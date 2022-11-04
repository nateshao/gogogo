---
title: day2-Go语言上手-工程实践
date: 2022-05-09 23:33:51
tags: 
- Go学习路线
- 字节跳动青训营
---







[TOC]

![day2-Go语言上手-工程实践](https://cdn.jsdelivr.net/gh/nateshao/images/20220512154323.jpg)

## 「Go 语言上手 - 工程实践」第三届字节跳动青训营 - 后端专场

这是我参与「第三届青训营 -后端场」笔记创作活动的的第2篇笔记

<img src="https://cdn.jsdelivr.net/gh/nateshao/images/20220510212909.png" style="zoom:50%;" />

同时这也是课表的第二天课程

# 01.语言进阶

从并发编程的视角待大家了解Go高性能的本质。



## 01.并发VS并行

Go可以充分发挥多核优势，高放运行

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508153156.png)

## 1.1 Goroutine

线程：用户态，轻量级线程，栈MB级别。

协程:内核态，线程跑多个协程，栈KB级别。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508153837.png)



在go里面快速开启一个协程

快速打印hello goroutine : 0~hello goroutine : 4

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508155110.png)

## 1.2 CSP (Communicating Sequential Processes)

提倡**通过通信共享内存**而不是通过共享内存而实现通信![通道-Gorountine.drawio](https://cdn.jsdelivr.net/gh/nateshao/images/20220508160810.png)



## 1.3 Channel

make(chan元素类型，[缓冲大小])

- 无缓冲通道make(chan int) 

- 有缓冲通道make(chan int,2)

  ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508161430.png)

![image-20220508162918192](https://cdn.jsdelivr.net/gh/nateshao/images/20220508162918.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508163823.png)



## 1.4并发安全Lock

对变量执行2000次+1操作，5个协程并发执行

| ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508170653.png) |
| ------------------------------------------------------------ |
| ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508190540.png) |
| ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508190617.png) |

## 小结

这里主要涉及3个方面：

1. 一个是协程Goroutine，通过高效的调度模型实现高并发操作；
2. 一个是通道channel,通过通信实现共享内存；
3. 最后sync相关关键字，实现并发安全操作和协程间的同步。



# 02.依赖管理

> 背景Go | 依赖管理演进 | Go Module实践

了解GO语言依赖管理的演进路线。

**背景**

对于hello world以及类似的单体函数只需要依赖原生SDK,而实际工程会相对复杂，我们不可能基于标准库0~1编码搭建，而更多的关注业务逻辑的实现，而其他的涉及框架、日志、driver、 以及collection等一 系列依赖都会通过sdk的方式引入，这样对依赖包的管理就显得尤为重要。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508192311.png)

## 2.1Go依赖管理演进

而Go的依赖管理主要经历了3个阶段，分别是，GOPATH-->Go Vendor-->Go Module 。到目前被广泛应用的go module,整个演进路线主要围绕实现两个目标来迭代发展的

![Go Module.drawio (1)](https://cdn.jsdelivr.net/gh/nateshao/images/20220508192804.png)



## 2.1.1 GOPATH

GOPATH是Go语言支持的一个环境变量，value是GO项目的工作区。
目录有以下结构: 

1. bin：存放Go项目编译生成的二进制文件。
2. pkg：存放编译的中间产物，加快编译速度。
3. src：存放Go项目的源码:。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508193234.png)







## 2.1.1 GOPATH-弊端

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508195748.png)

如图，同一个pkg，有2个版本，A->A0,B->B0,而src下只能有一 个版本存在， 那AB项目无法保证都能编译通过。也就是在gopath管理模式下， 如果多个项目依赖同一个库， 则依赖该库是同一份代码，所以不同项目不能依赖同一一个库的不同版本，这很显然不能满足我们的项目依赖需求。为了解决这问题，govender出现了 。

## 2.1.2 Go Vendor-弊端

Vendor是当前项目中的一一个目录，其中存放了当前项目依赖的副本。在Vendor机制下， 如果当前项目存在Vendor目录，会优先使用该目录下的依赖，如果依赖不存在，会从GOPATH中寻找; 但vendor无法很好解决依赖包的版本变动问题和一个项目依赖同一个包的不同版本的问题，下面我们看一 个场景

