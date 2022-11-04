---
title: day6-从需求到上线全流程+实操课
date: 2022-05-15 12:40:59
tags: 
- Go学习路线
- 字节跳动青训营
---

[TOC]

![day6-从需求到上线全流程+实操课](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220706095915.jpg)

这是我参与「第三届青训营 -后端场」笔记创作活动的的第6篇上午笔记

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220706102608.jpg)

# 「实战项目-Go语言笔记服务」上  第三届字节跳动青训营

同时这也是课表的第6天上午课程

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220703104814.png)

PC端阅读效果更佳，点击文末：**阅读原文**即可。

**一、课程背景与目标**

这篇文章主要是将前5节Go原理与实践课程的基础上，通过项目实战帮助大家把前面学过的知识应用起来

课程目标

- 将前面所学的知识应用到项目中
- 熟悉项目的代码,可以将项目正常运行
- 熟悉Kitex/Gorm的使用

**二、课前了解**

安装Docker/Postman/Git

**Kitex初体验**

通过阅读https://www.cloudwego.io/zh/docs/kitex/getting-started/尝试运行Kitex 的示例代码

1. kitex暂时没有针对Windows做支持，如果本地开发环境是Windows建议使用**WSL2**
2. 了解 etcd 是什么以及 opentracing(**链路追踪**) 是什么
3. 可以使用Minikube 或者使用Docker Desktop启动Docker
4. 安装Postman（浏览器安装**PostWoman Http接口调试插件**也可）
5. 安装Git安装教程

## 项目介绍

>项目简介：EasyNote 提供了一套比较完整的笔记后端API服务.

项目地址：https://github.com/cloudwego/kitex-examples/tree/main/bizdemo/easy_note

- 推荐版本 Golang >= 1.15

**项目模块介绍**

| 服务名称 |   模块介绍   |       技术框架        | 传输协议 | 注册中心 |  链路追踪   |
| :------: | :----------: | :-------------------: | :------: | :------: | :---------: |
| demoapi  |   API服务    | `Gorm` `Kitex ` `Gin` |   http   |   etcd   | opentracing |
| demouser | 用户数据管理 |    `Gorm` `Kitex `    | protobuf |   etcd   | opentracing |
| demonote | 笔记数据管理 |    `Gorm` `Kitex `    |  thrift  |   etcd   | opentracing |

**项目服务调用关系**

| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220703120541.jpg) | ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220703120625.png) |
| ------------------------------------------------------------ | ------------------------------------------------------------ |

**项目模块功能介绍**

| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220703192858.png) | ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220703192905.png) |
| ------------------------------------------------------------ | ------------------------------------------------------------ |

**项目技术栈**

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220703123008.png)

### 项目相关的使用框架资料

|                       |                                                              | 框架文档地址                                        | github地址                                                   | 拓展文档 |
| --------------------- | ------------------------------------------------------------ | --------------------------------------------------- | ------------------------------------------------------------ | -------- |
| RPC框架Kitex          | 框架文档                                                     | https://www.cloudwego.io/zh/docs/kitex/overview/    | https://github.com/cloudwego/kitex                           |          |
| Kitex-etcd扩展        | https://github.com/kitex-contrib/registry-etcd               | https://github.com/kitex-contrib/registry-etcd      | https://www.cloudwego.io/zh/docs/kitex/tutorials/framework-exten/registry/        https://www.cloudwego.io/zh/docs/kitex/tutorials/framework-exten/service_discovery/ |          |
|                       |                                                              |                                                     |                                                              |          |
| Kitex-OpenTracing扩展 | https://www.cloudwego.io/zh/docs/kitex/tutorials/service-governance/tracing/ | https://github.com/kitex-contrib/tracer-opentracing | https://www.cloudwego.io/zh/docs/kitex/tutorials/framework-exten/middleware/ |          |
| ORM框架Gorm           | 框架                                                         | https://gorm.cn/zh_CN/                              | https://github.com/go-gorm/gorm                              |          |
| Gorm-Opentracing扩展  | https://github.com/go-gorm/opentracing                       | https://github.com/go-gorm/opentracing              | https://gorm.cn/zh_CN/docs/write_plugins.html                |          |
| HTTP框架Gin           | 框架                                                         | https://github.com/gin-gonic/gin#gin-web-framework  | https://github.com/gin-gonic/gin                             |          |
| Gin-JWT扩展           | https://github.com/appleboy/gin-jwt#usage                    | https://github.com/appleboy/gin-jwt                 |                                                              |          |





