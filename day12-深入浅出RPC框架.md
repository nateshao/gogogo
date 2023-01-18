---
title: day12-深入浅出RPC框架
date: 2022-11-05 23:20:36
tags:
- Go学习路线
- 字节跳动青训营
---

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221105233121.jpg)

这是我参与「第三届青训营 - 后端场」笔记创作活动的的第12篇笔记。*PC端阅读效果更佳，点击文末：**阅读原文**即可。*

## 「深入浅出RPC框架」 第三届字节跳动青训营 - 后端专场

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221105115711.png)

## 课程预习

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221105233648.png)



### RPC 的基本概念

- RPC的概念模型：User、User-Stub、RPC-Runtime、Server-Stub、Server
  - 来自论文《[Implementing Remote Procedure Calls](https://link.juejin.cn?target=https%3A%2F%2Fweb.eecs.umich.edu%2F~mosharaf%2FReadings%2FRPC.pdf)》

- IDL(Interface Definition Language) 文件
  - Thrift
  - Protobuf

- 生成代码

- 编解码（序列化/反序列化）

- 通信协议
  - 应用层协议

- 网络通信
  - IO 网络模型
    - blocking IO
    - unblocking IO
    - IO multiplexing
    - signal driven IO
    - asynchronous IO
  - 传输层协议
    - TCP
    - UDP

### RPC 框架分层设计

- 编解码层
  - 数据格式：
    - 语言特定格式
    - 文本格式
    - 二进制编码
      - TLV 编码：Thrift 使用 TLV 编码
      - Varint 编码：Protobuf 使用 Varint 编码
  - 选项：
    - 兼容性
    - 通用型
    - 性能

- 传输协议层
  - 消息切分
    - 特殊结束符
    - 变长协议：length+body
  - 协议构造
    - 以 Thrift 的 [THeader](https://link.juejin.cn?target=https%3A%2F%2Fgithub.com%2Fapache%2Fthrift%2Fblob%2Fmaster%2Fdoc%2Fspecs%2FHeaderFormat.md) 协议为例讲解

- 网络通信层
  - 网络库
  - 核心指标
    - 吞吐高
    - 延迟低

### RPC 框架的核心指标

- 稳定性
  - 保障策略
    - 熔断
    - 限流
    - 超时
  - 请求成功率
    - 负载均衡
    - 重试
  - 长尾请求
    - BackupRequest

- 易用性
  - 开箱即用
  - 周边工具

- 扩展性

- 观测性
  - Log
  - Metric
  - Tracing
  - 内置观测性服务

- 高性能

### 字节内部 Kitex 实践分享

- [Kitex](https://link.juejin.cn/?target=https%3A%2F%2Fgithub.com%2Fcloudwego%2Fkitex) 整体架构

- 自研网络库 [Netpoll](https://link.juejin.cn/?target=https%3A%2F%2Fgithub.com%2Fcloudwego%2Fnetpoll)

- [性能优化](https://link.juejin.cn/?target=https%3A%2F%2Fwww.infoq.cn%2Farticle%2Fspasfyqgaaid5rguinl4)：
  - 网络库优化
  - 编解码优化

- 合并部署

---

## 课中

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106145535.png)

# 01.基本概念

## 1.1 本地函数调用

```go
func main( ){
	var a = 2
	var b = 3
	result := calculate(a, b)
	fmt.Println(result)
	return
}
func calculate(x, y int) {
	z := x*y
	return z
}
```

1. 将a和b的值压栈
2. 通过函数指针找到calculate函数，进入函数取出栈中的值2和3，将其赋予x和y 
3. 计算x*y，并将结果存在z
4. 将Z的值压栈，然后从calculate返回
5. 从栈中取出z返回值，并赋值给result

以上步骤只是为了说明原理。事实上编译器经常会做优化，对于参数和返回值少的情况会直接将其存放在寄存器，而不需要压栈弹栈的过程，甚至都不需要调用call，而直接做inline操作。

## 1.2远程函数调用(RPC - Remote Procedure Calls)

| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106150814.png) | ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106150912.png) |
| ------------------------------------------------------------ | ------------------------------------------------------------ |

#### 函数映射

**我们怎么告诉支付服务我们要调用付款这个函数，而不是退款或者充值呢？**

- 在本地调用中，函数体是直接通过**函数指针**来指定的，我们调用哪个方法，编译器就自动帮我们调用它相应的函数指针。

- 但是在远程调用中，函数指针是不行的，**因为两个进程的地址空间是完全不一样的**。所以函数都有自己的一个ID，在做 RPC的时候要附上这个 ID，还得有个 ID 和函数的对照关系表，通过 ID找到对应的函数并执行。 

**客户端怎么把参数值传给远程的函数呢？** 

- 在本地调用中，我们只需要把参数压到栈里，然后让函数自己去栈里读就行。
- 但是在远程过程调用时，客户端跟服务端是不同的进程，不能通过内存来传递参数。这时候就需要**客户端把参数先转成一个字节流，传给服务端后，再把字节流转成自己能读取的格式。**  

**远程调用往往用在网络上，如何保证在网络上高效稳定地传输数据？**

## 1.3 RPC概念模型

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106151621.png)

理解RPC调用的完整过程

## 1.4 一次RPC的完整过程

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106151806.png)

