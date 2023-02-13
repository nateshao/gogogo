---
title: day14-微服务架构原理与治理实践
date: 2023-01-15 12:58:41
tags:
- Go学习路线
- 字节跳动青训营
---

<img src="https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230117222648.jpg" style="zoom:200%;" />

这是我参与「第三届青训营 - 后端场」笔记创作活动的的第14篇笔记。*PC端阅读效果更佳，点击文末：**阅读原文**即可。*

## 「微服务架构原理与治理实践」 第三届字节跳动青训营 - 后端专场

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230117222718.png)

[TOC]

## 课前预习

### 微服务架构介绍

- 系统架构的演进历史
  - 单体架构
  - 垂直应用架构
  - 分布式架构
  - SOA架构
  - 微服务架构

- 微服务架构的三大要素
  - 服务治理
  - 可观测性
  - 安全

### 微服务架构原理及特征

- 微服务架构中的基本概念及组件
  - 服务、实例......

- 服务间通信
  - RPC、HTTP

- 服务注册及服务发现

### 核心服务治理功能

- 服务发布
  - 蓝绿部署
  - 灰度发布（金丝雀发布）

- 流量治理

- 负载均衡
  - Round Robin
  - Ring Hash
  - Random

- 稳定性治理
  - 限流
  - 熔断
  - 过载保护
  - 降级

### 字节跳动服务治理实践

- 请求重试的意义

- 请求重试的难点

## 课中

课程目录

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230117233539.png)

## 01.微服务架构介绍

### 1.1系统架构演变历史

> 为什么系统架构需要演进?

- 互联网的爆炸性发展
- 硬件设施的快速发展
- 需求复杂性的多样化
- 开发人员的急剧增加
- 计算机理论及技术的发展

主要是互联网行业的发展，日新月异 

- 硬件：包括 CPU MEM 存储 网络 

- 需求：文本 图片 音频 视频 VR 

- 开发人员：早期的精英程序员到如今的易于上手的开发平台 

- 计算机理论技术：算法（Paxos Raft） NoSQL 大数据

![架构演进](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118104405.png)

### 单体架构

![单体架构](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118104853.png)

|      |   优势   |        劣势        |
| :--: | :------: | :----------------: |
|  1   | 性能最高 |     debug 困难     |
|  2   |  冗余小  |    模块相互影响    |
|  3   |          | 模块分工、开发流程 |

all in one process，debug 自己写的程序就够头疼 

**痛点**： 

1. debug 困难 字节有上万研发，试想全部开发成一个程序，debug 会是什么体验？ 
2. 模块相互影响 非核心功能可能导致程序崩溃
3. 单个仓库的模块分工，依赖管理，开发流程 几乎无法分工（除了 google）

### 垂直应用架构

按照业务线垂直划分

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118115811.png)

|      |       优势       |       劣势       |
| :--: | :--------------: | :--------------: |
|  1   | 业务独立开发维护 | 不同业务存在冗余 |
|  2   |                  | 每个业务还是单体 |

**缺陷**：不同业务线的模块存在冗余，无法复用

### 分布式架构

抽出业务无关的公共模块

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118123500.png)

|      |        优势        |       劣势        |
| :--: | :----------------: | :---------------: |
|  1   | 业务无关的独立服务 | 服务模块bug可导致 |
|  2   |                    |     全站瘫痪      |
|  3   |                    |   调用关系复杂    |
|  4   |                    |   不同服务冗余    |

**缺陷**

- 一个模块服务有问题，可以导致整个系统崩溃
- 调用关系错综复杂
- 不同服务依然存在冗余

### SOA架构(Service Oriented Architecture)

面向服务，开始引入`服务注册`的概念

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118124301.png)

|      |   优势   |          劣势          |
| :--: | :------: | :--------------------: |
|  1   | 服务注册 | 整个系统设计是中心化的 |
|  2   | 单独部署 |    需要从上至下设计    |
|  3   |          |        重构困难        |

### 微服务架构

彻底的服务化

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118130110.png)

|      |     优势     |      劣势      |
| :--: | :----------: | :------------: |
|  1   |   开发效率   | 治理、运维难度 |
|  2   | 业务独立设计 |    观测挑战    |
|  3   |   自下而上   |     安全性     |
|  4   |   故障隔离   |   分布式系统   |

### 1.2微服务架构概览

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118130721.png)

### 1.3微服务架构核心要素

|      |  服务治理  | 可观测性 |   安全   |
| :--: | :--------: | :------: | :------: |
|  1   |  服务注册  | 日志采集 | 身份验证 |
|  2   |  服务发现  | 日志分析 | 认证授权 |
|  3   |  负载均衡  | 监控打点 | 访问令牌 |
|  4   |   扩缩容   | 监控大盘 |   审计   |
|  5   |  流量治理  | 异常报警 | 传输加密 |
|  6   | 稳定性治理 | 链路追踪 | 黑产攻击 |
|  7   |     …      |    …     |    …     |

