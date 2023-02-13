---
title: day15-走进消息队列
date: 2023-01-15 13:01:19
tags:
- Go学习路线
- 字节跳动青训营
---

<img src="https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118235050.jpg" style="zoom:200%;" />

这是我参与「第三届青训营 - 后端场」笔记创作活动的的第15篇笔记。*PC端阅读效果更佳，点击文末：**阅读原文**即可。*

## 「走进消息队列」 第三届字节跳动青训营 - 后端专场

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230117222718.png)

[TOC]

### 课程目录

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118235449.png)

### 背景1

看下面的这个的场景，有一天晚上晚上我们上完课，回到宿舍，想着新出的游戏机，但又摸了摸钱包，太贵了买不起，这个时候你突然想到，今天抖音直播搞活动，瞬间你掏出了手机打开抖音搜索，找到直播间以后，你点开了心心念念的游戏机详情页，看到价格只要500。

这个时候我们分析一下，就我们上面这几步操作，在我们的程序背后，做了什么事情？

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118235712.png)

## 案例一:系统崩溃

首先，请求会先到搜索商品这个服务上，并记录下你的搜索行为，然后点击商品的时候，又记录了我们的点击商品，这些数据最终都会通过计算分析，目的是为了下一次给你更准确的信息。

这个时候问题来了，如果这个时候，负责记录存储的数据库被一个小哥删库跑路了。我们的所有操作都动不了了，这个时候我们应该怎么办，带着这个问题，我们继续往下看

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230118235910.png)

**案例一:解决方案**

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119000738.png)

### 背景2

双十一，某宝的商品价格非常低，看到这个价格，你非常心动，定睛一看，商品即将在3分钟后开抢，这个价格必须要抢到啊！

但此时在无数台手机的后面，藏着无数和你一样饥渴的同学，时间快到了，心里面想着赶紧抢，我们再来看看，后面的程序又做了哪些事情呢？

## 案例二:服务能力有限

可以看到，一堆人都发起了订单请求，可是公司给的预算不够，服务器的配置太弱，订单服务只能同事处理10个订单请求。

这个时候我们又该怎么办呢。继续往下看

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119000259.png)

**案例二:解决方案**

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119000811.png)

## 

### 背景3

在我们点击提交订单之后，这个怎么一直转圈圈，卡在这个页面啊，等了半分钟后，啊终于抢到了，不过这个app也太慢了，下次不用了。

我们进一步看一下这次问题出在哪里了

## 案例三:链路耗时长尾

一通分析，发现库存服务和订单服务都挺快的，但是最后通知商家这一步咋这么慢，是不是还可以进行优化呀？

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119000440.png)

**案例三:解决方案**

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119000839.png)

## 案例四:日志存储

在大家都抢到了自己心仪商品准备睡去的时候，在字节跳动的会议室里传出了悲伤的声音，因为刚刚有服务器坏掉了，我们的本地日志都丢掉了，没有了日志，我们还怎么去修复那些刚刚出现的那些问题，周围一片寂静，突然**小张站出来缓缓的说了一句话**，众人才露出了微笑准备下班离开，大家能猜到小张到底说的什么吗

**思考时间**

四个场景，如何解决?

1. 系统奔溃
2. 服务处理能力有限
3. 链路耗时长尾
4. 日志如何处理

**案例四:解决方案**

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119000903.png)

> 什么是消息队列?

- 消息队列(MQ)，指保存消息的一个容器，本质是个队列。但这个队列呢，需要支持**高吞吐，高并发，并且高可用**。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119000951.png)

# 01. 前世今生

**消息队列发展历程**

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119190211.png)

消息中间件其实诞生的很早，早在1983年互联网应用还是一片荒芜的年代，有个在美国的印度小哥Vivek就设想了一种通用软件总线，世界上第一个现代消息队列软件The Information Bus(TIB)，他 TIB受到了企业的欢迎，这家公司的业务发展引起了当时最牛气的IT公司IBM的注意，于是他们一开始研发了自己消息队列软件，于是才有了后来的wesphere mq，再后来微软也加入了战团。

