---
title: day10-Git的正确使用姿势与最佳实践
date: 2022-07-16 15:32:33
tags: 
- Go学习路线
- 字节跳动青训营
---

[TOC]

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718150716.jpg)

## 「Git的正确使用姿势与最佳实践」 第三届字节跳动青训营 - 后端专场

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718150424.png)

同时这也是课表的第9天课程《Git的正确使用姿势与最佳实践》。PC端阅读效果更佳，点击文末：**阅读原文**即可。

## Git是什么

官网介绍：

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718151437.png)

Git 是一个免费和开源的分布式版本控制系统，旨在以速度和效率处理从小型到大型项目的所有内容。

Git易于学习， 占用空间小，性能快如闪电。它优于 SCM 工具，如 Subversion、CVS、Perforce 和 ClearCase，具有廉价的本地分支、方便的暂存区域和 多个工作流等功能。

**工作上用的比较多的就是Git了，像Git衍生出的Github，Gitee，Gitlab等等，当然也有其他的公司用Svn**

## 方向介绍

|   方向    |                           具体能力                           |
| :-------: | :----------------------------------------------------------: |
| 代码托管  | 负责管理公司内数十万的代码仓库,并在这之上对代码管理的相关功能进行迭代，提升研发活动的效率及质量 |
| 代码智能  | 提供更准确高效的代码搜索能力和代码导航能力，支持多种场景下的代码跳转，帮助用户更高效的去阅读代码 |
| 代码分析  | 提供一种代码检查能力，目的是在整个研发流程中自动的发现并反馈代码中存在的代码结构、代码漏洞、代码风格等问题 |
| 持续集成  | 一种软件开发实践，团队成员频繁将他们的工作成果集成在一起。每次提交后，自动触发运行一次包含自动化验证集的构建任务，以便能尽早发现集成问题 |
| Cloud IDE | 一个开箱即用的云端开发环境，支持node/python/go/java/c++等多种编程语言。可以在云端开发环境中编写、编译、运行和调试项目 |



## 为什么要学习Git

协同工作：业界绝大多数公司都是基于Git进行代码管理，因此**Git是一个程序员的必备技能**

开源社区：目前绝大多数的开源项目都是基于Git维护的，参与这些项目的开发都需要使用Git。

常见问题

1. 入职后按照文档进行Git配置，但是配置后依然拉取代码有问题，缺少自己排查配置问题的能力
2. 研发流程中进行一 些异常操作，不符合研发规范，不清楚保护分支，MR/ PR等概念

课程目标

1. 学习基本的Git命令，并了解原理，在遇到Git相关问题时，能自行排查并解决
2. 了解研发流程中的基本概念和规范，学会正确的使用Git

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718152448.png)

## 01 Git是什么
## 1.1 版本控制

Git是什么?

```sh
Git is a free and open source distributed version control system designed to handle everything from small to very large projects with speed and efficiency.
```

版本控制是什么?

- 一种记录一个或若干文件内容变化，以便将来查阅特定版本修订情况的系统

为什么需要版本控制?

- 更好的关注变更，了解到每个版本的改动是什么，方便对改动的代码进行检查，预防事故发生；也能够随时切换到不同的版本，回滚误删误改的问题代码; 

|  版本控制类型  | 代表性工具 |                          解决的问题                          |
| :------------: | :--------: | :----------------------------------------------------------: |
|  本地版本控制  |    RCS     |                      本地代码的版本控制                      |
| 集中式版本控制 |    SVN     | 提供一个远端服务器来维护代码版本，本地不保存代码版本，解决多人协作问题 |
| 分布式版本控制 |    Git     |  每个仓库都能记录版本历史，解决只有一个服务器保存版本的问题  |

### 1.1.1 本地版本控制-RCS

**最初的方式**：通过本地复制文件夹，来完成版本控制，- 般可以通过不同的文件名来区分版本

**解决方案**：开发了一些本地的版本控制软件，其中最流行的是RCS

**基本原理**：本地保存所有变更的补丁集，可以理解成就是所有的Diff,通过这些补J，我们可以计算出每个版本的实际的文件内容

**缺点**：RCS这种本地版本控制存在最致命的缺陷就是只能在本地使用，无法进行团队协作，因此使用的场景非常有限，因
此衍生出了集中式版本控制

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718155058.png)

###  1.1.2 集中版本控制-SVN
![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718155149.png)

1. 提供一个远端服务来保存文件， 所有用户的提交都提交到该服务器中。
2. 增量保存每次提交的Diff,如果提交的增量中和远端现存的文件存在冲突，则需要本地提前解决冲突。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718155149.png)

**优点:**

1. 学习简单，更容易操作。
2. 支持二进制文件，对大文件支持更友好。