微服务拆分后带来的挑战

- 日志采集
- 监控打点
- 链路追踪
- 认证授权

## 02. 微服务架构原理及特征

### 2.1基本概念

服务(service)

- 一组具有相同逻辑的运行实体。

实例(instance)

- 一个服务中，每个运行实体即为一个实例。

实例与进程的关系

- 实例与进程之间没有必然对应关系，可以一个实例可以对应一个或多个进程(反之不常见)

集群(cluster)

- 通常指服务内部的逻辑划分，包含多个实例。

常见的实例承载形式

- 进程、VM、k8s pod .....

有状态/无状态服务

- 服务的实例是否存储了可持久化的数据(例如磁盘文件)。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118131334.png)

如果把HDFS看做一组微服务

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118131656.png)



服务间通信

- 对于单体服务，不同模块通信只是简单的函数调用。
- 对于微服务，服务间通信意味着网络传输。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118131737.png)



### 2.2服务注册及发现

问题：在代码层面，如何指定调用一个目标服务的地址(ip:port) ?
hardcode?

```go
// Service A wants to call service B.
client := grpc.NewClient("10.23.45.67:8080")
```

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118132007.png)

对于网络通信功能(不管http还是rpc)，我们知道是需要指定远程的ip port的。既然服务之间存在通信关系，在实际代码层面，我们怎么写通信地址？hardcode的方式指定下游实例地址有什么问题?

回答

- 服务有多个实例，没法hardcode (记住一个服务的所有实例都是运行同一份代码)
- 服务实例ip port本身是动态变化的

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118132121.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118132143.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118132202.png)

了解了服务发现的基本机制，我们来看看如何基于服务发现来实现平滑无损的服务实例上下线流程

假设系统管理员需要下线service B的实例-3不能直接下线，因为还有流量。

**服务实例上线及下线过程**

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118132432.png)

如上图，如果是直接下掉实例3，那肯定是有问题

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118132445.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118132456.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118132508.png)

健康检测

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118132521.png)

服务注册

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118132536.png)

### 2.3流量特征

刚刚我们已经了解了微服务的基本概念和服务发现的思想，我们再从流量的视角看看微服务架构的全貌  

弱化连接的概念，强调“请求” 即同一个客户端长连接发出的请求，理论上可以到达服务中所有实例 API gateway 可以用作身份认证，进而将 token 附在请求上

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118132615.png)

## 03.核心服务治理功能

### 3.1服务发布

服务发布(deployment)， 即指让一个服务升级运行新的代码的过程。

服务发布的难点

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118153931.png)

### 蓝绿部署

项目逻辑上分为AB组，在项目系统时，首先把A组从负载均衡中摘除，进行新版本的部署。B组仍然继续提供服务。

当A组升级完毕，负载均衡重新接入A组，再把B组从负载列表中摘除，进行新版本的部署。A组重新提供服务。

最后，B组也升级完成，负载均衡重新接入B组，此时，AB组版本都已经升级完成，并且都对外提供服务。

| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118180655.png) |
| :----------------------------------------------------------: |
| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118180711.png) |
| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118180730.png) |
| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118180747.png) |
| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118180807.png) |
| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118180820.png) |

### 灰度发布(金丝雀发布)

灰度发布只升级部分服务，即让一部分用户继续用老版本，一部分用户开始用新版本，如果用户对新版本没什么意见，那么逐步扩大范围，把所有用户都迁移到新版本上面来。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118181456.png)

**部署过程**

- 从LB摘掉灰度服务器，升级成功后再加入LB；
- 少量用户流量到新版本；
- 如果灰度服务器测试成功，升级剩余服务器。

灰度发布是通过切换线上并存版本之间的路由权重，逐步从一个版本切换为另一个版本的过程。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118181440.png)

### 滚动发布

滚动发布是指每次只升级一个或多个服务，升级完成后加入生产环境，不断执行这个过程，直到集群中的全部旧版本升级新版本。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118181629.png)

- 红色：正在更新的实例
- 蓝色：更新完成并加入集群的实例
- 绿色：正在运行的实例

### 3.2流量治理

在微服务架构下，我们可以基于地区、集群、实例、请求等维度，对端到端流量的路由路径进行精确控制。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118181820.png)

### 3.3负载均衡

负载均衡(Load Balance)负责分配请求在每个下游实例上的分布。一个服务中，通常每个实例的负载应当是大体均衡一致的。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118182001.png)

常见的LB策略

- Round Robin
- Random
- Ring Hash
- Least Request

### 3.4稳定性治理

线上服务总是会出问题的，这与程序的正确性无关。

- 网络攻击
- 流量突增
- 机房断电
- 光纤被挖
- 机器故障
- 网络故障
- 机房空调故障
- …

微服务架构中典型的稳定性治理功能

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118182448.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118182515.png)