IDL (Interface description language)文件

- **IDL**通过一种中立的方式来描述接口，使得**在不同平台上运行的对象和用不同语言编写的程序可以相互通信**

生成代码

- 通过编译器具把IDL文件转换成语言对应的静态库

编解码

- 从内存中表示到字节序列的转换称为编码，反之为解码，也常叫做**序列化和反序列化**

通信协议

- 规范了数据在网络中的传输内容和格式。除必须的请求/响应数据外，通常还会包含额外的元数据

网络传输

- 通常基于成熟的网络库走TCP/UDP传输

相比本地函数调用，远程调用的话我们不知道对方有哪些方法，以及参数长什么样，所以需要有一种方式来描述或者说声明我有哪些方法，方法的参数都是什么样子的，这样的话大家就能按照这个来调用，这个描述文件就是 **IDL 文件**。 

服务双方是通过约定的规范进行远程调用，双方都依赖同一份IDL文件，需要通过工具来生成对应的生成文件，具体调用的时候用户代码需要依赖生成代码，所以可以把用户代码和生成代码看做一个整体。

编码只是解决了跨语言的数据交换格式，但是如何通讯呢？需要制定通讯协议，以及数据如何传输？我的网络模型如何呢？那就是这里的 transfer 要做的事情。

## 1.5 RPC的好处

1. 单一职责，有利于分工协作和运维开发
2. 可扩展性强，资源使用率更优
3. 故障隔离，服务的整体可靠性更高

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106212649.png)

单一职责，开发（采用不同的语言）、部署以及运维（上线独立）都是独立的。

可扩展性强，例如压力过大的时候可以独立扩充资源，底层基础服务可以复用，节省资源某个模块发生故障，不会影响整体的可靠性。

## 1.6 RPC带来的问题

1. 服务宕机，对方应该如何处理?
2. 在调用过程中发生网络异常，如何保证消息的可达性?
3. 请求量突增导致服务无法及时处理，有哪些应对措施?

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106213225.png)

# 02 分层设计

编解码层 | 协议层 | 网络通信层

## 2.1分层设计-以Apache Thrift为例

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106213443.png)

## 2.2 编解码层-生成代码

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106213646.png)

## 2.3 编解码层-数据格式

语言特定的格式

- 许多编程语言都内建了将内存对象编码为字节序列的支持，例如Java有java.io.Serializable

文本格式

- JSON、XML、 CSV 等文本格式，具有人类可读性

二进制编码

- 具备跨语言和高性能等优点，常见有Thrift 的BinaryProtocol, Protobuf 等

**语言特定编码格式**：这种编码形式好处是非常方便，可以用很少的额外代码实现内存对象的保存与恢复，这类编码通常与特定的编程语言深度绑定，其他语言很难读取这种数据。如果以这类编码存储或传输数据，那你就和这门语言绑死在一起了。安全和兼容性也是问题 

**文本格式**：文本格式具有人类可读性，数字的编码多有歧义之处，比如XML和CSV不能区分数字和字符串，JSON虽然区分字符串和数字，但是不区分整数和浮点数，而且不能指定精度，处理大量数据时，这个问题更严重了；没有强制模型约束，实际操作中往往只能采用文档方式来进行约定，这可能会给调试带来一些不便。 由于JSON在一些语言中的序列化和反序列化需要采用反射机制，所以在性能比较差； 

**二进制编码**：实现可以有很多种，TLV 编码 和 Varint 编码

## 2.5 编解码层-二进制编码

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106215826.png)