**项目代码目录结构介绍**

> IDL：接口定义语言

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220703193933.png)

### 项目运行

#### 运行基础依赖

```dockerfile
docker-compose up
```

执行上述命令启动 MySQL、Etcd、Jaeger 的 docker 镜像

#### 运行 demonote 服务

```sh
cd cmd/note 
sh build.sh 
sh output/bootstrap.sh
```

#### 运行 demouser 服务

```sh
cd cmd/user 
sh build.sh 
sh output/bootstrap.sh
```

#### 运行 demoapi 服务

```cmd
cd cmd/api 
chmod +x run.sh 
./run.sh
```

## 参考文档

作者：青训营官方账号
链接：https://juejin.cn/post/7095977466094682148

Docker安装：https://www.runoob.com/docker/windows-docker-install.html

Git安装：https://git-scm.com

WSL2：https://docs.microsoft.com/zh-cn/windows/wsl/install

Gorm初体验：https://gorm.cn/docs/#Install



# 「从需求到上线全流程」下  第三届字节跳动青训营

## 课程结构

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704100431.png)

## 01 为什么要有流程

团队规模和流程的关系 | 瀑布模型 | 敏捷开发 | 实际的例子

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704130418.jpg)

> 为什么要有流程？

1. 个人开发者是不需要流程的
2. 超过一个人的团队就需要协作
3. 随着团队规模上升，会出现全新的问题

## 1.1 团队规模和流程的关系

复杂项目没有流程会有什么问题：

- 需求阶段：每个人都有自己的想法，团队决策需要有一个过程
- 开发阶段：多人/多端协作开发，每个人有自己的安排，相互配合需要有一个流程
- 测试阶段：产物怎样交付，测试如何开展，BUG怎么修都需要流程
- 发布阶段：怎样确保发布过程平稳丝滑，版本和流量如何控制，需要有规范
- 运维阶段：线上问题如何应急响应，处理用户反馈和线上问题需要有流程

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704151221.png)

当我们要开发的产品复杂了之后，究竟会遇到哪些问题？ 按照阶段来看： 比如我们的想法可能不同，每个人都有自己的想法 甚至有一些还是馊主意，怎么去把这些各种各样的意见统一，这样团队就需要一个决策的过程；

多人或者多端配合开发的时候，每个人都有自己的安排，你的代码跟我的代码怎么合到一起，什么时候一起联调？都需要有流程 

再到后面开发出的产物交付测试，遇到了BUG怎么修，怎么让测试的功能之间不相互影响 最终好不容易我们的产品要上线了，上线要通过哪些方式，上线之后的运维又要做哪些事情。**这些都需要有一个规范和流程** 

> 当然小公司另当别论了，一般的小公司，需求过来，交给开发人员开发-测试-部署上线，就完事了。

## 1.2 传统的瀑布模型

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704151716.png)

最传统的模型：把整个软件开发按照各个阶段排成一条线，前一个阶段完成之后进行下一个阶段，这就是传统的**瀑布模型**。

这种方式有一定的优势，就是在一些非常重视流程的公司：比如银行，支付等公司就很有用，因为一旦出了问题都会对用户造成很严重的影响。这种流程非常低效，大家的工作是定死了的，到了一定的时间你就要做一定的事情，因此会有很多时间是在等待前面的流程完成。

## 1.3 敏捷开发

后来互联网公司的流程逐渐倾向于**敏捷**，其实这里的**敏捷**指的更多的是一种思想。

追根溯源的话 2001年的时候，有一帮程序员，在美国犹他州的一个滑雪圣地举行的一次聚会，在聚会上他们定了这个宣言 就是因为当时的软件开发流程太过于重视流程本身，没法快速灵活的适应市场变化

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704153504.png)

 右项有其价值，其实就是指的在传统的开发流程里大家太过于重视流程文档 而忽视了跟客户交流快速响应变化

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704153623.png)

敏捷开发简单来说就是以更小的团队，更快速的进行迭代。因为团队小，所以大家可以围绕着一个很具体的目标开展工作，大家的合作也更加紧密；

在敏捷开发的概念体系里，有很多具体的方法，比如：scrum，kanban等等 scrum这个词的来源，就是橄榄球中的争球 大家肩并肩，共同围绕着一个目标前进。

## 1.4 The Scaled Agile Framework (SAFe)简介

在实践当中，现在敏捷已经发展出了一套规模化的管理的框架 也就说所谓的SAFe 