**缺点:**

1. 本地不存储版本管理的概念，所有提交都只能联上服务器后才可以提交。
2. 分支上的支持不够好，对于大型项目团队合作比较困难。
3. 用户本地不保存所有版本的代码，如果服务端故障容易导致历史版本的丢失。

### 1.1.3 分布式版本控制-Git

**基本原理：**

1. 每个库都存有完整的提交历史，可以直接在本地进行代码提交
2. 每次提交记录的都是完整的文件快照，而不是记录增量

3. 通过Push等操作来完成和远端代码的同步

**优点:**

1. 分布式开发，每个库都是完整的提交历史，支持本地提交,强调个体
2. 分支管理功能强大，方便团队合作，多人协同开发
3. 校验和机制保证完整性，-般只添加数据，很少执行删除操作，不容易导致代码丢失

**缺点:**

1. 相对SVN更复杂，学习成本更高
2. 对于大文件的支持不是特别好(git-lfs 工具可以弥补这个功能)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718155629.png)





## 1.2 Git发展历史

作者 Linus Torvalds(就是Linux这个项目的作者，同时也是Git的作者)。

开发原因：怀疑Linux团队对BitKeeper (另一种分布式版本控制系统，专有软件)进行了逆向工程，BitKeeper 不允许Linux团队
继续无偿使用。因此决定自己开发一个分布式版本控制系统。

开发时间：大概花了两周时间，就完成了Git的代码第一个版本，后续Linux项目就开始使用Git进行维护。

Github：https://github.com

- 全球最大的代码托管平台，大部分的开源项目都放在这个平台上。

Gitlab：https://gitlab.com/gitlab-org

- 全球最大的开源代码托管平台，项目的所有代码都是开源的，便于在自己的服务器上完成Gitlab的搭建。

Gerrit：https://android-review.googlesource.com/

- 由Google开发的一个代码托管平台，Android 这个开源项目就托管在Gerrit 之上。

随着 Git 的发展，基于 Git 也衍生出了很多平台。除此之外，还有 BitBucket, Coding, 码云，阿里云效平台等等，每个平台都有自己的使用场景和优势。



## 02 Git的基本使用方式

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718160953.png)



## 2.1 Git目录介绍

项目初始化

```sh
mkdir study
cd study
git init
```

其他参数

```sh
--initial-branch 初始化的分支
--bare 创建一个裸仓库(纯Git目录，没有工作目录)
--template 可以通过模版来创建预先构建好的自定义git目录
```

在这里我们要重点关心一下这个 git 目录，因为我们后续每一个 git 操作都会映射到这个 git 目录之中，通过这里面的文件我们可以映射出所有版本的代码

|                           Git仓库                            |                        工作区&暂存区                         |
| :----------------------------------------------------------: | :----------------------------------------------------------: |
| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718165622.png) | ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718165641.png) |

我们刚刚看到 git 目录里面有个 config 文件，那这个 git 配置到底是个什么东西呢，我们又可以配置哪些内容呢，我们一起来了解一下 Git 配置这个概念

### 2.1.1 Git Config

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718165906.png)

每个级别的配置可能重复,但是低级别的配置会覆盖高级别的配置

### 2.1.2 常见Git配置

用户名配置

```bash
// 对当前用户的所有仓库有效
git config --global user.email "你的名字"
git config --global user.email "你的邮箱"
```

git config命令查看用户名，邮箱

```bash
git config user.name
git config user.email
```

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718172324.png)

## 2.2 Git Remote

查看Remote

```bash
git remote -v
```

添加Remote

```bash
git remote add origin_ssh git@github.com.git/git.git
git remote add origin_http https://github.com/git/git.git
```

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718172718.png)

### 2.2.1 HTTP Remote

URL: https://github.com/git/git.git

（引出免密配置）我们本地是如何与 remote 进行通信的呢，一般会通过 http 和 ssh 两种协议，这两种协议都需要对身份进行认证，类似 go 这种语言，依赖库很多，所以我们需要不断的输入认证的账号密码，肯定是一件很麻烦的事情，因此我们需要配置一下免密的认证方式

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718173328.png)

### 2.2.2 SSH Remote

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718173512.png)

## 2.3 Git Add

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718173602.png)

## 2.4 Git Commit

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718173702.png)

## 2.5 Objects

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718173803.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718173823.png)

## 2.6 Refs

除了 objects 文件有变化，我们发现 refs 的文件内容也有变化

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718174830.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718174855.png)



**Branch**

- git checkout -b可以创建一个新分支
- 分支一般用于开发阶段，是可以不断添加Commit进行迭代的

**Tag**

- 标签一般表示的是一 个稳定版本，指向的Commit一般不会变更
- 通过git tag命令生成tag