第一个byte是类型，主要用来表示是string还是int还是list等等。这里不写key的字符串了，比如上面的userName，favoriteNumber等等，取而代之的是一个field tag的东西，这个会设置成1,2,3和上面的schema中key字符串前面的数字，也就是用这里来取代了具体的key值，从而减小的总体的大小，这里打包后压缩到 59个字节 TLV编码结构简单清晰，并且扩展性较好，但是由于增加了Type和Length两个冗余信息，有额外的内存开销，特别是在大部分字段都是基本类型的情况下有不小的空间浪费。



## 2.6 编解码层-选型

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106220517.png)

通用性： 通用性有两个层面的意义： 

- 第一、技术层面，序列化协议是否支持跨平台、跨语言。如果不支持，在技术层面上的通用性就大大降低了。 
- 第二、流行程度，序列化和反序列化需要多方参与，很少人使用的协议往往意味着昂贵的学习成本；另一方面，流行度低的协议，往往缺乏稳定而成熟的跨语言、跨平台的公共包。  

兼容性： 移动互联时代，业务系统需求的更新周期变得更快，新的需求不断涌现，而老的系统还是需要继续维护。如果序列化协议具有良好的可扩展性，支持自动增加新的业务字段，而不影响老的服务，这将大大提供系统的灵活度。  

性能： 

- 第一、空间开销（Verbosity）， 序列化需要在原有的数据上加上描述字段，以为反序列化解析之用。如果序列化过程引入的额外开销过高，可能会导致过大的网络，磁盘等各方面的压力。对于海量分布式存储系统，数据量往往以TB为单位，巨大的的额外空间开销意味着高昂的成本。 
- 第二、时间开销（Complexity），复杂的序列化协议会导致较长的解析时间，这可能会使得序列化和反序列化阶段成为整个系统的瓶颈。

## 2.8 协议层-概念

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106221355.png)

协议是双方确定的交流语义，比如：我们设计一个字符串传输的协议，它允许客户端发送一个字符串，服务端接收到对应的字符串。这个协议很简单，首先发送一个4字节的消息总长度，然后再发送1字节的字符集charset长度，接下来就是消息的payload，字符集名称和字符串正文。 

特殊结束符：过于简单，对于一个协议单元必须要全部读入才能够进行处理，除此之外必须要防止用户传输的数据不能同结束符相同，否则就会出现紊乱 HTTP 协议头就是以回车(CR)加换行(LF)符号序列结尾。

 变长协议：一般都是自定义协议，有 header 和 payload 组成，会以定长加不定长的部分组成，其中定长的部分需要描述不定长的内容长度，使用比较广泛

## 2.9 协议层-协议构造

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106222651.png)

**LENGTH** 字段 32bits，包括数据包剩余部分的字节大小，不包含 LENGTH 自身长度 

**HEADER MAGIC** 字段16bits，值为：0x1000，用于标识 协议版本信息，协议解析的时候可以快速校验 FLAGS 字段 16bits，为预留字段，暂未使用，默认值为 0x0000 

**SEQUENCE NUMBER** 字段 32bits，表示数据包的 seqId，可用于多路复用，最好确保单个连接内递增 

**HEADER SIZE** 字段 16bits，等于头部长度字节数/4，头部长度计算从第14个字节开始计算，一直到 PAYLOAD 前（备注：header 的最大长度为 64K） 

**PROTOCOL ID** 字段 uint8 编码，取值有：~ 

ProtocolIDBinary = 0 

ProtocolIDCompact = 2 

**NUM TRANSFORMS** 字段 uint8 编码，表示 TRANSFORM 个数 

**TRANSFORM ID** 字段 uint8 编码，具体取值参考下文，表示压缩方式 zlib or snappy 

**INFO ID** 字段 uint8 编码，具体取值参考下文，用于传递一些定制的 meta 信息 

**PAYLOAD** 消息内容

## 2.10 协议层-协议解析

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106223140.png)

## 2.12 网络通信层- Sockets API

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106223243.png)

套接字编程中的客户端必须知道两个信息：**服务器的 IP 地址，以及端口号**。 