接近2000年的时候，互联网时代已经初见曙光，全球的应用程序得到了极大地丰富，对于程序之间互联互通的需求越来越强烈，但是各大IT公司之间还是牢牢建立着各种技术壁垒，以此来保证自己的商业利益，所以消息中间件在那个时候是大型企业才能够用的起的高级玩意。 

但是时代的洪流不可逆转，有壁垒就有打破壁垒的后来者，2001年sun发布了jms技术，试图在各大厂商的层面上再包装一层统一的java规范。java程序只需要针对jms api编程就可以了，不需要再关注使用了什么样的消息中间件，但是jms仅仅适用于java。

2004年AMQP（高级消息队列协议）诞生了，才是真正促进了消息队列的繁荣发展，任何人都可以针对AMQP的标准进行编码。有好的协议指导，再加上互联网分布式应用的迅猛发展成为了消息中间件一飞冲天的最大动力，程序应用的互联互通，发布订阅，最大契合了消息中间件的最初的设计初衷。除了刚才介绍过的收费中间件，后来开源消息中间件开始层出不穷。

常见比较流行的有ActiveMQ、RabbitMQ 、Kafak、阿里的RocketMQ，以及目前存算分离的Pulsar，在目前互联网应用中消息队列中间件基本上成为标配。



**业界消息队列对比**

到目前为止，比较流行的MQ是以下几个：

- `Kafka`：分布式的、分区的、多副本的日志提交服务，在高吞吐场景下发挥较为出色

- `RocketMQ`：低延迟、强一致、高性能、高可靠、万亿级容量和灵活的可扩展性，在一些实时场景中运用较广

- `Pulsar`：是下一代云原生分布式消息流平台，集消息、存储、轻量化函数式计算为一体、采用存算分离的架构设计

- `BMQ`：和Pulsar架构类似，存算分离，初期定位是承接高吞吐的离线业务场景，逐步替换掉对应的Kafka集群

# 02.消息队列- Kafka

## 2.1使用场景

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119191752.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119192152.png)

1. 第一步：首先需要创建一个Kafka集群，但如果你是在字节工作，恭喜你这一步消息团队的小伙伴已经帮你完成了 
2. 第二步：需要在这个集群中创建一个Topic，并且设置好分片数量 
3. 第三步：引入对应语言的SDK，配置好集群和Topic等参数，初始化一个生产者，调用Send方法，将你的Hello World发送出去 

4. 第四步：引入对应语言的SDK，配置好集群和Topic等参数，初始化一个消费者，调用Poll方法，你将收到你刚刚发送的Hello World  

## 2.2基本概念

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119210552.png)

- Topic：Kakfa中的逻辑队列，可以理解成每一个不同的业务场景就是一个不同的topic，对于这个业务来说，所有的数据都存储在这个topic中 
- Cluster：Kafka的物理集群，每个集群中可以新建多个不同的topic 
- Producer：顾名思义，也就是消息的生产端，负责将业务消息发送到Topic当中 
- Consumer：消息的消费端，负责消费已经发送到topic中的消息 
- Partition：通常topic会有多个分片，不同分片直接消息是可以并发来处理的，这样提高单个Topic的吞吐

## 2.3Offset

Offset：消息在partition内的相对位置信息，可以理解为唯一ID, 在partition内部严格递增。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119210922.png)

对于每一个Partition来说，每一条消息都有一个唯一的Offset，消息在partition内的相对位置信息，并且严格递增

### Replica

每个分片有多个Replica，Leader Replica 将会从ISR中选出。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119211144.png)

`Replica`：分片的副本，分布在不同的机器上，可用来容灾，Leader对外服务，Follower异步去拉取leader的数据进行一个同步，如果leader挂掉了，可以将Follower提升成leader再堆外进行服务 