![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718175105.png)



## 2.7 Annotation Tag

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718175123.png)

## 2.8 追溯历史版本

获取当前版本代码

- 通过Ref指向的Commit可以获取唯一的代码版本。

获取历史版本代码

- Commit里面会存有parent commit字段，通过commit的串联获取历史版本代码。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718175423.png)

| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718175846.png) | ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718175901.png) |
| ------------------------------------------------------------ | ------------------------------------------------------------ |

## 2.9 修改历史版本

1. commit - amend

  通过这个命令可以修改最近的一次commit信息，**修改之后commit id会变**

2. rebase

  通过git rebase -i HEAD~3可以实现对最近三个commit的修改

  1. 合并commit
  2. 修改具体的commit message
  3. 删除某个commit

3. filter - branch

  该命令可以指定删除所有提交中的某个文件或者全局修改邮箱地址等操作

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718180152.png)

## 2.10 Objects

| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718180333.png) | ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718180349.png) |
| ------------------------------------------------------------ | ------------------------------------------------------------ |

## 2.11 Git GC

GC

- 通过git gc命令，可以删除一些不需要的object，以及会对object进行一些打包压缩来减少仓库的体积。

Reflog

- reflog是用于记录操作日志，防止误操作后数据丢失，通过reflog来找到丢失的数据，手动将日志设置为过期。

指定时间

- git gc prune = now指定的是修剪多久之前的对象，默认是两周前

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718180713.png)

**完整的Git视图**

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718180818.png)

## 2.13 Git Clone & Pull & Fetch

**Clone**

- 拉取完整的仓库到本地目录，可以指定分支，深度。

**Fetch**

- 将远端某些分支最新代码拉取到本地，不会执行merge操作，会修改refs/remote内的分支信息，如果需要和本地代码合并需要手动操作。

**Pull**

- 拉取远端某分支，并和本地代码进行合并，操作等同于git fetch + git merge,也可以通过git pull --rebase完成git fetch + git rebase操作。可能存在冲突，需要解决冲突。



## 2.14 Git Push
Push是将本地代码同步至远端的方式。

常用命令

一般使用git push origin master命令即可完成
冲突问题

1. 如果本地的commit记录和远端的commit历史不一致，则会产生冲突，比如git commit --amend or git rebase
   都有可能导致这个问题。

2. 如果该分支就自己一个人使用，或者团队内确认过可以修改历史则可以通过git push origin master -f
来完成强制推送，一般不推荐主干分支进行该操作，正常都应该解决冲突后再进行推送。

推送规则限制

- 可以通过保护分支，来配置一些保护规则， 防止误操作，或者一些不合规的操作出现，导致代码丢失。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718181548.png)





**02. 常见问题**

1. 为什么我明明配置了Git配置，但是依然没有办法拉取代码?
   - 免密认证没有配。
   - Instead Of配置没有配，配的SSH免密配置，但是使用的还是HTTP协议访问。

2. 为什么我Fetch了远端分支，但是我看本地当前的分支历史还是没有变化?
   - Fetch会把代码拉取到本地的远端分支，但是并不会合并到当前分支，所以当前分支历史没有变化。

## 03 Git研发流程

**常见问题**

1. 在Gerrit平台上使用Merge的方式合入代码
2. 不了解保护分支，Code Review, CI等概念，研发流程不规范
3. 代码历史混乱，代码合并方式不清晰

## 3.1 不同的工作流

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718184321.png)

## 3.2 集中式工作流

什么是集中式工作流?

- 只依托于master分支进行研发活动

工作方式

1. 获取远端master代码
2. 直接在master分支完成修改
3. 提交前拉取最新的master代码和本地代码进行合并(使用rebase),如果有冲突需要解决冲突
4. 提交本地代码到master

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718184718.png)



### 3.2.1集中式工作流-Gerrit

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718184927.png)

Gerrit是由Google开发的一款代码托管平台，主要的特点就是能够很好的进行代码评审。

在aosp (android open source project)中使用的很广，Gerrit 的开发流程就是一种集中式工作流。

基本原理

1. 依托于Change ID概念，每个提交生成一个单独的代码评审。
2. 提交上去的代码不会存储在真正的refs/heads/下的分支中，而是存在一个refs/for/ 的引用下。

3. 通过refs/meta/config下的文件存储代码的配置，包括权限，评审等配置,每个Change都必须要完成Review后才能合入。

**优点**

1. 提供强制的代码评审机制，保证代码的质量
2. 提供更丰富的权限功能，可以针对分支做细粒度的权限管控
3. 保证master的历史整洁性

4. Aosp 多仓的场景支持更好

**缺点**