socket函数创建一个套接字，bind 将一个套接字绑定到一个地址上。listen 监听进来的连接，backlog的含义有点复杂，这里先简单的描述：指定挂起的连接队列的长度，当客户端连接的时候，服务器可能正在处理其他逻辑而未调用accept接受连接，此时会导致这个连接被挂起，内核维护挂起的连接队列，backlog则指定这个队列的长度，accept函数从队列中取出连接请求并接收它，然后这个连接就从挂起队列移除。如果队列未满，客户端调用connect马上成功，如果满了可能会阻塞等待队列未满（实际上在Linux中测试并不是这样的结果，这个后面再专门来研究）。Linux的backlog默认是128，通常情况下，我们也指定为128即可。 

connect 客户端向服务器发起连接，accept 接收一个连接请求，如果没有连接则会一直阻塞直到有连接进来。得到客户端的fd之后，就可以调用read, write函数和客户端通讯，读写方式和其他I/O类似 ∂ 

read 从fd读数据，socket默认是阻塞模式的，如果对方没有写数据，read会一直阻塞着： 

write 写fd写数据，socket默认是阻塞模式的，如果对方没有写数据，write会一直阻塞着： 

socket 关闭套接字，当另一端socket关闭后，这一端读写的情况： 尝试去读会得到一个EOF，并返回0。 尝试去写会触发一个SIGPIPE信号，并返回-1和errno=EPIPE，SIGPIPE的默认行为是终止程序，所以通常我们应该忽略这个信号，避免程序终止。 如果这一端不去读写，我们可能没有办法知道对端的socket关闭了。

## 2.13 网络通信层-网络库

提供易用API

- 封装底层Socket API
- 连接管理和事件分发

功能

- 协议支持: tcp、 udp和uds等
- 优雅退出、异常处理等

性能

- 应用层buffer减少copy
- 高性能定时器、对象池等

# 03 关键指标

## 3.1 稳定性-保障策略

熔断：保护调用方，防止被调用的服务出现问题而影响到整个链路

限流：保护被调用方，防止大流量把服务压垮

超时控制：避免浪费资源在不可用节点上

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106230507.png)

熔断： 一个服务 A 调用服务 B 时，服务 B 的业务逻辑又调用了服务 C，而这时服务 C 响应超时了，由于服务 B 依赖服务 C，C 超时直接导致 B 的业务逻辑一直等待，而这个时候服务 A 继续频繁地调用服务 B，服务 B 就可能会因为堆积大量的请求而导致服务宕机，由此就导致了服务雪崩的问题。

限流： 当调用端发送请求过来时，服务端在执行业务逻辑之前先执行检查限流逻辑，如果发现访问量过大并且超出了限流条件，就让服务端直接降级处理或者返回给调用方一个限流异常。

超时： 当下游的服务因为某种原因响应过慢，下游服务主动停掉一些不太重要的业务，释放出服务器资源，避免浪费资源  

从某种程度上讲超时、限流和熔断也是一种服务降级的手段。

## 3.2 稳定性-请求成功率

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106231432.png)

注意，因为重试有放大故障的风险。

首先，重试会加大直接下游的负载。如下图，假设 A 服务调用 B 服务，重试次数设置为 r（包括首次请求），当 B 高负载时很可能调用不成功，这时 A 调用失败重试 B ，B 服务的被调用量快速增大，最坏情况下可能放大到 r 倍，不仅不能请求成功，还可能导致 B 的负载继续升高，甚至直接打挂。  

防止重试风暴，限制单点重试和限制链路重试

## 3.3 稳定性-长尾请求

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106232034.png)

长尾请求一般是指明显高于均值的那部分占比较小的请求。 业界关于延迟有一个常用的P99标准， P99 单个请求响应耗时从小到大排列，顺序处于99%位置的值即为P99 值，那后面这 1%就可以认为是长尾请求。在较复杂的系统中，长尾延时总是会存在。造成这个的原因非常多，常见的有网络抖动，GC，系统调度。 

我们预先设定一个阈值 t3（比超时时间小，通常建议是 RPC 请求延时的 pct99 ），当 Req1 发出去后超过 t3 时间都没有返回，那我们直接发起重试请求 Req2 ，这样相当于同时有两个请求运行。然后等待请求返回，只要 Resp1 或者 Resp2 任意一个返回成功的结果，就可以立即结束这次请求，这样整体的耗时就是 t4 ，它表示从第一个请求发出到第一个成功结果返回之间的时间，相比于等待超时后再发出请求，这种机制能大大减少整体延时。

## 3.4 稳定性-注册中间件

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106232216.png)

Kitex Client 和 Server 的创建接口均采用 Option 模式，提供了极大的灵活性，很方便就能注入这些稳定性策略

