---
title: day16-分布式定时任务那些事儿
date: 2023-01-15 12:59:16
tags:
- Go学习路线
- 字节跳动青训营
---







<img src="https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230120095422.jpg" style="zoom:200%;" />

这是我参与「第三届青训营 - 后端场」笔记创作活动的的第16篇笔记。*PC端阅读效果更佳，点击文末：**阅读原文**即可。*

## 「分布式定时任务那些事儿」 第三届字节跳动青训营 - 后端专场

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230120002153.png)

[TOC]

## 课前

### 分布式定时任务发展历史

- Linux命令-CronJob

- 单机定时任务-Timer、Ticker

- 单机定时任务-ScheduledExecutorService

- 任务调度- Quartz

- 分布式定时任务

### 分布式定时任务核心架构

- 控制台Admin

- 触发器Trigger

- 调度器Scheduler

- 执行器Executor

### 知识点扩充

- 时间轮

- 延时消息

- 离线计算引擎 Hive

- 实时计算引擎 Flink

## 课中

## 这篇课程我能学到什么

知识面扩充

- 对分布式定时任务建立起宏观的认知，并深入了解其实现原理
- 了解关联的单机定时任务、大数据处理引擎，通过了解不同实现方案的优劣来拓展知识面

项目实践能力加强

- 了解在哪些实际业务场景中使用分布式定时任务
- 对于实际业务场景的中间件选型、技术方案设计做到成竹在胸

主要有知识的**广度**、**深度**还有**项目实践**能力三大的部分。

首先是知识的广度这个课程可以帮助同学们对于分布式定时任务建立起比较宏观的一个认知。

另外除了分布式定时任务之外，也会去讲述它相关联的像单机定时任务以及它数据处理引擎，从而去扩充同学们的一个技术和知识面。

在知识深度方面的话。所以这门课程会深入去了解分布式定时任务的一系列的实现原理。 然后在这个过程当中，我们会深入地讲述不同的公司在同样一个技术问题上中的技术选型，通过去了解这种不同的实现方案来拓展知识的一个深度。

最后在项目实践能力方面，可以在这个过程当中了解到未来在哪些实际的业务场景中可以去使用分布式定时任务，并且后续如果遇到类似的业务场景，那对于像中间件的一个选型技术方案的设计能够做得承竹在胸。到底是该用这个分支定时任务还是用单机定时任务？还是用离线技术？离线处理不还是用大数据处理引擎对自己的一些见解。

## 课程目录

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230120151636.png)

# 01.前言

## 01.春节集卡瓜分20亿

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230120152002.png)

作为后端开发同学，怎么设计最终开奖环节技术方案?

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230120152045.png)



![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230120152112.png)

# 02.发展历程

## 2.1 Windows批处理

Case 1: 10分钟后Windows电脑自动关机

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230120152232.png)

1. Step1：桌面空白处右键单击新建文本文档
2. Step2：更改文件名和后缀为"自动关机.bat"
3. Step3：修改文件内容为”shutdown -S -t 600"，代表10分钟后关机
4. Step4：双击运行该批处理文件，电脑将会在10分钟之后自动关机

## 2.2 Windows任务计划程序

Case 2：每天12:00自动疫情打卡

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230120152434.png)

## 2.3 Linux命令-CronJob

Case 3：每天02:30定时清理机器日志

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230120152734.png)

## 2.4单机定时任务-Timer、Ticker

Case 4：每隔5分钟定时刷新本地缓存数据



Java

```java
    public static void main(String[] args) {
        Timer timer = new Timer();
        timer.schedule(new TimerTask() {
            @Override
            public void run() {
                SyncLocalCache();
            }
        },5000,5*60*1000);
    }
```

Golang

```go
package main

import "time"

func main() {
	ticker := time.NewTicker(5 * time.Minute)
	for {
		select {
		case <-ticker.C:
			SyncLocalCache()
		}
	}
}
```

### 单机定时任务-ScheduledExecutorService

Case 5：每隔5分钟定时执行多个任务

```go
public class Code_ScheduledExecutorService {
    public static ScheduledExecutorService scheduler;

    public static void main(String[] args) {
        scheduler = Executors.newScheduledThreadPool(5);
        scheduler.scheduleAtFixedRate(((
                        new Runnable() {
                            @Override
                            public void run() {
                                DoSomething();
                            }
                        })),
                TimeUnit.SECONDS);
    }
}
```