`ISR`：意思是同步中的副本，对于Follower来说，始终和leader是有一定差距的，但当这个差距比较小的时候，我们就可以将这个follower副本加入到ISR中，不在ISR中的副本是不允许提升成Leader的

## 2.4数据复制

下面这幅图代表着Kafka中副本的分布图。

图中Broker代表每一个Kafka的节点，所有的Broker节点最终组成了一个集群。整个图表示，图中整个集群，包含了4个Broker机器节点，集群有两个Topic，分别是Topic1和Topic2，Topic1有两个分片，Topic2有1个分片，每个分片都是三副本的状态。这里中间有一个Broker同时也扮演了Controller的角色，Controller是整个集群的大脑，负责对副本和Broker进行分配。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119211318.png)

## 2.5 Kafka架构

而在集群的基础上，还有一个模块是`ZooKeeper`，这个模块其实是存储了集群的元数据信息，比如副本的分配信息等等，Controller计算好的方案都会放到这个地方

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119211452.png)

## 2.6一条消息的自述

从一条消息的视角，看看为什么Kafka 能支撑这么高的吞吐?

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119213710.png)

> 思考：如果发送条消息， 等到其成功后再发一条会有什么问题？

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119213807.png)

## 2.7 Producer-批量发送

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119213945.png)

**数据压缩**

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119214021.png)



## 2.8 Broker-数据的存储

如何写入到磁盘呢，我们先来看一下Kafka最终存储的文件结构是什么样子的？

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119214109.png)

### Broker消息文件结构

在每一个Broker，都分布着不同Topic的不同分片

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119214206.png)

### 磁盘结构

移动磁头找到对应磁道，磁盘转动，找到对应扇区，最后写入。寻道成本比较高，因此顺序写可以减少寻道所带来的时间成本。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119214358.png)

### Broker-顺序写

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119214533.png)

采用顺序写的方式进行写入，以提高写入效率

### Broker-如何找到消息

Consumer通过发送FetchRequest请求消息数据，Broker 会将指定Offset处的消息,按照时间窗口和消息

大小窗口发送给Consumer，寻找数据这个细节是如何做到的呢?

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119214638.png)

此时我们的消息写入到Broker的磁盘上了，那这些数据又该怎么被找到然后用来消费呢

### Broker偏移量索引文件

介绍文件：文件名是文件中第一条消息的offset 然后，第一步，通过二分找到小于目标文件的最大文件

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119220154.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119221316.png)

通过二分找到小于目标offset最大的索引位置，再遍历找到目标offset

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119221354.png)

如果我们需要使用时间戳来寻找的时候，和offset相比只是多加了以及索引，也就是通过二分找到时间戳对应的offset，再重复之前的步骤找到相应的文件数据

### Broker-传统数据拷贝

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119221447.png)

### Broker-零拷贝

Consumer从Broker中读取数据，通过sendfile的方式，将磁盘读到os内核缓冲区后，直接转到socket buffer进行网络发送 Producer生产的数据持久化到broker，采用mmap文件映射，实现顺序的快速写入

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119221825.png)

## 2.9 Consumer-消息的接收端

对于一个Consumer Group来说，多个分片可以并发的消费，这样可以大大提高消费的效率，但需要解决的问题是，Consumer和Partition的分配问题，也就是对于每一个Partition来讲，该由哪一个Consumer来消费的问题。

对于这个问题，我们一般有两种解决方法，**手动分配**和**自动分配**

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119221931.png)



### Consumer- Low Level

通过手动进行分配，哪一个Consumer消费哪一个Partition完全由业务来决定。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119222557.png)

第一，手动分配，也就是Kafka中所说的Low Level消费方式进行消费，这种分配方式的一个好处就是启动比较快，因为对于每一个Consumer来说，启动的时候就已经知道了自己应该去消费哪个消费方式，就好比图中的Consumer Group1来说，Consumer1去消费Partition1,2,3     Consumer2，去消费456，    Consumer3去消费78。

