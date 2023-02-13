---
title: day19-TOS对象存储实战
date: 2023-01-15 13:02:47
tags:
- Go学习路线
- 字节跳动青训营
---

<img src="https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123111046.jpg" style="zoom:200%;" />

这是我参与「第三届青训营 - 后端场」笔记创作活动的的第19篇笔记。*PC端阅读效果更佳，点击文末：**阅读原文**即可。*

## 「TOS对象存储实战」 第三届字节跳动青训营 - 后端专场

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123111113.png)

[TOC]



## 课程目录

01.抖音背后的存储

- 发/刷抖音背后有何流程?背后有何种存储需求?

02.为什么对象存储

- 为什么需要对象存储呢?

03.对象存储怎么用

04.TOS字节内部实践

- TOS在字节面临的场景有哪些?工程.上的解法是?



# 01.抖音背后的存储

发/刷抖音背后有何流程？背后有何种存储需求？

## 背景

同学小明，进入了互联网顶尖大厂子虚乌有公司工作。

某一天，领导说：短视频这么火，我们也做一个**短视频APP**，这个重担就交给小明你了。

小明：机会来咯，要是做成抖音那样火爆，那升职加薪，走向人生巅峰，岂不是分分钟！

##  短视频架构初探

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123153149.png)

通过仔细思考，设计了一个如图的架构；

客户端我们肯定需要有一个，然后也要有账号，也要有评论服务；

但核心的来说，向抖音这类APP都是UGC，就是用户生产内容的平台，所以重中之重**我们的内容生产到推荐**这一条链路； 这一条链路上，会有片源，也就是承接用户上传的内容，然后还有个审核服务，对这些内容进行合规的审核，最后这些内容会根据机器学习打上标签，通过各位同学历史观看兴趣通过推荐服务推荐到各位同学的手机上

## 存储需求

把短视频生产/消费链路做更细粒度分解，小明发现到处都有视频/图片的公共存储需求。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123153412.png)

这里小明对短视频生产/消费链路做了更细粒度分解，分成了如图的几个服务。

1. 第1步是**片源服务**，将用户上传的视频存起来，这个用户原始的视频我们叫源视频；

2. 然后会有**转码服务**，将源视频转为不同码率的视频，为什么需要转码呢，因为不同的客户端（如不同的手机/电脑）型号，能接受的分辨率是不一样的，因此需要转码来适配，这里转码完后的也需要存储起来，以供后续客户端拉取；

3. 然后图下面还有个**抽帧服务**，将视频抽成不一样的帧，用于审核服务做审核，这里也需要有个存储把这些图片存储起来；

4. 最后就是**推荐服务**了，将对应视频推荐出去，然后客户端就可以来拉取观看啦；

## 存储需求量细化

小明根据未来可能的用户量，简单做数学计算，发现存储量真的大

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123155009.png)



单视频假设10MB： 

- 转码3个码率：10MB 

- 抽帧20张图片：1MB 

- 总计：21MB，24个视频/图片

 1000/S视频上传： 

- Day：1730TB,  20亿个视频/图片 
- Month：51900TB, 622亿个视频/图片 
- Year：631450TB, 7568亿个视频/图片 

一块磁盘4TB：

- Day： 432块盘 
- Month：12960块盘 
- Year：157680块盘

## 寻找天选存储

小明现在的当务之急，是寻找一个合适的存储，要求：

- 海量：从前面分析，这个存储系统一定要能够存储如此大的海量 

- 易用：好的存储能够解放业务，让业务专注于业务逻辑开发 

- 便宜：这么大的存储量，越便宜就越能省下宝贵的经费

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123155831.png)

# 02.为什么对象存储
[带你认识存储&数据库](https://mp.weixin.qq.com/s?__biz=MzIyNjE0MDI1NQ==&mid=2247490421&idx=1&sn=12b53f286191059def1d8330ce731766&chksm=e8745245df03db53034863bcac738b954757c9410ff70075b41210fc58c839a5cec8c8d3b43c&token=1386937348&lang=zh_CN#rd)这篇有讲到存储系统的分类

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123180754.png)

存储对比

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123181038.png)

这里详细类比了各类存储。单机存储我们先排除，数据量实在是太大了，单机肯定存不下，单机数据库同理。

那分布式数据库呢，现代的分布式数据库在容量和弹性上面都有很大进展，是否可以呢？答案也是No，因为分布式数据库只适合存储结构化or半结构化数据。

**什么是结构化数据呢**，就是数据和数据之间有一定关系，如用户信息，一个用户会关联包括电话号码，性别等各种维度的信息，这些信息一般都是不超过KB级别的，超过MB级别就不适合使用数据库处理了； 

唯一的选择就是分布式存储，**分布式存储是针对海量存储场景特别设计的存储，能够存储海量的大数据**；

## 分布式存储选型

但分布式存储也有分布式文件系统和对象存储，应该选择哪个呢?

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123182605.png)

## 易用性：接口对比

刚刚对比还是比较抽象的，我们来更详细的对比HDFS和TOS的接口； 

接口对于任何一种存储来说，都是最重要的表征，为什么呢，因为一种新的存储之所以提出来，是因为它在新场景下解决了其他存储很难解决的问题，针对这类场景提供了新的接口可大幅简化业务逻辑。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123183128.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123183226.png)

Bucket/Object语义

- Bucket：存储对象的桶，可类比一个云上的Map 
- Object：对象，包含如下三个部分
  - Key：对象的名字，可类比Map的Key
  - Data：对象的内容，例如视频/图片内容
  - MetaData：对象的一些元信息，如对象大小，对象Content-Type,也可以存储自定义元信息

