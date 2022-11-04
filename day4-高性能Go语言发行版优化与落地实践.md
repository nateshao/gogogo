---
title: day4-高性能Go语言发行版优化与落地实践
date: 2022-05-14 09:53:54
tags: 
- Go学习路线
- 字节跳动青训营
---

[TOC]

这是我参与「第三届青训营 -后端场」笔记创作活动的的第4篇笔记

![day4-高性能 Go 语言发行版优化与落地实践 ](https://cdn.jsdelivr.net/gh/nateshao/images/20220514100008.jpg)

## 「高性能 Go 语言发行版优化与落地实践 」 第三届字节跳动青训营 - 后端专场

同时这也是课表的第4天课程

<img src="https://cdn.jsdelivr.net/gh/nateshao/images/20220514095557.png" style="zoom:50%;" />

PC端阅读效果更佳，点击文末：**阅读原文**即可。

## 这节课开源收获什么？

《高性能Go语言发行版优化与落地实践》

- 优化
  - 内存管理优化
  - 编译器优化
- 背景
  - 自动内存管理和Go内存管理机制
  - 编译器优化的基本问题和思路
- 实践：字节跳动遇到的性能问题以及优化方案

### 追求极致性能

**性能优化是什么?**

- 提升软件系统处理能力，减少不必要的消耗，充分发掘计算机算力

**为什么要做性能优化?**

- 用户体验：带来用户体验的提升。让刷抖音更丝滑，让双十一购物不再卡顿
- 资源高效利用：降低成本，提高效率一很小的优化乘以海量机器会是显著的性能提升和成本节约

### 性能优化的层面

**业务层优化**

- 针对特定场景，具体问题，具体分析
- 容易获得较大性能收益

**语言运行时优化**

- 解决更通用的性能问题
- 考虑更多场景
- `Tradeoffs`

**数据驱动**

- 自动化性能分析工具----pprof
- 依靠数据而非猜测
- 首先优化最大瓶颈

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220514184635.png)

---

**性能优化与软件质量**

- 软件质量至关重要
- **在保证接口稳定的前提下改进具体实现**
- 测试用例：覆盖尽可能多的场景，方便回归
- 文档：做了什么，没做什么，能达到怎样的效果
- 隔离：通过选项控制是否开启优化
- 可观测：必要的日志输出

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220514184425.png)

---

## 目录

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220514185111.png)

## 01.自动内存管理

> 概念 | Tracing garbage collection | Generational GC | Reference counting

`Tracing garbage collection`：跟踪垃圾回收

`Generational GC `：分代GC

`Reference counting`：引用技术

## 1.1自动内存管理

- 动态内存
  - 程序在运行时根据需求动态分配的内存：`nalIoc()`

- 自动内存管理( 垃圾回收)：由程序语言的运行时系统管理动态内存
  - 避免手动内存管理，专注于实现业务逻辑
  - 保证内存使用的**正确性和安全性**: double-free problem, use-after-free problem

- 三个任务
  - 为新对象分配空间
  - 找到存活对象
  - 回收死亡对象的内存空间

## 1.1 自动内存管理-相关概念

- `Mutator`：业务线程，分配新对象，修改对象指向关系
- `Collector`：GC线程，找到存活对象，回收死亡对象的内存空间
- `Serial GC`：只有一个`collector`
- `Parallel GC`：支持多个`collectors`同时回收的GC算法
- `Concurrent GC`：` mutator(s)`和`collector(s) `可以同时执行

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220514191538.png)

`Collectors `必须感知对象指向关系的改变!

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220514191712.png)



---

评价GC算法（四个方面）

- 安全性(Safety)：不能回收存活的对象 **基本要求**

- 吞吐率(Throughput): 1 - （GC时间/程序执行总时间） **花在GC上的时间**

- 暂停时间(Pause time)：stop the world (STW) **业务是否感知**

- 内存开销(Space overhead) GC元数据开销

**追踪垃圾回收(Tracing garbage collection)**

**引用计数(Reference counting)**

## 1.2追踪垃圾回收