这些Consumer再启动的时候就已经知道分配方案了，但这样这种方式的缺点又是什么呢，想象一下，如果我们的Consumer3挂掉了，我们的7,8分片是不是就停止消费了。又或者，如果我们新增了一台Consumer4，那是不是又需要停掉整个集群，重新修改配置再上线，保证Consumer4也可以消费数据，其实上面两个问题，有时候对于线上业务来说是致命的。



### Consumer-High Level

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119223022.png)

所以Kafka也提供了自动分配的方式，这里也叫做High Level的消费方式。

简单的来说，就是在我们的Broker集群中，对于不同的Consumer Group来讲，都会选取一台Broker当做Coordinator，而Coordinator的作用就是帮助Consumer Group进行分片的分配，也叫做分片的rebalance，使用这种方式，如果ConsumerGroup中有发生宕机，或者有新的Consumer加入，整个partition和Consumer都会重新进行分配来达到一个稳定的消费状态

## 2.10 Consumer Rebalance







![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119223132.png)

刚刚总共讲了哪一些可以帮助Kafka提高吞吐或者稳定性的功能?

- Producer：批量发送、数据压缩

- Broker：顺序写，消息索引，零拷贝

- Consumer： Rebalance

## 2.11 Kafka-数据复制问题

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119223240.png)

通过前面的介绍我们可以知道，对于Kafka来说，每一个Broker上都有不同topic分区的不同副本，而每一个副本，会将其数据存储到该Kafka节点上面，对于不同的节点之间，通过副本直接的数据复制，来保证数据的最终一致性，与集群的高可用。

## 2.12 Kafka-重启操作

举个例子来说，如果我们对一个机器进行重启首先，我们会关闭一个Broker，此时如果该Broker上存在副本的Leader，那么该副本将发生leader切换，切换到其他节点上面并且在ISR中的Follower副本，可以看到图中是切换到了第二个Broker上面 

而此时，因为数据在不断的写入，对于刚刚关闭重启的Broker来说，和新Leader之间一定会存在数据的滞后，此时这个Broker会追赶数据，重新加入到ISR当中 

当数据追赶完成之后，我们需要回切leader，这一步叫做prefer leader，这一步的目的是为了避免，在一个集群长期运行后，所有的leader都分布在少数节点上，导致数据的不均衡 通过上面的一个流程分析，我们可以发现对于一个Broker的重启来说，需要进行数据复制，所以时间成本会比较大，比如一个节点重启需要10分钟，一个集群有1000个节点，如果该集群需要重启升级，则需要10000分钟，那差不多就是一个星期，这样的时间成本是非常大的。 有同学可能会说，老师可以不可以并发多台重启呀，问的好，不可以。为什么呢，在一个两副本的集群中，重启了两台机器，对某一分片来讲，可能两个分片都在这台机器上面，则会导致该集群处于不可用的状态。这是更不能接受的。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119224549.png)



## 2.13 Kafka-替换、扩容、缩容

如果是替换，或者扩容，或者缩容操作呢，我们来看看。 

如果是替换，和刚刚的重启有什么区别，其实替换，本质上来讲就是一个需要追更多数据的重启操作，因为正常重启只需要追一小部分，而替换，则是需要复制整个leader的数据，时间会更长 扩容呢，当分片分配到新的机器上以后，也是相当于要从0开始复制一些新的副本 而缩容，缩容节点上面的分片也会分片到集群中剩余节点上面，分配过去的副本也会从0开始去复制数据 

以上三个操作均有数据复制所带来的时间成本问题，所以对于Kafka来说，运维操作所带来的时间成本是不容忽视的

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119224809.png)

## 2.14 Kafka-负载不均衡