![Untitled Diagram.drawio (3)](https://cdn.jsdelivr.net/gh/nateshao/images/20220508205622.png)



## 2.1.3 Go Module

1. 通过go.mod文件管理依赖包版本
2. 通过goget/gomod指令工具管理依赖包

GO Modules是Go语言言方推出的依赖管理系统，解决了之前依赖管理系统存在的诸如无法依赖同一个库的多 个版本等问题，go moule从Go 1.11开始实验性引入，Go 1.16默认开启;我们般都读为g mod,我们也先统下名称

## 2.2依赖管理三要素

对于Java选手而言就是可以类比下maven。

1. 配置文件，描述依赖go .mod
2. 中心仓库管理依赖库Proxy
3. 本地工具go get/mod

## 2.3.1 依赖配置- go.mod

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508213931.png)



## 2.3.2依赖配置version



go path和go vendor都是源码副本方式依赖，没有版本规则概念，而go mod为 了放方便管理则定义了版本规则，分为语义化版本; 其中语义化版本包括不同的MAJOR版本表示是不兼容的AP,所以即使是同一个库，MAJOR 版本不同也会被认为是不同的模块: MINOR 版本通常是新增函数或功能，向后兼容；而patch 版本一般是修复 bug 基础版本前缀是和语义化版本一样的;时间戳(yyymmddhhmmss),也就是提交Commit的时间，最后是校验码(abcdefabcdef),包含12位的哈希前缀;每次提交commit后Go都会默认生成一个伪版本号。

## 2.3.3依赖配置- indirect

indirect后缀，表示go.mod对应的当前模块，没有直接导入该依赖模块的包，也就是非直接依赖，表示间接依赖

## 2.3.4依赖配置- incompatible

下一个常见是的是incompatible,主版本2 +模块会在模块路径增加/vN后缀，这能让go module按照不同的模块来处理同一个项目不同主版本的依赖。由于go module是1。 11实验性引入所以这项规则提出之前已经有一些仓库打 上了2或者更高版本的tag了，为了兼容这部分仓库，对于没有go.mod文件并且主版本在2或者以上的依赖，会在版本号后加上+incompatible后缀

## 2.3.4依赖配置依赖图

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508214440.png)



如果X项目依赖了A、B两个项目，且A、B分别依赖了C项目的v1.3、v1.4两个版本，最终编译时所使用的C项目的版本为如下哪个选项?
(单选)
A. v1.3
**B. v1.4 **
C. A用到C时用v1. 3编译, B用到C时用v1.4编译

答案：B,  选择最低的兼容版本

## 2.3.5依赖分发-回源

> go module的依赖分发。 也就是从哪里下载，如何下载的问题?

github是比较常见给的代码托管系统平台，而Go Modules系统中定义的依赖，最终可以对应到多版本代码管理系统中某一项目的特定提交或版本， 这样的话，对于go.mod中定义的依赖，则直接可以从对应仓库中下载指定软件依赖，从而完成依赖分发。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508214905.png)

但直接使用版本管理仓库下载依赖，存在多个问题，首先无法保证构建确定性:软件作者可以直接代码平台增加修改/删除软件版本，导致下次构建使用另外版本的依赖，或者找不到依赖版本。无法保证依赖可用性:依赖软件作者可以直接代码平台删除软件，导致依赖不可用;大幅增加第三方代码托管平台压力。

## 2.3.5依赖分发-Proxy

而go proxy就是解决这些问题的方案，Go Proxy是一个服务站点， 它会缓源站中的软件内容，缓存的软件版本不会改变，并且在源站软件删除之后依然可用，从而实现了供"immutability"和"available”的依赖分发;使用Go Proxy之后，构建时会直接从Go Proxy站点拉取依赖。类比项目中，下游无法满足我们上游的需求

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508215434.png)

## 2.3.6依赖分发-变量GOPROXY

GOPROXY="https://proxy1.cn, https://proxy2.cn ,direct'

服务站点URL列表，“direct" 表示源站