<img src="https://cdn.jsdelivr.net/gh/nateshao/images/20220514195724.png" style="zoom:50%;" />

- 对象被回收的条件：指针指向关系不可达的对象

主要分3步

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220514200839.png)

1. 标记根对象
   - 静态变量、全局变量、常量、线程栈等

2. 标记：找到可达对象
   - 求指针指向关系的传递闭包：从根对象出发，找到所有可达对象

3. 清理：所有不可达对象
   - 将存活对象复制到另外的内存空间(Copying GC)
   - 将死亡对象的内存标记为可分配(Mark-sweep GC)
   - 移动并整理存活对象(Mark compact GC)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220514201749.png)

如何选择：**根据对象的生命周期，使用不同的标记和清理策略**



## 1.3 分代GC (Generational GC)

- 分代假说(Generational hypothesis)：most objects die young
- `Intuition`：很多对象在分配出来后很快就不再使用了
- 每个对象都有年龄：经历过GC的次数
- 目的：对年轻和老年的对象，制定不同的GC策略，**降低整体内存管理的开销**
- 不同年龄的对象处于heap的不同区域

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220514202905.png)

年轻代(Young generation)  ：**适合Copying GC**

- 常规的对象分配
- 由于存活对象很少，可以采用`copying collection`
- GC吞吐率很高

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220514203337.png)

老年代(Old generation)：**适合Mark-sweep GC**

- 对象趋向于一直活着，反复复制开销较大
- 可以采用`mark-sweep collection`

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220514203354.png)

## 1.4引用计数

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220514231114.png)

每个对象都有一个与之关联的引用数目

对象存活的条件：当且仅当引用数大于0

- **优点**

  - 内存管理的操作被平摊到程序执行过程中

    ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220514231742.png)

  - 内存管理不需要了解`runtime`的实现细节: C++智能指针(smart pointer)



- **缺点**

  - 维护引用计数的开销较大:通过**原子操作**保证对引用计数操作的**原子性**和**可见性**

    ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220514231916.png)

  - 无法回收环形数据结构----`weak reference`

    ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220514232022.png)

  - 内存开销：每个对象都引入的额外内存空间存储引用数目

  - 回收内存时依然可能引发暂停

    ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220514232207.png)



### 01总结

- 自动内存管理的背景和意义
- 概念和评价方法
- 追踪垃圾回收
- 引用计数
- 分代GC
- 学术界和工业界在一直在致力于解决自动内存管理技术的不足之处
  PL .DI22 Low-L .atency, High-Throughput Garbage Collection



## 02.Go内存管理及优化

> Go内存分配 | Go内存管理优化



## 2.1 GO内存分配-分块

- 目标：为对象在`heap`上分配内存

- 提前将内存分块

  - 调用系统调用`mmap()`向OS申请一大块内存，例如4 MB

  - 先将内存划分成大块，例如8 KB，称作`mspan`

  - 再将大块继续划分成特定大小的小块，用于对象分配

  - `noscan mspan`：分配不包含指针的对象一GC 不需要扫描

  - `scan mspan`：分配包含指针的对象一GC 需要扫描

    ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515084502.png)

- 对象分配：根据对象的大小，选择最合适的块返回

## 2.1 GO内存分配一缓存

- `TCMalloc`：`thread caching`
- 每个p包含一个mcache用于快速分配，用于为绑定于p上的g分配对象
- `mcache`管理一组`mspan`
- 当`mcache`中的`mspan`分配完毕，向`mcentral`申请带有末分配块的`mspan`
- 当`mspan`中没有分配的对象，`mspan`会被缓存在`mcentral`中，而不是立刻释放并归还给OS

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515085802.png)

## 2.2 Go内存管理优化
![image-20220515085948402](https://cdn.jsdelivr.net/gh/nateshao/images/20220515085948.png)

- 对象分配是非常高频的操作：每秒分配GB级别的内存
- 小对象占比较高
- Go内存分配比较耗时
  - 分配路径长: g -> m -> p -> mcache -> mspan -> memory block -> return pointer
  - pprof：对象分配的函数是最频繁调用的函数之一

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515090108.png)