这个场景当中，同一个Topic有4个分片，两副本，可以看到，对于分片1来说，数据量是明显比其他分片要大的，当我们机器IO达到瓶颈的时候，可能就需要把第一台Broker上面的Partition3迁移到其他负载小的Broker上面，接着往下看

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119224948.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119225030.png)

## 02. Kafka-问题总结

1. 运维成本高
2. 对于负载不均衡的场景，解决方案复杂
3. 没有自己的缓存，完全依赖`Page Cache`
4. `Controller`和`Coordinator`和`Broker`在同一进程中，大量IO会造成其性能下降

我们对以上两个问题进行总结：

1. 第一，因为有数据复制的问题，所以`Kafka`运维的时间成本和人力人本都不低。

2. 第二，对于负载不均衡的场景，我们需要有一个较为复杂的解决方案进行数据迁移，从而来权衡IO升高的问题 除了以上两个问题以外，`Kafka`自身还存在其他的问题 比如，`Kafka`没有自己的缓存，在进行数据读取的时候，只有Page Cache可以用，所以不是很灵活 另外在前面的介绍当中，相信大家也了解到了，`Kafka`的`Controller`和`Coordinator`都是和`Broker`部署在一起的，`Broker`因为承载大量`IO`的原因，会导致`Controller`和`Coordinator`的性能下降，如果到一定程度，可能会影响整个集群的可用性

# 03.消息队列-BMQ

## 3.1 BMQ简介

兼容Kafka协议，存算分云原生消息队列。

BMQ兼容Kafka协议，存算分离，云原生消息队列，初期定位是承接高吞吐的离线业务场景，逐步替换掉对应的Kafka集群，我们来了解一下BMQ的架构特点

### BMQ介绍

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119225805.png)

Producer -> Consumer -> Proxy -> Broker -> HDFS -> Controller -> Coordinator -> Meta 着重强调一下Proxy和Broker无状态，为下面运维比较做铺垫

这里简单介绍一下存算分离，适配Kafka协议，为什么不选择Pulsar的原因

## 3.2运维操作对比

实际上对于所有节点变更的操作，都仅仅只是集群元数据的变化，通常情况下都能秒级完成，而真正的数据已经移到下层分布式文件存储去了，所以运维操作不需要额外关心数据复制所带来的时间成本

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119225920.png)

## 3.3 HDFS写文件流程

通过前面的介绍，我们知道了，同一个副本是由多个segment组成，我们来看看BMQ对于单个文件写入的机制是怎么样的？

首先客户端写入前会选择一定数量的DataNode，这个数量是副本数，然后将一个文件写入到这三个节点上，切换到下一个segment之后，又会重新选择三个节点进行写入。

这样一来，对于单个副本的所有segment来讲，会随机的分配到分布式文件系统的整个集群中。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119230125.png)

## 3.4 BMQ文件结构

对于Kafka分片数据的写入，是通过先在Leader上面写好文件，然后同步到Follower上，所以对于同一个副本的所有Segment都在同一台机器上面。就会存在之前我们所说到的单分片过大导致负载不均衡的问题，但在BMQ集群中，因为对于单个副本来讲，是随机分配到不同的节点上面的，因此不会存在Kafka的负载不均问题。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119230324.png)

## 3.5 Broker-Partition 状态机

其实对于写入的逻辑来说，我们还有一个状态机的机制，用来保证不会出现同一个分片在两个Broker上同时启动的情况，另外也能够保证一个分片的正常运行。

首先，Controller做好分片的分配之后，如果在该Broker分配到了Broker，首先会start这个分片，然后进入Recover状态，这个状态主要有两个目的获取分片写入权利，也就是说，对于hdfs来讲，只会允许我一个分片进行写入，只有拿到这个权利的分片我才能写入，