## 任务调度- Quartz

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121095330.png)

## 2.6分布式定时任务

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121095415.png)

什么是分布式定时任务?

- 定时任务是指系统为了自动完成特定任务，**实时、延时、周期性**完成任务调度的过程。

- **分布式定时任务**是把分散的、可靠性差的定时任务纳入统一的平台， 并实现集群管理调度和分布式部署的一种定时任务的管理方式。

按触发时机分类:

- 定时任务：特定时间触发，比如今天15:06执行
- 延时任务：延时触发，比如10s后执行
- 周期任务：固定周期时间，或固定频率周期调度触发，比如每隔5s或者每天12点执行

### 分布式定时任务特点

- 自动化：全自动完成定时任务的调度和执行
- 平台化：基于平台化的思维管控一系列的分布式定时任务
- 分布式：在分布式系统环境下运行任务调度，突破单机定时任务的性能瓶颈
- 伸缩性：采用集群方式部署，可以随时按需扩缩容
- 高可用：单点故障不影响最终任务结果，可以做到故障转移

### 执行方式

- 单机任务：随机触发一台机器执行任务， 适用于计算量小、并发度低的任务

- 广播任务：广播到所有机器上执行同一一个任务，比如所有机器一起清理日志

- Map任务：一个任务可以分出多个子任务，每个子任务负责-部分的计算。 适用于计算量大，单机无法满足要求的任务

- MapReduce任务：在Map任务的基础上，还可以对所有子任务的结果做汇总计算，适用于计算量大，并且需要对子任务结果做汇总的任务



### 执行方式VS春节集卡

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121102808.png)

### 业内定时任务框架

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121103040.png)

## 2.7业内定时任务框架Xxl-job

- Xxl-job是大众点评员工许雪里于2015年发布的分布式任务调度平台，是一个轻量级分布式任务调度框架，其核心设计目标是开发迅速、学习简单、轻量级、易扩展。XXL-JOB支持分片，简单支持任务依赖，支持子任务依赖，不是跨平台的。

- Xxl-job很大一个优势在于开源且免费，并且轻量级，开箱即用，操作简易，上手快，企业维护起来成本不高，因而在中小型公司使用非常广泛。

### SchedulerX

分布式任务调度SchedulerX 2.0是阿里巴巴基于Akka架构自研的新一代

分布式任务调度平台，提供定时调度、调度任务编排和分布式批量处理等功能。

SchedulerX可在阿里云付费使用。它功能非常强大，在阿里巴巴内部广泛使用并久经考验。

### TCT

分布式任务调度服务(Tencent Cloud Task)是腾讯云自主研发的一款高性能、高可靠通用的分布式任务调度中间件，通过指定时间规则严格触发调度任务，保障调度任务的可靠有序执行。该服务支持国际通用的时间表达式、调度任务执行生命周期管理，解决传统定时调度任务单点及并发性能问题。同时，支持任务分片、流程编排复杂调度任务处理能力，覆盖广泛的任务调度
应用场景。

TCT仅在腾讯内部使用，未开源，也未商用。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121105713.png)

# 03.实现原理

## 3.1核心架构

分布式定时任务核心要解决**触发、调度、执行**三个关键问题

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121105830.png)

### 3.1.1数据流

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121110027.png)

## 3.1.2功能架构

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121110917.png)



## 3.2 控制台

### 3.2.1基本概念

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121111101.png)

任务: job, 任务元数据

任务实例: JobInstance，周期任务会生成多个任务实例

任务结果: JobResult，任务实例运行的结果

任务历史: JobHistory，用户可以修改任务信息，任务实例对应的任务元数据可以不同，因而使用任务历史存储

### 3.2.2基本概念-任务元数据

任务元数据(Job) 是用户对任务属性定义，包括任务类型调度时机、执行行为等。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121111329.png)

### 3.2.3基本概念任务实例

任务实例(JobInstance) 是一个确定的Job的一次运行实例。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121111434.png)

## 3.3触发器

### 3.3.1触发器核心职责

核心职责

- 给定一系列任务，解析它们的触发规则，在规定的时间点触发任务的调度

设计约束

- 需支持大量任务
- 需支持秒级的调度
- 周期任务需要多次执行
- 需保证秒级扫描的高性能，并避免资源浪费