HTTP接口

- 任何时间、任何地点、任何互联网设备上传和下载数据
- 支持HTTP协议的各种客户端都可访问

接口速览

- GET：下载对象
- HEAD：查看对象元信息
- PUT：上传对象
- DELETE：删除对象
  …

## 使用场景

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123183732.png)

# 03.对象存储怎么用

确定对象存储就是那个天选存储，选型完毕开启开发，首先要了解对象存储的用法。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123183841.png)

## Restful接口

申请完Bucket后，要开启开发，了解到对象存储对外提供的一般都是Restful风格的接口。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123184216.png)

了解对象存储PUT/GET/HEAD/DELETE接口的基本使用

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123184434.png)

## MultiUpload接口

随着开发的深入，小明发现一个问题，自己上传数GB的大视频时，由于网络不好，总是上传到99%就网络卡住了，他很恼火，翻遍对象存储各类手册，终于发现了一个解决此场景的闪电三连鞭：**MultiUpload**

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123185134.png)

## 分页列举：Listprefix接口

把上传/下载/删除对象存储等基本场景搞定后，但是想看看桶里面有哪些对象，这时候他又犯难了，继续翻遍开发手册后，他发现了一个分页列举接口: **ListPrefix**

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123185400.png)

ListPrefix接口是用于查看桶下面有哪些对象，它是一个分页接口，循环调用该接口可以遍历桶下面所有对象

# 04.TOS字节内部实践
## 开发一个对象存储

短视频应用上线后，大获成功，对象存储在公有云的使用量特别大，数据量大了，领导决定自研对象存储，这个重任又交给了小明。小明很快想出了一个经典的三层架构：

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123185532.png)

- 接入层：接入解析并处理接口请求
- 元信息层：存储对象元信息
- 存储引擎层：存储对象内容

但是架构如何细化呢，小明想了又想，先梳理当前经典业务场景，总结挑战对症下药

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123185755.png)

## 可扩展性解法之Partition

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123214237.png)

**分布式存储 = 分布式 + 单机存储**

分布式：

- 存储均匀分布
- 计算均匀分布
- 压力均匀分布

分布式系统相当于一个蜂群，每个节点都负责一小部分数据存储和计算，达到1+1 >= 2的效果

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123215539.png)

## 持久度解法之Replication

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123215615.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123215635.png)



复制(Replication)

- 数据复制多份，即多个副本
- 副本放置策略:
  - 多机架：可抵抗机架级别故障
    多机房：可抵抗机房级别故障
  - 多Region：可抵抗Region级别故障

带来

- 高持久度：不丢数据
- 强吞吐能力：多个副本可以提供服务

思考

- Replication的拷贝方式有哪些?
- Replication如何解决一致性问题?



## 成本解法之EC

Replication虽然可以解决持久度问题，但是单纯多副本拷贝成本也非常高！

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123220829.png)

EC (Erasure Coding)

- 冗余编码：可达到和多副本一样的持久度

特点

- 低冗余度：成本较单纯多副本低
- 额外计算：增加了额外的编码计算步骤

思考

- 当前的EC编码算法有哪些?
- 多机房的EC如何实现?

## 成本解法之温冷转换

数据都是有温度的，将冷数据转移到性能更差但更廉价的存储介质不就可以省下来成本么?

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123220953.png)

## 架构细化

通过上面的分析，小明把自研对象存储的架构细化了下来

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123222748.png)

API：接入层

- Bucket Meta： Bucket元信息服务
- Object Meta：对象元信息服务
- Distributed KV： Range Partition的分布式KV，用于持久化对象元数据
- Storage Engine：对象内容存取服务
- Distributed Storage Pool：分布式存储池，三副本or EC存储
- GC：垃圾回收后台服务
- Lifecycle：温冷转换后台服务

## 存储需求量细化

小明自研的对象存储上线后，公司其他业务也想使用，但是他们的要求不一样:超高的可用性!

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123222946.png)

单视频假设10MB： 

- 转码3个码率：10MB 
- 抽帧20张图片：1MB 
- 总计：21MB，24个视频/图片 

1000/S视频上传： 

- Day：1730TB,  20亿个视频/图片 
- Month：51900TB, 622亿个视频/图片 
- Year：631450TB, 7568亿个视频/图片 

一块磁盘4TB： 

- Day： 432块盘 
- Month：12960块盘 
- Year：157680块盘

## 高可用解法之拆分降低爆炸半径

首先想到的就是，一个集群拆分成多个集群，有效降低爆炸半径

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123223455.png)

## 高可用之粤核酸的启发

但对于很多业务，例如飞书等，是希望无论情况，都保证服务可用，那应该怎么办呢? ! ! !

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123223534.png)

## 高可用之镜像灾备

完全镜像的主备Bucket,出现问题随时切换，真正100%的可用性！

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123223650.png)

![未来展望](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230123223728.png)

## 课后大作业

#### 实现一个对象存储客户端

**作业要求：**

1. 在任意一个公有云中申请一个对象存储 Bucket
2. 使用你熟悉的语言，实现一个对象存储命令行客户端，要求该客户端能够
   1. 创建对象：超过 1GB 的对象使用 MultiUpload 上传，小于 1GB 的使用 Put 上传
   2. 下载对象
   3. 删除对象
   4. 查看对象是否存在
   5. 列举对象及 CommonPrefix





参考文献：

1. 青训营官方账号：https://juejin.cn/post/7101135488974585870

2. https://bytedance.feishu.cn/file/boxcn3SOO7DkrVWzb93qx08jtch