第二一个目的是如果上次分片是异常中断的，没有进行save checkpoint，这里会重新进行一次save checkpoint，然后就进入了正常的写流程状态，创建文件，写入数据，到一定大小之后又开始建立新的文件进行写入。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119230710.png)

### Broker-写文件流程

数据校验：CRC , 参数是否合法 校验完成后，会把数据放入Buffer中 通过一个异步的Write Thread线程将数据最终写入到底层的存储系统当中 这里有一个地方需要注意一下，就是对于业务的写入来说，可以配置返回方式，可以在写完缓存之后直接返回，另外我也可以数据真正写入存储系统后再返回，对于这两个来说前者损失了数据的可靠性，带来了吞吐性能的优势，因为只写入内存是比较快的，但如果在下一次flush前发生宕机了，这个时候数据就有可能丢失了，后者的话，因为数据已经写入了存储系统，这个时候也不需要担心数据丢失，相应的来说吞吐就会小一些 我们再来看看Thread的具体逻辑，首先会将Buffer中的数据取出来，调用底层写入逻辑，在一定的时间周期上去flush，flush完成后开始建立Index，也就是offset和timestamp对于消息具体位置的映射关系 Index建立好以后，会save一次checkpoint，也就表示，checkpoint后的数据是可以被消费的辣，我们想一下，如果没有checkpoint的情况下会发生什么问题，如果flush完成之后宕机，index还没有建立，这个数据是不应该被消费的 最后当文件到达一定大小之后，需要建立一个新的segment文件来写入

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119230914.png)

### Broker-写文件Failover

我们之前说到了，建立一个新的文件，会随机挑选与副本数量相当的数据节点进行写入，那如果此时我们挑选节点中有一个出现了问题，导致不能正常写入了，我们应该怎么处理，是需要在这里等着这个节点恢复吗，当然不行，谁知道这个节点什么恢复，既然你不行，那就把你换了，可以重新找正常的节点创建新的文件进行写入，这样也就保证了我们写入的可用性

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119231017.png)

## 3.6 Proxy

首先Consumer发送一个Fetch Request，然后会有一个Wait流程，那么他的左右是什么呢，想象一个Topic，如果一直没有数据写入，那么，此时consumer就会一直发送Fetch Request，如果Consumer数量过多，BMQ的server端是扛不住这个请求的，因此，我们设置了一个等待机制，如果没有fetch到指定大小的数据，那么proxy会等待一定的时间，再返回给用户侧，这样也就降低了fetch请求的IO次数，经过我们的wait流程后，我们会到我们的Cache里面去找到是否有存在我们想要的数据，如果有直接返回，如果没有，再开始去存储系统当中寻找，首先会Open这个文件，然后通过Index找到数据所在的具体位置，从这个位置开始读取数据

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119231201.png)

## 3.7多机房部署

为什么需要多机房部署，其实对于一个高可用的服务，除了要防止单机故障所带来的的影响意外，也要防止机房级故障所带来的影响，比如机房断点，机房之间网络故障等等。

那我们来看看BMQ的多机房部署是怎么做的Proxy -> Broker -> Meta -> HDFS

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119231251.png)

## 3.8 BMQ-高级特性

泳道 -> Databus -> Mirror -> Index -> Parquet

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119231442.png)

## 3.9泳道消息

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119231902.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119231932.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119231954.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119232017.png)



![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119232334.png)

## 3.10 Databus

直接使用原生SDK会有什么问题？

1. 客户端配置较为复杂
2. 不支持动态配置，更改配置需要停掉服务
3. 对于latency不是很敏感的业务，batch 效果不佳

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119232445.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119232528.png)

1. 简化消息队列客户端复杂度
2. 解耦业务与Topic
3. 缓解集群压力，提高吞吐

## 3.11 Mirror

思考一下，我们是否可以通过多机房部署的方式，解决跨`Region`读写的问题?

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119232709.png)

使用`Mirror`通过最终一致的方式， 解决跨`Region`读写问题。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119232749.png)