### 3.3.2触发器方案

定期扫描 + 延时消息(腾讯、字节方案)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121111818.png)

### 3.3.3触发器-方案2

时间轮(Quartz所用方案)

时间轮是一种高效利用线程资源进行批量化调度的一种调度模型。时间轮是一个存储环形队列， 底层采用数组实现，数组中的每个元素可以存放一个定时任务列表。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121112131.png)

### 3.3.3触发器方案2

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121112424.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121112521.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121112553.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121112622.png)

### 3.3.4触发器-高可用

核心问题

- 不同业务之间，任务的调度相互影响怎么办?
- 负责扫描和触发的机器挂了怎么办?

解法思路

- 存储上，不同国别、业务做资源隔离
- 运行时，不同国别、业务分开执行
- 部署时，采用多机房集群化部署，避免单点故障，通过数据库锁或分布式锁保证任务只被触发一次

**问题引出**

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121112807.png)

### 数据库行锁模式

在触发调度之前，更新数据库中Joblnstance的状态， 成功抢锁的才会触发调度

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121112901.png)

### 分布式锁模式

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121113049.png)

# 3.4调度器

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121113140.png)

### 3.4.1资源来源

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121113305.png)

### 3.4.2资源调度-节点选择

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121113429.png)

### 资源调度任务分片

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121113534.png)

通过任务分片来提高任务执行的效率和资源的利用率

## 3.4.3高级特性-任务编排

使用有向无环图DAG(Directed Acyclic Graph)进行可视化任务编排

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121113649.png)

故障转移：确保部分执行单元任务失败时，任务最终成功。

分片任务基于一致性hash策略分发任务， 当某Executor异常时， 调度器会将任务分发到其他Executor

## 3.4.4调度器高可用

调度器可以集群部署，做到完全的无状态，靠消息队列的重试机制保障任务一定会被调度。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121113930.png)

# 3.5执行器

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121114329.png)

核心架构

- 控制台Admin、触发器Trigger、调度器Scheduler、 执行器Executor

业务模型

- 任务元数据Job、任务实例JobInstance、任务结果JobResult、 任务历史JobHistory

触发器

- 定时描+延时消息

时间轮

- 链表、最小堆、时间轮、多级时间轮

调度器

- 资源来源

- 资源调度：节点选择、任务分片、任务编排、故障转移

执行器：注册、调度、回调、心跳检测

# 业务应用

所有需要定时、延时、周期性执行任务的业务场景，都可以考虑使用分布式定时任务。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121114925.png)

发货后超过10天未收货时系统自动确认收货

- 使用分布式定时任务的延时任务
- 使用消息队列的延时消息或者定时消息

春节集卡活动统计完成集卡的用户个数和总翻倍数

- 使用分布式定时任务的MapReduce任务
- 使用大数据离线处理引擎Hive离线做统计
- 使用大数据实时处理引擎Flink实时做累计

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121115251.png)







































青训营官方账号：https://juejin.cn/post/7100051825939709983

过年这会儿是弯道超车的好机会哈！强烈建议大家不要放纵自己！！！该学习还是要学习。
23:00-7:00 睡觉
7: 00-7: 30 赖床
7: 30-8: 00 洗漱吃饭
8: 00-8: 01 复习 Rust 基础
8: 01-8: 02 学习并发相关知识
8: 02-8: 03 学习 mysql 相关知识
8: 03-8: 05 学习 Go 语言
8: 05-8: 07 学习分布式相关理论
8: 07-8: 09 看技术书籍
8: 09-8: 10 学习 Linux
8:10-12:10 刷一会会短视频
12:10-12:30 午饭时间
12:30-13:00 午休时间
13:00-13:05 学习消息中间件
13:06-13: 10 看理财书籍
13:10-13: 15 阅读研报
13:15-13:16 看技术书籍
13:16-13:18 学习微服务相关理论
13:18-13:20 学习网关
13:20-13:21 学习分布式配置中心
13:21-18:20 打一会会游戏，放松下
18:20-18:50 晚饭时间
18:50-18:51 复习操作系统相关知识
18:51-18:52 复习计算网络相关知识
18:52-18:53 复习数据结构和算法相关知识
18:53-23:00 打一会会游戏，放松下
23:00-7:00 睡觉