这套框架是为企业中实施敏捷开发提供一套方法论。如果说敏捷开发是一个团队内部的协作方式，那么**SAFe**就是在企业中，多个敏捷团队之间怎样配合。如果大家到了目前一些比较大的厂工作，应该实际接触的就是这套模型。

比较理论的东西我们不去深入讨论：比如精益产品开发，敏捷软件开发，系统思考等等 

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704153928.png)

现代的Scrum

- 敏捷教练Scrum Master
- 产品负责人Product Owner
- 敏捷团队Scrum Team
- 敏捷发布火车Agile Release Train

现在软件开发的阵型更像是特种部队，每个人都是身怀绝技，虽然大家会分工，但是并不是说产品设计只由产品经理负责，领导负责分配任务。

大家的决策和配合都是非常有凝聚力的行动。如果一个scrum就是一个战术小队，敏捷教练就好比是小队的队长，产品负责人是负责联络指挥部和发布任务的人，其他团队成员就是特种兵 大家也不是按照一个方阵去前进，而是用更敏捷的方式去前进，也就说**敏捷发布火车**。 

## 1.5 我们团队的流程

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704154332.png)

从时间表上可以看到，我们会在迭代的第二周周二开grooming会议，周五开planning会议 至于评审和反思会议，一般要是需要的时候才会开。 回到我们开头的例子，如果要开发一个类似抖音一样的复杂产品，那么需要的不是一个程序员，而是像这个的几十上百个这样的团队，通力合作 背后也不仅仅是程序员，还有各种各样的团队成员共同努力，才能最终把我们的想法变成实际的产品。

## 02 有哪些流程

> 这门课程叫做从**需求到上线全流程**，所以究竟有哪些流程才是我们这门课的重点，这里按照需求阶段、开发阶段、测试阶段、发布阶段、运维阶段。这几个阶段来介绍一下我们在实践当中究竟要做什么？

## 2.1 需求阶段

>不要浪费时间讨论不应该存在的问题

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704160638.png)

我们作为后端开发，除了砍需求，还要学会站在用户的角度评估需求。

**MVP( minimum viable product,最小化可行产品)思想**

- 站在用户的角度思考，

- 收集用户反馈，快速迭代



那么需求究竟应该怎么评估 这里有一个MVP的思维方式：如果我们要给用户造一辆车，我们不应该第一天给他一个轮子，第二天给两个轮子，第三天给他一个底盘，第四天才让他开上车 

**我们应该先给用户一个简单能用的产品，比如一个滑板车，一个自行车，根据用户的反馈我们再逐步把车的功能升级，最终变成用户想要的产品。** 

---

另外一个评估需求的方法：**四象限法**： 当你很多任务的时候，可以按照这个坐标，把他们按照重要性和紧急程度分类

1. 有些事情既重要又紧急，那么我们就应该先去做；
2. 不重要又不紧急的可以最后做；
3. 而有些事情虽然重要但是不紧急他们的优先级应该比那些紧急但是不重要的事情高，因为如果我们不去处理，后面他们就会变得又重要又紧急 这个理论的原则是先判断事情的重要性，再判断紧急程度 一个高效的占比，应该是大多数时间在处理重要但不紧急的事情，因为一旦一件事情变成了紧急，那我们就容易犯错误，因此如果每天大部分时间都在处理重要且紧急的事情，那么其实是不健康的。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704181428.png)



## 2.2 开发阶段

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704182404.png)

## 2.2 云原生下的开发

云原生下的开发，一个最大的区别就是**部署的形式不同** 

- **传统虚拟机上的服务开发**，是在物理机或者是更底层上虚拟出多个虚拟主机，然后在每个虚拟主机中安装软件和依赖 虚拟机需要有专门的运维人员维护 本地开发的时候也大多是直接在电脑上运行程序 

- 开发云原生的后端程序，容器是从操作系统中虚拟出来的，所有容器共享宿主机的系统，通过cgroup，namespace和union mount实现了容器之间的隔离。因此在部署的时候，应用和其依赖的系统是整体打包成一个镜像的。后端开发不再依赖运维人员创建程序的运行时环境

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704195937.png)

云原生带来的另外一个改变就是**微服务**。web应用的主要架构是SOA架构，在一个服务中多个不同的模块构成了一个部署单元，各个模块作为一个整体部署和伸缩。

这种架构下往往服务会形成一个很大的代码仓库，大家共同维护一个大型的系统。好处是模块之间的调用不需要通过RPC，但是坏处就是加减机器的时候不能安模块处理；开发的工作也需要多人共同进行充分的集成测试，保证不会把别人的东西改坏。