## 3.12 Index

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119232851.png)

## 3.13 Index

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119232923.png)

直接在BMQ中将数据结构化，配置索引DDL，异步构建索引后，通过Index Query服务读出数据。

## 3.14 Parquet

Apache Parquet是Hadoop生态圈中一种**新型列式存储格式**，它可以兼容Hadoop生态圈中大多数计算框架(Hadoop、Spark等)， 被多种查询引擎支持(Hive、Impala、 Dill等)。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119233153.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119233206.png)

## 03.小结

1. BMQ的架构模型(解决Kafka存在的问题)
2. BMQ读写流程(Failover 机制，写入状态机)
3. BMQ高级特性(泳道、Databus、 Mirror、 Index、 Parquet)

# 04消息队列-RocketMQ

例如，针对电商业务线，其业务涉及广泛，如注册、订单、库存、物流等；同时，也会涉及许多业务峰值时刻，如秒杀活动、周年庆、定期特惠等

## 4.1 RocketMQ基本概念

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119233413.png)

根据我们刚刚的介绍，可以看到Producer，Consumer，Broker这三个部分，Kafka和RocketMQ是一样的，而Kafka中的Partition概念在这里叫做ConsumerQueue

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119233515.png)

## 4.2 RocketMQ架构

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119233613.png)

先说数据流也是通过`Producer`发送给`Broker`集群，再由`Consumer`进行消费 `Broker`节点有`Master`和`Slave`的概念 `NameServer`为集群提供轻量级服务发现和路由。

## 4.3存储模型

接下来我们来看看RocketMQ消息的存储模型，对于一个Broker来说所有的消息的会append到一个CommitLog上面，然后按照不同的Queue，重新Dispatch到不同的Consumer中，这样Consumer就可以按照Queue进行拉取消费。

但需要注意的是，这里的ConsumerQueue所存储的并不是真实的数据，真实的数据其实只存在CommitLog中，这里存的仅仅是这个Queue所有消息在CommitLog上面的位置，相当于是这个Queue的一个密集索引

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119233809.png)

## 4.4 RocketMQ-高级特性

RocketMQ的高级特性有哪些？

那RocketMQ所提供的高级特性有哪些呢，我们通过一些我们有可能在业务场景中遇到的问题，一起来看一下

### 事务场景

先看一下我们最开始说的这个场景，正常情况下，这个下单的流程应该是这个样子。

首先我保证库存足够能够顺利-1，这个时候再消息队列让我其他系统来处理，比如订单系统和商家系统，但这里有个比较重要的点，我库存服务和消息队列必须要是在同一个事务内的，大家还记不记得事务的基本特性是什么。ACID，这里库存记录和往消息队列里面发的消息这两个事情，是需要有事务保证的，这样不至于发生，库存已经-1了，但我的订单没有增加，或者商家也没有收到通知要发货。

因此RocketMQ提供事务消息来保证类似的场景，我们来看看其原理是怎么样的

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119234002.png)

### 事务消息

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119234104.png)

### 延迟发送

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119234200.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119234232.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119234247.png)

### 处理失败

该如何处理失败的消息呢？

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119234314.png)

### 消费重试和死信队列

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230119234419.png)

## 04.小结

1. RocketMQ的基本概念(Queue, Tag)
2. RocketMQ的底层原理(架构模型、存储模型)
3. RocketMQ的高级特性(事务消息、重试和死信队列， 延迟队列)

# 课程总结

- 前世今生：消息队列发展历程
- Kafka：基本概念、架构设计、底层原理、架构缺点
- BMQ：架构设计、底层原理、Kafka比较、高级特性
- RocketMQ：架构设计、底层原理、高级特性

参考文献：

1. https://bytedance.feishu.cn/file/boxcnKXjTxtmdCpdidtJZGckf5b
2. https://juejin.cn/post/7100051825939709983#heading-0





















































































































































































































































