![Untitled Diagram.drawio (6)](https://cdn.jsdelivr.net/gh/nateshao/images/20220508215749.png)



下面讲一下go proxy的使用，Go Modules通过GOPROXY环境变量控制如何使用Go Proxy; GOPROXY是一个Go Proxy站点URL列表，可以使用directr表示源站。对于示例配置，整体的依赖寻址路径， 会优先从proxy1下载依赖， 如果proxy1不存在，后下钻proxy2寻找，如果proxy2, 中不存在则会回源到源站直接下载依赖，缓存到proxy站点中。

## 2.3.7工具- go get

![简单](https://cdn.jsdelivr.net/gh/nateshao/images/20220508220420.png)

## 2.3.8工具-go mod

尽量提交之前执行下go tidy,减少构建时无效依赖包的拉取

![简单 (1)](https://cdn.jsdelivr.net/gh/nateshao/images/20220508220630.png)

**依赖管理二要素**

1. 配置文件，描述依赖   go.mod
2. 中心仓库管理依赖库  Proxy
3. 本地工具 go get/mod

## 02小结

1. Go依赖管理演进
2. Go Module依赖管理方案







# 03.测试

在实际工程开发中，另一个重要概念就是单元测试，这里我们主要讲解go测试相关的内容，包括**单元测试 、Mock测试以及基准测试**。

从单元测试实践出发，提升大家的质量意识。

**事故**

1. 营销配置错误，导致非预期用户享受权益，资金损失10w+
2. 用户提现，幂等失效，短时间可以多次提现，资金损失20w+
3. 代码逻辑错误，广告位被占，无法出广告，收入损失500w+
4. 代码指针使用错误，导致APP不可用，损失上kw+。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508223548.png)

测试是避免事故的最后一道屏障

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508223814.png)

测试一般分为， 回归测试一般是QA同学 手动通过终端回归一些固定的主流程场景，集成测试是对系统功能维度做测试验证，而单元测试测试开发阶段，开发者对单独的函数、模块做功能验证，层级从上至下，测试成本逐渐减低，而测试覆盖率确逐步上升，所以单元测试的覆盖率-定程度上决定这代码的质量。

## 3.1单元测试

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508223938.png)

单元测试主要包括，输入，测试单元，输出，以及校对，单元的概念比较广，包括接口，函数，模块等;用最后的校对来保证代码的功能与我们的预期相符;单侧一方面可以保证质量，在整体覆盖率足够的情况下，一定程度上既保证了新功能本身的正确性，又未破坏原有代码的正确性。另一方面可以提升效率，在代码有bug的情况下，通过编写单测，可以在一个较短周期内定位和修复问题。



## 3.1.1单元测试-规则

从文件上就很好了区分源码和测试代码，以Test开头， 且连接的第一 个字母大写

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508224450.png)



## 3.1.2单元测试例子

```go
func HelloTom() string {
	return "Tom"
}
```

测试代码

```go
func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectOutput := "Tom"
	assert.Equal(t, expectOutput, output)
}
```

## 3.1.3单元测试-运行

go test [flags] [packages]

![image-20220508225855252](https://cdn.jsdelivr.net/gh/nateshao/images/20220508225855.png)



## 3.1.4单元测试- assert

```go
package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectOutput := "Tom"
	assert.Equal(t, expectOutput, output)
}
```

## 3.1.5单元测试-覆盖率

- 一般覆盖率: 50%~60%，较高覆盖率80%+。

- 测试分支相互独立、全面覆盖。

- 测试单元粒度足够小，函数单一职责。(要求函数体足够小，这样就比较简单的提升覆盖率，也符合函数设计的单一职责。)

对于资金型服务，覆盖率可能要求达到80%以上

## 3.2单元测试-依赖

我们的单测需要保证稳定性和幕等性，稳定是指相互隔离，能在任何时间，任何环境，运行测试。幂等是指每一次测试运行都应该产生与之前一样的结果。而要实现这一 目的就要用到mock机制。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508231111.png)

## 3.3单元测试-文件处理

```go
package test

import (
	"bufio"
	"os"
	"strings"
)

func ReadFirstLine() string {
	open, err := os.Open("log")
	defer open.Close()
	if err != nil {
		return ""
	}
	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func ProcessFirstLine() string {
	line := ReadFirstLine()
	destLine := strings.ReplaceAll(line, "11", "00")
	return destLine
}
```

测试类

```go
package test

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessFirstLine(t *testing.T) {
	firstLine := ProcessFirstLine()
	assert.Equal(t, "line00", firstLine)
}

func TestProcessFirstLineWithMock(t *testing.T) {
	monkey.Patch(ReadFirstLine, func() string {
		return "line110"
	})
	defer monkey.Unpatch(ReadFirstLine)
	line := ProcessFirstLine()
	assert.Equal(t, "line000", line)
}
```





## 3.4单元测试- Mock

monkey : https://github.com/bouk/monkey