而微服务架构是把模块拆到不同的服务中去，拆分的粒度更细，可以让每个模块独立的扩容/缩容 同时可以让少数几个人维护一个仓库，更适合敏捷的开发流程。

---

云原生让开发环境也逐渐云化。而云原生的IDE就可以很好的解决这个问题。借助容器技术，可以轻松创建一个模拟线上的开发环境。比如你不要再纠结本地多个java的版本，go和python的还有依赖的库也可以随意切换。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704201146.png)

还有最大因素就是团队的**分支策略** 

不同的分支策略往往影响团队的开发流程。我们日常工作中写代码都是基于主干的某个版本进行修改，改完之后再把代码合并回主干形成新的版本。这里就会有一些协作上的问题： **多个团队成员各自用什么分支？修改有冲突怎么解决？出了问题代码如何回退？** 为了应对这些问题： 有些团队会有一个专门的分支叫做release分支，大家都把代码合并到release分支，然后测试，发布，之后再把release分支合会master 有些团队会直接把开发的分支合入master，然后再用某个master上的commit发布 之所以有各种各样的分支策略，就是因为我们在后续的测试和发布阶段要按照对应的分支和commit进行交付 

**git的知识也是需要重点掌握的**

## 2.4 代码规范、自测和文档

**代码规范**

- 养成良好的注释习惯，超过三个月的代码，自己都会忘了当时在想什么
- 不要有魔法数字， 魔法字符串
- 重复的逻辑抽象成公共的方法，不要copy代码
- 正确使用IDE的重构功能，防止修改错误

**自测**

- 单元测试
- 功能环境测试
- 测试数据构造

**文档**

- 大型改造需要有技术设计文档，方案评审
- 好的接口文档能更方便的和前端进行沟通

在开发阶段，代码规范，自测和文档也是非常重要的 一个有经验的开发和新人写的代码，**往往最大的差别不是功能的实现，而是在于代码的风格和规范**。

比如有一些原则可以很容易遵循： **良好的注释习惯**，有复杂的地方时间长了自己都会忘。

**不要有魔法数字和魔法字符串**，比如在判断条件里判断某个变量等于2，2代表什么？可以用常量来定义 重复的逻辑可以抽象成公共的方法，不要到处copy代码，不然每次修改都要改很多地方。

正确使用IDE的功能进行重构，不要手动去编辑或者全局替换。另外在开发阶段还要进行自测和文档编写，在开发过程中，随时进行一些静态代码扫描，也能显著提高代码质量。遵守开发的规范就像大家随手捡起手边的垃圾，是一个对团队有益的好习惯。

## 2.3 测试阶段

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704202309.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704202356.png)

左边这个图是传统的测试金字塔模型，他的意思是：**越底层的测试粒度越细，就需要越多的数量去覆盖所有场景，越顶层的测试越能用少量的case覆盖大多数场景**。但是有一个软件开发中的常识：**越早发现的缺陷解决成本越低**。因为85%的缺陷是在开发阶段引入的，而如果要在上线之后修复他们，话费的成本可能是一开始就解决他们的数百倍 所以我们还是要尽可能早的发现bug，进行充分的单元测试 

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704224315.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704224653.png)

我们日常开发的应用，如图，由客户端发送请求到网关，网关请求到后端服务器。比如我们抖音的测试同学，可能每人手里会有好几个手机，分别用来测试不同的功能 对于后端的服务，也可能有不同的版本，因此在测试中会有一个虚拟的环境概念 我们用特定的设备可以通过某些设置，让他请求到对应的后端服务器，从而达到测试对应的后端服务的目的 这样一个从客户度到服务端的一整套体系，称为一个环境。

在实际的工作中，我们一般至少需要三套环境： 

1. **功能环境**是用于开发和测试新开发的功能的
2. **集成环境**是为了把不同的功能合并在一起测试
3. **回归环境**是为了验证新的功能对老功能没有影响。具体要根据你开发的应用采用的架构，这里只是一个最简化的模型

## 2.4 发布阶段

**发布过程中要做的事情**

**发布负责人**

- 负责按照计划执行发布
- 需要通知各个相关人员发布进展
- 观察各个服务的发布状态，及时处理异常

**变更服务的相关RD**

- 按照上线checklist检查服务的日志，监控，响应上线过程中的告警
- 对于自己负责的改动，在小流量或者是预览环境进行功能验证
- 执行发布计划中的其他操作(如线上配置，数据处理等)