## 3.5 易用性

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106232342.png)

Kitex 使用 Suite 来打包自定义的功能，提供「一键配置基础依赖」的体验 

## 3.6 扩展性

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106232458.png)

一次请求发起首先会经过治理层面，治理相关的逻辑被封装在middleware中，这些middleware会被构造成一个有序调用链逐个执行，比如服务发现、路由、负载均衡、超时控制等，mw执行后就会进入到remote 模块，完成与远端的通信

## 3.7 观测性

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106232553.png)

除了传统的 Log、Metric、Tracing 三件套之外，对于框架来说可能还不够，还有些框架自身状态需要暴露出来，例如当前的环境变量、配置、Client/Server初始化参数、缓存信息等

## 3.8 高性能

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106232702.png)

这个图需要换一下 这里分两个维度，

高性能意味着高吞吐和低延迟，两者都很重要，甚至大部分场景下低延迟更重要。 多路复用可以大大减少了连接带来的资源消耗，并且提升了服务端性能，我们的测试中服务端吞吐可提升30%。 

右边的图帮助大家理解连接多路复用 调用端向服务端的一个节点发送请求，并发场景下，如果是非连接多路复用，每个请求都会持有一个连接，直到请求结束连接才会被关闭或者放入连接池复用，并发量与连接数是对等的关系。 

而使用连接多路复用，所有请求都可以在一个连接上完成，大家可以明显看到连接资源利用上的差异 

# 04 企业实践

> Kitex  企业内部大范围使用 go 语言进行开发，而 kitex 是内部多年最佳实践沉淀出来的一个高性能高可扩展性的 go RPC 框架，在内部有几万个微服务在使用，在去年也开源了回馈给了社区，并且收获了 4K stars。

## 4.1整体架构- Kitex

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106233249.png)

core是它的的主干逻辑，定义了框架的层次结构、接口，还有接口的默认实现，如中间蓝色部分所示，最上面client和server是对用户暴露的，client/server option的配置都是在这两个package中提供的，还有client/server的初始化，在第二节介绍kitex_gen生成代码时，大家应该注意到里面有client.go和server.go，虽然我们在初始化client时调用的是kitex_gen中的方法，其实大家看下kitex_gen下service package代码就知道，里面是对这里的 client/server的封装。 

client/server下面的是框架治理层面的功能模块和交互元信息，remote是与对端交互的模块，包括编解码和网络通信。 

右边绿色的byted是对字节内部的扩展，集成了内部的二方库还有与字节相关的非通用的实现，在第二节高级特性中关于如何扩展kitex里有介绍过，byted部分是在生成代码中初始化client和server时通过suite集成进来的，这样实现的好处是与字节的内部特性解耦，方便后续开源拆分。 

左边的tool则是与生成代码相关的实现，我们的生成代码工具就是编译这个包得到的，里面包括idl解析、校验、代码生成、插件支持、自更新等，未来生成代码逻辑还会做一些拆分，便于给用户提供更友好的扩展

## 4.2 自研网络库-背景

原生库无法感知连接状态

- 在使用连接池时，池中存在失效连接，影响连接池的复用。

原生库存在goroutine暴涨的风险

- 一个连接个goroutine的模式，由于连接利用率低下，存在大量goroutine占用调度开销，影响性能。

1. Go Net 使用 Epoll ET ，Netpoll 使用 LT。 
2. Netpoll 在大包场景下会占用更多的内存。 
3. Go Net 只有一个 Epoll 事件循环（因为 ET 模式被唤醒的少，且事件循环内无需负责读写，所以干的活少），而 Netpoll 允许有多个事件循环（循环内需要负责读写，干的活多，读写越重，越需要开更多 Loops）。 
4. Go Net 一个连接一个 Goroutine，Netpoll 连接数和 Goroutine 数量没有关系，和请求数有一定关系，但是有 Gopool 重用。
5. Go Net 不支持 Zero Copy，甚至于如果用户想要实现 BufferdConnection 这类缓存读取，还会产生二次拷贝。Netpoll 支持管理一个 Buffer 池直接交给用户，且上层用户可以不使用 Read(p []byte) 接口而使用特定零拷贝读取接口对 Buffer 进行管理，实现零拷贝能力的传递。

## 4.3 自研网络库- Netpoll

解决无法感知连接状态问题

- 引入epoll主动监听机制，感知连接状态