快速Mock函数

- 为一个函数打桩
- 为一个方法打桩

```go
package test

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessFirstLine(t *testing.T) {
	firstLine := ProcessFirstLine()
	assert.Equal(t, "line00", firstLine)
}

func TestProcessFirstLineWithMock(t *testing.T) {
	monkey.Patch(ReadFirstLine, func() string {
		return "line110"
	})
	defer monkey.Unpatch(ReadFirstLine)
	line := ProcessFirstLine()
	assert.Equal(t, "line000", line)
}
```

## 3.5基准测试

Go语言还提供了基准测试框架，基准测试是指测试一段程序的运行性能及耗费CPU的程度。而我们在实际项目开发中，经常会遇到代码性能瓶颈，为了定位问题经常要对代码做性能分析，这就用到了基准测试。使用方法类似于单元测试.

- 优化代码，需要对当前代码分析
- 内置的测试框架提供了基准测试的能力

## 3.5.1基准测试例子

这里举一个服务器负载均衡的例子，首先我们有10个服务器列表，每次随机执行select函数随机选择一 个执行。

代码如下：

```go
package benchmark

import (
	"github.com/bytedance/gopkg/lang/fastrand"
	"math/rand"
)

var ServerIndex [10]int

func InitServerIndex() {
	for i := 0; i < 10; i++ {
		ServerIndex[i] = i+100
	}
}

func Select() int {
	return ServerIndex[rand.Intn(10)]
}

func FastSelect() int {
	return ServerIndex[fastrand.Intn(10)]
}
```





## 3.5.3基准测试-优化

https://github.com/bytedance/gopkg

```go
func FastSelect() int {
	return ServerIndex[fastrand.Intn(10)]
}
```

而公司为了解决这一随机性能问题， 开源了一个高性能随机数方法fastrand, 下面有开源地址;我们这边再做一下基准测试， 性能提升了百倍。主要的思路是牺牲了一定的数列一 致性，在大多数场景是适用的，同学在后面遇到随机的场景可以尝试用一下。

# 04.项目实践 

通过项目需求、需求拆解、逻辑设计、代码实现带领大家感受下真实的项目开发。

主要包含：需求设计 代码开发 测试运行

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508231908.png)

大家应该都是从掘金的社区话题入口报名的，都看到过这个页面，页面的功能包括话题详情，回帖列表，支持回帖，点赞，和回帖回复，我们今天就以此为需求模型，开发一个该页面交涉及的服务端小功能。

## 4.1需求描述

社区话题页面

- 展示话题(标题，文字描述)和回帖列表
- 暂不考虑前端页面实现，仅仅实现一个本地web服务
- 话题和回帖数据用文件存储

## 4.2需求用例

主要涉及功能点，用户浏览消费，涉及页面的展示，包括话题内容和回帖的列表，其实从图中我们应该会抽出2个实体的，而实体的属性有哪些，他们之间的联系