**值班同学**

- 发布过程中的监控和告警需要特别关注，如果有异常需要立刻判断
- 是否由变更引起
- 如果有变更引起的告警或者用户反馈，需要及时中止发布

## 2.4 发布模式-蛮力发布

>简单粗暴，直接用新版本覆盖老版本。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704232136.png)

## 2.4 发布模式-金丝崔发布

由于金丝崔对瓦斯极其敏感，因此以前矿工开矿下矿洞前，先会放一只金丝崔进去探是否有有毒气体，看金丝崔能否活下来，金丝崔发布由此得名。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704232846.png)



## 2.4 发布模式-浪动发布

>每个实例都通过金丝雀的方式逐步放大流量，对用户影响小，体验平滑。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704232722.png)





## 2.4 发布模式-蓝绿发布

> 把服务分成蓝绿两组，先把蓝组流量摘掉然后升级，只用绿组提供服务，之后切换全部流量，只用蓝组提供服务，然后升级绿组服务，最终两组全部升级。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704233022.png)





## 2.4 发布模式-红黑发布

> 和蓝绿发布类似，但是发布时会动态扩容出一组新的服务，而不需要常备两组服务。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704233309.png)

其实发布的模式还不止这几种。实际工作中，我们的发布使用的是滚动发布，发布的负责人需要**关注滚动的粒度和时间，以及具体执行的进度**。因为这种方式对用户的体验最平滑，同时公司也有强大的流量控制能力，能够平滑的切换流量能够支持滚动。

但是仍然有一些场景需要使用蓝绿发布 可能有些公司因为发布需要在用户低峰期进行，所以在**那些公司发布的时候，往往都是在夜深人静的时候**，这也就是程序员经常要晚上加班的很大一个原因

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704233554.png)

## 2.5 运维阶段

当故障发生之后，我们是有几个关键的动作的： 

1. 首先是止损，尽快去让服务恢复功能 
2. 其次是要让服务的上下游感知到出了问题。

当上面两个动作做了之后大家才需要去定位和修复问题 所以大家不要反过来，线上页面都打不开了，第一时间打开IDE开始看代码 这个其实是有问题的

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704233646.png)

以字节为例：

公司在发展过程中，逐渐形成了十分复杂的超大规模微服务体系。为了实现对这些复杂微服务的监控，我们往往会在微服务中添加埋点采
集Metrics、Logging、 分布式Trace等多种数据。

![字节内部的监控平台](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704233931.png)

## 03 流程怎样优化

## 3.1 怎样让生活更美好

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704234131.png)

- 在重视质量的团队，效丰往往比较低
- 在重视效率的团队，事故往往比较多

---

但是，时代又变了 我们看到技术的发展，让质量和效率同时得到了提高 把规范融入流程，把流程自动化 从需求到上线全流程自动化，我们就同时提高了质量和效率

- 技术的发展会带来质量和效率的同时提高

- 将质量保障融入到流程，将流程自动化
- 从需求到上线全流程自动化，同时提高质量和效率

## 3.2 DevOps

所谓DevOps，就是由左侧的Dev 和 右侧的 Ops组成 从需求开始，写代码，编译，测试，发布，运维，监控 形成了一个闭环 于是我们可以进行持续集成，持续交付。也就说CI和CD，整个流程是密不可分的。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704234349.png)

在DevOps的基础之上，我们还可以怎样优化？ 这里就引出了一个效率竖井的概念。可以看到，从需求到交付的过程中 我们真正产生价值的开发，测试等等动作占比是很低的。大量的时间可能是在等待和传递 比如测试一直在等开发把环境部署好 另外人和人之间的沟通，也很慢

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220704234539.png)

## 3.3 全流程自动化
通过效能平台串联各个阶段

- 需求发起研发流程的自动化
- 写代码，测试环境部署的自动化
- 自动化测试触发和报告分析
- 发布过程可观测融入流程

减少无价值的等待

- 分析整个流程的耗时，计算真正产生价值的时间
- 不断优化流程，让有价值的流程时间占比上升





## 参考文档

作者：青训营官方账号

- 链接：https://juejin.cn/post/7097126973163454494#heading-0
- https://bytedance.feishu.cn/docx/doxcnfOBkQhewRL6P35KBFbNG8L
- https://bytedance.feishu.cn/file/boxcnBFwiH4ItgbiBBADJcFTKBc



















































































