解决goroutine暴涨的风险

- 建立goroutine池，复用goroutine

提升性能

- 引入Nocopy Buffer,向上层提供NoCopy的调用接口，编解码层面零拷贝



1. go net 无法检测连接对端关闭（无法感知连接状态）  
   1. 在使用长连接池时，池中存在失效连接，严重影响了连接池的使用和效率。
   2. 希望通过引入 epoll 主动监听机制，感知连接状态。 

2. go net 缺乏对协程数量的管理  
   1. Kite 采取一个连接一个 goroutine 模式，由于连接利用率低，服务存在较多无用的 goroutine，占用调度开销，影响性能。\
   2. 希望建立协程池，提升性能。

netpoll基于epoll，同时采用Reactor模型，对于服务端则是主从Reactor模型，如右图所示：服务端的主reactor 用于接受调用端的连接，然后将建立好的连接注册到某个从Reactor上，从Reactor负责监听连接上的读写事件，然后将读写事件分发到协程池里进行处理。 

3. 为了提升性能，引入了 Nocopy Buffer，向上层提供 NoCopy 的调用接口，编解码层面零拷贝

## 4.4扩展性设计

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106233858.png)

kitex支持多协议的并且也是可扩展的，交互方式上前面已经说过支持ping-pong、streaming、oneway 

编解码支持thrift、Protobuf 

应用层协议支持TTHeader、Http2、也支持裸的thrift协议 

传输层目前支持TCP，未来考虑支持UDP、kernel-bypass的RDMA 如右图所示，框架内部不强依赖任何协议和网络模块，可以基于接口扩展，在传输层上则可以集成其他库进行扩展。

 目前集成的有自研的Netpoll，基于netpoll实现的http2库，用于mesh场景通过共享内存高效通信的shm-ipc，以后也可以增加对RDMA支持的扩展

## 4.5 性能优化-网络库优化

**调度优化**

- epoll wait在调度上的控制
- gopool重用goroutine降低同时运行协程数

**LinkBuffer**

- 读写并行无锁，支持nocopy地流式读写
- 高效扩缩容
- Nocopy Buffer池化，减少GC

**Pool**

- 引入内存池和对象池，减少GC开销



## 4.6 性能优化-编解码优化

**Codegen**

- 预计算并预分配内存，减少内存操作次数，包括内存分配和拷贝
- Inline减少函数调用次数和避免不必要的反射操作等
- 自研了Go语言实现的Thrift IDL解析和代码生成器，支持完善的Thrift IDL语法和语义检
  查，并支持了插件机制- Thriftgo

 **JIT**

- 使用JIT编译技术改善用户体验的同时带来更强的编解码性能，减轻用户维护生成代码的负担
- 基于JIT编译技术的高性能动态Thrift 编解码器- **Frugal**

序列化和反序列的性能优化从大的方面来看可以从时间和空间两个维度进行优化。

从兼容已有的 Binary 协议来看，空间上的优化似乎不太可行，只能从时间维度进行优化，包括下面的几点：

代码生成 code-gen 的优点是库开发者实现起来相对简单，缺点是增加业务代码的维护成本和局限性。 

 JIT编译（just-in-time compilation）狭义来说是当某段代码即将第一次被执行时进行编译，因而叫“即时编译”。 

即时编译 JIT 则将编译过程移到了程序的加载（或首次解析）阶段，可以一次性编译生成对应的 codec 并高效执行，目前公司内部正在尝试，压测数据表明性能收益还是挺不错的，目的是不损失性能的前提下，减轻用户的维护负担生成代码的负担。



## 4.7 合并部署

- 微服务过微，传输和序列化开销越来越大
- 将亲和性强的服务实例尽可能调度到同一个物理机，远程RPC调用优化为本地IPC调用

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106234246.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20221106234306.png)

## 课程总结

1. 从本地函数调用引出RPC的基本概念

2. 重点讲解了RPC框架的核心的三层，编解码层、协议层和网络传输层

3. 围绕RPC框架的核心指标，例如稳定性、可扩展性和高性能等，展开讲解相关的知识

4. 分享了字节跳动高性能RPC框架Kitex的相关实践





1. https://juejin.cn/post/7099742161540743198/
2. https://bytedance.feishu.cn/file/boxcnQ289aixfQUmvgtDDn98mmf
3. 青训营官方账号：https://juejin.cn/post/7099665398655615006/
   







