![Untitled Diagram.drawio (10)](https://cdn.jsdelivr.net/gh/nateshao/images/20220508232132.png)

## 4.3 ER图-Entity Relationship Diagram

结构设计。E-R国，用来典型的分层结构设计模型。 有了模型实体，属性以及之间的联系，对我们后续做开发就提供了比较清晰的思路。回到需求。两个个实体主要包括，实体的属性，有了实体模型，下一步就是思考代码

![image-20220508232234260](https://cdn.jsdelivr.net/gh/nateshao/images/20220508232234.png)



## 4.4分层结构

整体分为三层，repository数据层， service逻辑层， controoler视图层。

数据层关联底层数据模型，也就是这里的model,封装外部数据的增删改查，我们的数据存储在本地文件， 通过文件操作拉取话题， 帖子数据;数据层面向逻辑层，对service层透明， 屏蔽下游数据差异，也就是不管下游是文件，还是数据库，还是微服务等，对service层的接模型是不变的。

Servcie逻辑层处理核心业务逻辑，计算打包业务实体entiy,对应我们的需求，就是话题页面，包括话题和回帖列表，并上送给视图层;

Cortroller视图层负责处理和外部的交互逻辑，以view视图的形式返回给客户端，对于我们需求，我们封装json格式化的请求结果，api形式访问就好,

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220508232343.png)

- 数据层：数据Model,外部数据的增删改查
- 逻辑层：业务Entity,处理核心业务逻辑输出
- V视图层：视图view,处理和外部的交互逻辑

## 4.5组件工具

1. **Gin** 高性能go web框架 https://github.com/gin-gonic/gin#installation
   
2. Go Mod
   go mod init

   go get  gopkg.in/gin-gonic/gin.v1@v1.3.0

介绍下开发涉及的基础组件和工具，首先是gin, 高性能开源的go web框架，我们基于gin搭建web服务器，在课程手册应该提到了，这里我们只是简单的使用，主要涉及路由分发，不会涉及其他复杂的概念。

因为我们引入了web框架，所以就涉及go module依赖管理，如前面依赖管理课程内容讲解，我们首先通过go mod是初始化go mod管理配置文件，然后go get下载gin依赖，这里显示用了V1.3.0版本。

有了框架依赖，我们只需要关注业务本身的实现，从reposity --> service  --> contoller我们一步步实现。希望大家能跟上我的节奏，从0~1 实现这个项目，如果时间问题，大家可以一步步copy一下，主要是走一半开发思路。

## 4.6 Repository

|                            Topic                             |                             Post                             |
| :----------------------------------------------------------: | :----------------------------------------------------------: |
| ![image-20220508232757464](https://cdn.jsdelivr.net/gh/nateshao/images/20220508232757.png) | ![image-20220508232808335](https://cdn.jsdelivr.net/gh/nateshao/images/20220508232808.png) |
|                        QueryTopicByld                        |                     QueryPostsByParentld                     |



## 4.6 Repository-index

好的，一方面查询我们可以用全扫描遍历的方式，但是这虽然能达到我们的目的，但是并非高效的方式，所以这里引出索引的概念，索引就像书的目录，可以引导我们快速查找定位我们需要的结果；这里我们用map实现内存索引，在服务对外暴露前，利用文件元数据初始化全局内存索引，这样就可以实现0 (1) 的时间复杂度查找操作。



![](https://cdn.jsdelivr.net/gh/nateshao/images/20220509225729.png)

Ok， 下面是具体的实现，我们过一-下， 首先是打开文件，基于file初始化scanner, 通过迭代器方式遍历数据行，转化为结构体存储至内存map，ok,这就是初始化话题内存索引。

![image-20220509230404980](https://cdn.jsdelivr.net/gh/nateshao/images/20220509230405.png)

## 4.6 Repository-查询

有了4.6 Repository-查询内存索引，下一步就是实现查询操作就比较简单了，直接根据查询key获得map中的value就好了 ，这里用到了sync.once,主要适用高并发的场景下只执行一次的场景， 这里的基于once的实现模式就是我们平常说的单例模式， 减少存储的浪费。

![image-20220509230838568](https://cdn.jsdelivr.net/gh/nateshao/images/20220509230950.png)



## 4.7 Service

实体

```go
type PageInfo struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}
```

流程：

![image-20220509231153125](https://cdn.jsdelivr.net/gh/nateshao/images/20220509231153.png)

```go
// 代码流程编排
func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}
```

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220509231617.png)



关于prepareInfo方法，话题和回帖信息的获取都依赖topicid,这样2这就可以并行执行，提高执行效率。大家在后期做项目开发中，一定要思考流程是否可以并，通过压榨CPU，降低接口耗时，不要一味的串行实现， 浪费多核cpu的资源。

## 4.8 Controller

Service实现完成，下面就是controller层。 这里我们定义一个view对象，通过code msg打包业务状态信息，用data承载业务实体信息

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220509231929.png)



## 4.9 Router



```go
func main() {
	// 初始化数据索引
	if err := Init("./data/"); err != nil {
		os.Exit(-1)
	}
	// 初始化引|擎配置
	r := gin.Default()
	// 构建路由
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := cotroller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	// 启动服务
	err := r.Run()
	if err != nil {
		return
	}
}
```





## 4.10运行

终端执行：curl --location --request GET 'http://127.0.0.1:8080/community/page/get/2'

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220509232651.png)

控制台输出：

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220510213341.png)

好的，以上就是对社区话题页面需求的整个实现流程，这样我们从项目拆解，代码设计落地，最后测试运行就跑通了整个的项目流程，为大家后期实现项目提供了一定的开发思路。 当然实际项目较我们实现的需求会复杂很多，不过大家也不必担心，可以通过大拆小的思路，将大需求拆解为小需求的思路来分析解决，遇到问题，各个击破，同时做好充分的测试。