1. 开发人员较多的情况下，更容易出现冲突。
2. 对于多分支的支持较差，想要区分多个版本的线上代码时，更容易出现问题。
3. 一般只有管理员才能创建仓库，比较难以在项目之间形成代码复用，比如类似的fork操作就不支持。

## 3.3 分支管理工作流

| 分支管理工作流 | 特点                                                         |
| -------------- | ------------------------------------------------------------ |
| Git Flow       | 分支类型丰富，规范严格                                       |
| Github Flow    | 只有主干分支和开发分支，规则简单                             |
| Gitlab Flow    | 在主干分支和开发分支之上构建环境分支，版本分支，满足不同发布or环境的需要 |

### 3.3.1 分支管理T作流-Git Flow

Git Flow时比较早期出现的分支管理策略。

**包含五种类型的分支**

- Master：主干分支
- Develop：开发分支
- Feature：特性分支
- Release：发布分支
- Hotfix：热修复分支

**优点**
如果能按照定义的标准严格执行，代码会很清晰，并且很难出现混乱。

**缺点**
流程过于复杂，上线的节奏 会比较慢。由于太复杂，研发容易不按照标准执行，从而导致代码出现混乱。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718190108.png)





### 3.3.2 分支管理工作流Github Flow

Github的工作流，只有一个主干分支，基于Pull Request往主干分支中提交代码。

选择团队合作的方式

1. owner创建好仓库后，其他用户通过Fork的方式来创建自己的仓库，并在fork的仓库上进行开发
2. owner创建好仓库后，统一给团队内成员分配权限， 直接在同一个仓库内进行开发

| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718191456.png) | ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718191810.png) |
| ------------------------------------------------------------ | ------------------------------------------------------------ |

创建一个Pull Request

1. 创建一个main主分支
2. 创建一个feature分支
3. 创建一个feature 到main 的Pull Request

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718192305.png)

可以在Pull Request页面执行CI/CA/CR等操作，都检查通过后，执行合入。

| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718192800.png) |
| ------------------------------------------------------------ |
| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718192807.png) |
| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718192815.png) |
| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718192819.png) |
| 可以通过进行一些保护分支设置，来限制合入的策略，以及限制直接的push操作。 |
| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718193030.png) |

### 3.3.2 分支管理工作流-Gitlab Flow

Gitlab推荐的工作流是在GitFlow和Github Flow 上做出优化，既保持了单一主分支的简便， 又可以适应不同的开发环境。

原则: upstream first上游优先

只有在上游分支采纳的代码才可以进入到下游分支，一般上游分支就是master.

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718205130.png)

## 3.4 代码合并

**Fast-Forward**

不会产生一个merge节点，合并后保持一个线性历史， 如果target分支有了更新，则需要通过rebase操作更新source branch后才可以合入。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718205306.png)

Three-Way Merge：三方合并，会产生一个新的merge节点

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220718205417.png)



## 3.5 如何选择合适的工作流
选择原则：**没有最好的，只有最合适的**

针对小型团队合作，推荐使用Github工作流即可

1. 尽量保证少量多次，最好不要一次性提交上千行代码
2. 提交Pull Request后最少需要保证有CR后再合入

3. 主干分支尽量保持整洁，使用fast-forward 合入方式，合入前进行rebase

**大型团队合作，根据自己的需要指定不同的工作流，不需要局限在某种流程中。**

**常见问题**

1. 在Gerrit平台上使用Merge的方式合入代码。
   - Gerrit是集中式工作流，不推荐使用Merge方式合入代码，应该是在主干分支开发后，直接Push。

2. 不了解保护分支，Code Review, CI 等概念，研发流程不规范。
   - 保护分支：防止用户直接向主干分支提交代码，必须通过PR来进行合入。
     Code Review, CI： 都是在合入前的检查策略，Code Review是人工进行检查，CI 则是通过一些定制化的脚本来进行一些校验。

3. 代码历史混乱，代码合并方式不清晰。
   - 不理解Fast Forward和Three Way Merge的区别，本地代码更新频繁的使用Three Way的方式，导致生成过多的Merge节点，使提交历史变得复杂不清晰。

总结：

课程内容 Git 是一个分布式版本控制工具，由 linus 开发，衍生出 github gitlab gerrit 等平台 Git 配置，Git 代码提交，Git 代码同步基本命令，以及 git 管理代码的原理；帮助我们更好的知道如何正确使用 Git 命令 讲述不同的研发流程，有以 gerrit 为代表的集中式工作流，和 gitlab/github 为代表的分支管理工作流，讲述了一些代码提交规范，保护分支，codereview 等概念，帮助我们规范研发流程 。

参考链接：

https://bytedance.feishu.cn/file/boxcnsl8BgASj1TA9lzYb0TevEg