## 2.3我们的优化方索: Balanced GC 

- 每个`g`都绑定一大块内存(1 KB)，称作goroutine allocation buffer (GAB)

- `GAB`用于`noscan`类型的小对象分配：< 128 B

- 使用三个指针维护GAB： base, end, top

- `Bump pointer` (指针碰撞)风格对象分配
  - 无须和其他分配请求互斥
  - 分配动作简单高效

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515090507.png)

## 2.3 Balanced GC

- `GAB`对于Go内存管理来说是**一个大对象**
- 本质：**将多个小对象的分配合并成一次大对象的分配**
- 问题：`GAB`的对象分配方式会导致内存被延迟释放

- 方案：移动`GAB`中存活的对象
  - 当`GAB`总大小超过一定阈值时，将`GAB`中存活的对象复制到另外分配的`GAB`中
  - 原先的`GAB`可以释放，避免内存泄漏
  - **本质：用copying GC的算法管理小对象 **       **{根据对象的生命周期，使用不同的标记和清理策略}**

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515093523.png)

## 2.3 Balanced GC 性能收益

高峰期CPU usage降低4.6%，核心接口时延下降4.5%~7.7%

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515092719.png)

### 02总结

- Go内存管理一分块

- Go内存管理一缓存

- Go对象分配的性能问题

  - 分配路径过长

  - 小对象居多

- Balanced GC

  - 指针碰撞风格的对象分配
  - 实现了copying GC
  - 性能收益

## 03.编译器和静态分析

基本介绍 | 数据流和控制流 | 过程内和过程间分析

## 3.1编译器的结构

- 重要的系统软件
  - 识别符合语法和非法的程序
  - 生成正确且高效的代码

- 分析部分(前端front end)
  - 词法分析，生成词素(lexeme)
  - 语法分析，生成语法树
  - 语议分析，收集类型信息，进行语义检查
  - 中间代码生成，生成intermediate representation (IR)

- 综合部分(后端back end)
  - 代码优化，机器无关优化，生成优化后的 IR
  - 代码生成，生成目标代码

**主要学习编译器后端优化**

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515093948.png)

## 3.2 静态分析

- 静态分析：**不执行程序代码**，推导程序的行为，分析程序的性质。
- 控制流(Control flow)：程序执行的流程
- 数据流(Data flow)：数据在控制流上的传递

通过**分析控制流和数据流，我们可以知道更多关于程序的性质(properties)**。根据这些性质优化代码

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515094753.png)

## 3.3过程内分析和过程间分析

过程内分析(Intra-procedural analysis)

- 仅在过程内部进行分析

过程间分析(Inter-procedural analysis)

- 考虑过程调用时参数传递和返回值的数据流和控制流

为什么过程间分析是个问题?

- 需要通过**数据流分析**得知i的具体类型，才能知道i.fod)调用的是哪个foo( )
- 根据i的具体类型，**产生了新的控制流**，i.foo(), 分析继续
- 过程间分析需要**同时**分析控制流和数据流-----**联合求解，比较复杂**

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515095325.png)

### 03.总结

- 编译器的结构与编译的流程
- 编译器后端优化
- 数据流分析和控制流分析
- 过程内分析和过程间分析

## 04. Go编译器优化
函数内联 | 逃逸分析

- 为什么做编译器优化
  - 用户无感知，重新编译即可获得性能收益
  - 通用性优化
- 现状
  - 采用的优化少
  - 编译时间较短，没有进行较复杂的代码分析和优化
- 编译优化的思路
  - 场景：面向后端长期执行任务
  - Tradeoff：**用编译时间换取更高效的机器码**
- `Beast mode`
  - **函数内联**
  - **逃逸分析**
  - 默认栈大小调整
  - 边界检查消除
  - 循环展开
    ....

## 4.1函数内联(Inlining)

- 内联：将被调用函数的函数体(callee)的副本替换到调用位置(caller).上，同时重写代码以反映参数的绑定

