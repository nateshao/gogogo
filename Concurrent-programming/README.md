## 并发介绍

#### 进程与线程

1. 进程是程序在操作系统中一次执行过程，系统进行资源分配和调度的一个独立的单位。
2. 线程是进程的一个执行实体，是CPU调度和分配的基本单位，它是比进程更小的能独立运行的基本单位。
3. 一个进程可以创建和撤销多个线程；同一个进程中的多个线程之间可以并发执行。

#### 并发和并行

1. 多线程程序在一个核的CPU上运行，就是并发。
2. 多线程程序在多个核的CPU上进行，就是并行。

#### 并发

![并发](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/91b959478442a949362e188240f835cb.png)

#### 并行

![并行](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/fb6f8af99ec207b1b67c1d64a531cce6.png)