## 04.字节跳动服务治理实践

### 4.1重试的意义

**本地函数调用**

本地函数调用基本没有重试的意义

```go
func LocalFunc(x int) int {
	res := calculate(x * 2)
	return res
}
```

可能有哪些异常?

- 参数非法
- OOM (`Out Of Memory`)
- NPE (`Null Pointer Exception`)
- 边界case
- 系统崩溃
- 死循环
- 程序异常退出
- ….

**远程函数调用**

```go
func RemoteFunc(ctx context.Context, x int) (int, error) {
	ctx2, defer_func := context.WithTimeout(ctx, time.Second)
	defer defer_func()
	res, err := grpc_client.Calculate(ctx2, x*2)
	return res, err
}
```

可能有哪些异常?

- 网络抖动
- 下游负载高导致超时
- 下游机器宕机
- 本地机器负载高，调度超时
- 下游熔断、限流
- …

重试可以避免掉偶发的错误，提高SLA(Service-Level Agreement)

```go
func RemoteFunc(ctx context.Context, x int) (int, error) {
	ctx2, defer_func := context.WithTimeout(ctx, time.Second)
	defer defer_func()
	res, err := grpc_client.Calculate(ctx2, x*2)
	return res, err
}

func RemoteFuncRetry(ctx context.Context, x int) (res int, err error) {
	for i := 0; i < 3; i++ {
		if res, err = RemoteFunc(ctx, x); err == nil {
			return
		}
	}
	return
}
```

**重试的意义**

- 降低错误率
  - 假设单次请求的错误概率为0.01，那么连续两次错误概率则为0.0001。
- 降低长尾延时
  - 对于偶尔耗时较长的请求，重试请求有机会提前返回。
- 容忍暂时性错误
  - 某些时候系统会有暂时性异常(例如网络抖动)，重试可以尽量规避。
- 避开下游故障实例
  - 一个服务中可能会有少量实例故障(例如机器故障)，重试其他实例可以成功。

### 4.2重试的难点

既然重试这么多好处，为什么 默认不用呢? 

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118184918.png)

1. 幂等性：多次请求可能会造成数据不一致
2. 重试风暴：随着调用深度的增加，重试次数会指数级上涨(稍后分析)
3. 超时设置：假设一个调用正常是1s的超时时间，如果允许一次重试，那么第一次请求经过多少时间时，才开始重试呢?



**重试风暴**

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118204146.png)



### 4.3重试策略

**限制重试比例**

设定一个重试比例阈值(例如1%)，重试次数占所有请求比例不超过该阈值。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118204318.png)

重试只有在大部分请求都成功，只有少量请求失败时，才有必要 如果大部分请求都失败，重试只会加剧问题严重性 因此，可以定义，比如重试次数不能超过正常成功请求次数的 1%

**防止链路重试**

链路层面的防重试风暴的核心是限制每层都发生重试，理想情况下只有最下一层发生重试。可以返回特殊的status表明“请求失败，但别重试”。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118204425.png)

**Hedged requests**

对于可能超时(或延时高)的请求，重新向另一个下游实例发送一个相同的请求，并等待先到达的响应。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118204502.png)

### 4.4重试效果验证

实际验证经过上述重试策略后，在链路上发生的重试放大效应。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118204933.png)



## 课后

1. 结合 CAP 等原理，思考微服务架构有哪些缺陷？
   - 由于在运行上具有分布式的特征，服务调用依赖网络通信能力，如果出现网络故障，就可能出现数据不一致或不可用的情况，而且这种方式会增加大量内网流量，对网络运维和设备的要求比较高。
2. 微服务是否拆分得越“微”越好？为什么？
   -  微服务的互相调用依赖网络通信，拆的越细越容易在网络通信上出现性能瓶颈。其次是拆的太细太多，服务运维就非常困难。而且环节越多，越容易产生意想不到的放大效应。
3. Service Mesh 这一架构是为了解决微服务架构的什么问题？
   - 主要是解决微服务架构中服务间可靠调用、服务治理等问题。
4. 有没有可能有这样一种架构，从开发上线运维体验上是微服务，但实际运行又类似单体服务？
   - 在部署时考虑把耦合性很强的微服务成组部署，相互调用的通信走回环网卡流量，但是又会导致单体服务的易连锁宕机的问题。

参考链接

1. 青训营官方账号：https://juejin.cn/post/7099665398655615006#heading-32
2. https://bytedance.feishu.cn/file/boxcnPjF5oJxpZh4ZQYDwVCSxib
2. http://www.lanxinbase.com/?p=2553
2. https://juejin.cn/post/6960282794081812511
2. https://link.juejin.cn/?target=https%3A%2F%2Fwww.infoq.cn%2Farticle%2Fasgjevrm8islszo7ixzh
2. https://medium.com/swlh/a-design-analysis-of-cloud-based-microservices-architecture-at-netflix-98836b2da45f

