- 优点
  - 消除函数调用开销，例如传递参数、保存寄存器等
  - **将过程间分析转化为过程内分析**，帮助其他优化，例如**逃逸分析**
- 缺点
  - 函数体变大，instruction cache (icache) 不友好
  - 编译生成的Go镜像变大
- 内联策略
  - 调用和被调函数的规模
  - …

- **函数内联能多大程度影响性能?**  ---使用micro-benchmark验证一下

  ```go
  func BenchmarkInline(b *testing.B) {
  	x := genInteger()
  	y := genInteger()
  	for i := 0; i < b.N; i++ {
  		addInLine(x, y)
  	}
  }
  func addInline(a, b int) int {
  	return a + b
  }
  ```

  ```go
  func BenchmarkInlineDisabLed(b *testing.B) {
  	x := genInteger()
  	y := genInteger()
  	for i := 0; i < b.N; i++ {
  		addNoIntine(x, y)
  	}
  }
  
  //go :noinline
  func addNoInLine(a, b int) int {
  	return a + b
  }
  ```

使用micro-benchmark快速验证和对比性能优化结果

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515101538.png)

## 4.2 Beast Mode

- Go函数内联受到的限制较多
  - 语言特性，例如`interface`, `defer`等，限制了函数内联
  - 内联策略非常保守

- `Beast mode`：调整函数内联的策略，使更多函数被内联
  - 降低函数调用的开销
  - 增加了其他优化的机会：**逃逸分析**

- 开销
  - Go镜像增加~ 10%
  - 编译时间增加

## 4.3逃逸分析

- 逃逸分析：分析代码中指针的动态作用域:指针在何处可以被访问

- 大致思路
  - 从对象分配处出发，沿着控制流，观察对象的数据流
  - 若发现指针p在当前作用域s:
    - 作为参数传递给其他函数
    - 传递给全局变量
    - 传递给其他的`goroutine`
    - 传递给已逃逸的指针指向的对象
- 则指针p指向的对象逃逸出s,反之则没有逃逸出s

- `Beast mode`：函数内联拓展了函数边界，更多对象**不逃逸**

- 优化：未逃逸的对象可以在**栈上分配**
  - 对象在栈上分配和回收很快：移动`sp`
  - 减少在`heap`上的分配，降低`GC`负担

## 4.2 Beast Mode-性能收益

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220515102718.png)

### 04.总结
- Go编译器优化的问题
- Beast mode
- 函数内联
- 逃逸分析
- 通过micro-benchmark快速验证性能优化
- 性能收益

## 课程总结

- 本节课程：高性能Go语言发行版优化与落地实践

- 性能优化
  - 自动内存管理
  - Go内存管理
  - 编译器与静态分析
  - 编译器优化

- 实践
  - Balanced GC优化对象分配
  - Beast mode提升代码性能

- **分析问题的方法与解决问题的思路，不仅适用于Go语言，其他语言的优化也同样适用**

## 参考文献

1. The Garbage Collection Handbook -- the art of automatic memory management
2. https://plumbr.io/handbook/what-is-garbage-collection
3. JEP 333: ZGC: A Scalable Low-Latency Garbage Collector https://openjdk.java.net/jeps/333
4. 数据密集型应用系统设计Designing Data-Intensive Applications: The Big ldeas Behind Reliable, Scalable, and Maintainable Systems
5. 编译原理The Dragon book, Compilers: Principles, Techniques, and Tool
6. 编译器设计Engineering a Compiler
7. 编译原理Principles and Techniques of Compilers https://silverbulltt.bitbucket.io/courses/compiler- 2022/index.html
8. 静态程序分析Static Program Analysis https://pascal-group.bitbucket.io/teaching.html
9. JVM Anatomy Quark #4: TLAB allocation https://shipilev.net/jvm/anatomy-quarks/4-tlab-allocation/
10. Constant folding, https://en.wikipedia.org/wiki/Constant folding
11. Choi, Jong-Deok, et al. "Escape analysis for Java." Acm Sigplan Notices 34.10 (1999): 1-19.





















































































