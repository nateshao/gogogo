---
title: day18-深入理解RDBMS
date: 2023-01-15 13:02:29
tags:
- Go学习路线
- 字节跳动青训营
---



![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122120920.jpg)

这是我参与「第三届青训营 - 后端场」笔记创作活动的的第18篇笔记。*PC端阅读效果更佳，点击文末：**阅读原文**即可。*

## 「深入理解RDBMS」 第三届字节跳动青训营 - 后端专场

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230121123806.png)



[TOC]



## 课程目录

01.经典案例

- 喜闻乐见：从一场抖音红包雨说起~

02.发展历史

- 源远流长: RDBMS从1970年第一篇论文发 布至今已有50余年，衍生出Oracle, DB2，MySQL，SQL Sever, Aurora等一系列知名数据库产品。

03.关键技术

- 万变归宗：无论RDBMS如何演变，其核心都是SQL引擎、存储引擎、事务引擎。

04.企业实践

- 繁花似锦：RDBMS广泛的应用于互联网、金融、电信、电力等领域，成为了各类企业级应用的数据基石。

# 01.经典案例

## 从一场红包雨说起

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122192909.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122192921.png)



```sql
-- 从抖音的账户上扣除一个小目标
UPDATE account table SET balance = balance - '小目标' WHERE name = '抖音';
-- 给羊老师的账户加上一个小目标
UPDATE account table SET balance = balance + '小目标' WHERE name = '羊老师';
```

## RDBMS事务ACID

事务(Transaction)：是由一组SQL 语句组成的一个程序执行单元(Unit)，它需要满足ACID特性。

```sql
BEGIN;
-- 从抖音的账户上扣除一个小目标
UPDATE account table SET balance = balance - '小目标' WHERE name = '抖音';
-- 给羊老师的账户加上一个小目标
UPDATE account table SET balance = balance + '小目标' WHERE name = '羊老师';
COMMIT;
```

ACID：

- 原子性( **A** tomicity)：事务是一个不可再分割的工作单元， 事务中的操作要么都发生，要么都不发生。
- 一致性( **C**onsistency)：数据库事务不能破坏关系数据的完整性以及业务逻辑上的一致性。
- 隔离性( **I**solation)：多个事务并发访问时，事务之间是隔离的，一个事务不应该影响其它事务运行效果。
- 持久性( **D**urability)：在事务完成以后,该事务所对数据库所作的更改便持久的保存在数据库之中，并不会被回滚。

| 原子性 | ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122194314.png) |
| ------ | ------------------------------------------------------------ |
| 一致性 | ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122194333.png) |
| 隔离性 | ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122194352.png) |
| 持久性 | ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122194418.png) |

## 红包雨与高并发

全国14亿人，假设有10亿人同时开抢红包，每秒处理一个请求， 那需要31年才能完成。春节完了，抖音可能也被大家嫌弃了... 

-------高并发Concurrency

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122200451.png)

Case 6：假设除夕晚上大家正在愉快的从抖音身上"薅羊毛”，这时候服务器挂了，程序员花了一个小时，头发都掉光了，终于修好了。这时候发现李谷一老师《难忘今宵》都唱完了。“抖音宕机"秒上热搜...------------高可靠、高可用High Reliability/Availability

## 大家一起"薅羊毛”

大家今年红包雨都抢到多少钱呢？弹幕打出来~

1. 超过100块。 欧皇就是你!
2. 超过50， 不到100块。LOL至臻皮肤你值得拥有~
3. 超过10, 不到50块。点份杨国福麻辣烫安慰一下自己。
4. 超过1, 不到10块。买袋薯片，隔壁小孩都馋哭了。
5. 不到1块。 怎么说呢，还不够电费的...

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122200949.png)

# 02.发展历史

## 前DBMS时代一人工管理

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122201052.png)

**文件系统**

1950s，现代计算机的雏形基本出现。1956年IBM发布 了第一个的磁盘驱动器 -- Model 305 RAMAC，从此数据存储进入磁盘时代。在这个阶段，数据管理直接通过文件系统来实现。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122201253.png)

## DBMS时代

1960s，传统的文件系统已经不能满足人们的需要，数据库管理系统(DBMS)应运而生。

DBMS：按照某种数据模型来组织、存储和管理数据的仓库。所以通常按照数据模型的特点将传统数据库系统分成网状数据库、层次数据库和关系数据库三类。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122201752.png)

传统的文件系统难以应对数据增长的挑战，也无法满足多用户共享数据和快速检索数据的需求。 层次型、网状型和关系型数据库划分的原则是数据之间的联系方式。

- 层次数据库是按记录来存取数据的；

- 网状数据库是采用网状原理和方法来存储数据；

- 关系型数据库是以行和列的形式存储数据。

### DBMS数据模型-网状模型

网状数据库所基于的网状数据模型建立的数据之间的联系，能反映现实世界中信息的关联，是许多空间对象的自然表达形式。

1964年，世界上第一个数据库系统集成数据存储(Integrated Data Storage, IDS) 诞生于通用电气公司。IDS是世界上第一一个网状数据库，奠定了数据库发展的基础，在当时得到了广泛的应用。在1970s网状数据库系统十分流行，在数据库系统产品中占据主导地位。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122202123.png)

网状数据模型是以记录类型为结点的网络结构，即一个结点可以有一个或多个下级结点，也可以有一个或多个上级结点，两个结点之间甚至可以有多种联系，例如“教师”与“课程”两个记录类型，可以有“任课”和“辅导”两种联系，称之为复合链。 

两个记录类型之间的值可以是多对多的联系，例如一门课程被多个学生修读，一个学生选修多门课程。

### DBMS数据模型-层次模型

1968年，世界上第一个层次数据库信息管理系统(Information Management System, IMS) 诞生于IBM公司，这也是世界上第一个大型商用的数据库系统。层次数据模型，即使用树形结构来描述实体及其之间关系的数据模型。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122203217.png)

层次数据库就是树结构。每棵树都有且仅有一个根节点，其余的节点都是非根节点。

每个节点表示一个记录类型对应与实体的概念，记录类型的各个字段对应实体的各个属性。各个记录类型及其字段都必须记录。

### DBMS数据模型-关系模型

1970年，IBM的研究员E.F.Codd博士发表了一篇名为"https://www.seas.upenn.edu/~zives/03f/cis550/codd.pdf"的论文，提出了关系模型的概念，奠定了关系模型的理论基础。1979年Oracle首次将 关系型数据库商业化，后续DB2, SAP Sysbase ASE, and Informix等知名数据库产品也纷纷面世。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122203845.png)

使用表格表示实体和实体之间关系的数据模型称之为关系数据模型。 

关系数据模型中，无论是是实体、还是实体之间的联系都是被映射成统一的关系一张二维表，在关系模型中，操作的对象和结果都是一张二维表，它由行和列组成； 

关系型数据库可用于表示实体之间的多对多的关系，只是此时要借助第三个关系—表，来实现多对多的关系；

**DBMS数据模型**

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122204027.png)

1974年ACM牵头组织了一次研讨会，会上开展了一场分别以Codd和Bachman为首的支持和反对关系数据库两派之间的辩论。

这次著名的辩论推动了关系数据库的发展，使其最终成为现代数据库产品的主流。

## SQL语言

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122204403.png)

高度非过程化 非关系数据模型的数据操纵语言是面向过程的语言，用其完成用户请求时，必须指定存取路径。

而用SQL进行数据操作，用户只需提出“做什么”，而不必指明“怎么做”，因此用户无须了解存取路径，存取路径的选择以及SQL语句的操作过程由系统自动完成。这不但大大减轻了用户负担，而且有利于提高数据独立性。 

面向集合的操作方式 SQL采用集合操作方式，不仅查找结果可以是元组的集合，而且一次插入、删除、更新操作的对象也可以是元组的集合。语言简洁，易学易用 SQL功能极强，但由于设计巧妙，语言十分简洁，完成数据定义、数据操纵、数据控制的核心功能只用了9个动词: **CREATE、 ALTER、DROP、 SELECT、 INSERT、 UPDATE、 DELETE、GRANT、 REVOKE**。且SQL语言语法简单，接近英语口语，因此容易学习，也容易使用。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122213355.png)

# 03.关键技术

sql引擎，存储引擎，事务引擎

## 一条SQL的一生

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122214850.png)

SQL引擎 

**查询解析**：SQL 语言接近自然语言，入门容易。但是各种关键字、操作符组合起来，可以表达丰富的语意。因此想要处理SQL命令，首先将文本解析成结构化数据，也就是抽象语法树 （AST）。 

**查询优化**：SQL 是一门表意的语言，只是说『要做什么』，而不说『怎么做』。所以需要一些复杂的逻辑选择『如何拿数据』，也就是选择一个好的查询计划。优化器的作用根据AST优化产生最优执行计划（Plan Tree）。 

**查询执行**：根据查询计划，完成数据读取、处理、写入等操作。 

**事务引擎**：处理事务一致性、并发、读写隔离等 

**存储引擎**：内存中的数据缓存区、数据文件、日志文件



## SQL引擎-Parser
![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122215340.png)

所有的代码在执行之前，都存在一个解析编译的过程，差异点无非在于是静态解析编译还是动态的。 SQL语言也类似，在SQL查询执行前的第一步就是查询解析。 

**词法分析**：将一条SQL语句对应的字符串分割为一个个token，这些token可以简单分类。 

**语法分析**：把词法分析的结果转为语法树。根据token序列匹配不同的语法规则，比如这里匹配的是update语法规则，类似的还有insert、delete、select、create、drop等等语法规则。根据语法规则匹配SQL语句中的关键字，最终输出一个结构化的数据结构。 

**语义分析**：对语法树中的信息进行合法性校验。 

## SQL引擎-Optimizer

为什么需要一个优化器(Optimizer)？

```sql
SELECT * FROM A, B,C WHERE A.a1 = B.b1 and A.a1 = C.b1;
```

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122220419.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122220651.png)



| ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122221007.png) | ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122221026.png) | ![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122221039.png) |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |

比如红绿灯最少，这就是一个规则。 再举个例子，大家知道中关村软件园附近有一条路叫后厂村路，非常的堵。网上有一个广为流传的段子：问：制约中国互联网未来10年发展最大的瓶颈是什么？ 答：后厂村路。那么这里可以加一个规则：不能走后厂村路。 

到达一个目的地，有不同的路线，选择不同的路线有不同的代价。这里的代价可能是时间，也可能是路程。比如我们赶时间的时候，就会选择时间最短的。如果时间没那么赶，那么我们可能选择路程最短的。因为这样省油啊，毕竟现在油价这么高。 对于数据库也是这样，一个查询有不同的执行方案。 那对于数据库而言，什么是一条SQL执行的代价呢？ 其实，对于用户只能感知到查询时间这个代价，底层用了多少资源他是不在乎的。但是在并发的情况下，就得考虑资源消耗了，这个用户的查询占用的资源多了，其他用户的资源就少了。所以资源也是必须考虑的一点。

对于**InnoDB 存储引擎**来说，全表扫描的意思就是把聚簇索引中的记录都依次和给定的搜索条件做一下比较，把符合搜索条件的记录加入到结果集，所以需要将聚簇索引对应的页面加载到内存中，然后再检测记录是否符合搜索条件。 对于使用二级索引 + 回表方式的查询，设计MySQL 的大叔计算这种查询的成本依赖两个方面的数据：**范围区间数量，需要回表数据量**  

## SQL引擎-Executor

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122222753.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122222817.png)

向量化执行更适合于大批量数据处理，对于很多单行数据处理并没有优势。而且往往搭配列式存储使用。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122224352.png)

## 存储引擎-InnoDB

![](https://dev.mysql.com/doc/refman/8.0/en/images/innodb-architecture-8-0.png)

图片来源：https://dev.mysql.com/doc/refman/8.0/en/innodb-architecture.html

In-Memory：

- Buffer Pool
- Change Buffer
- Adaptive Hash Index
- Log Buffer

On-Disk：

- System Tablespace(ibdata1)
- General Tablespacesxxx.ibd)
- Undo Tablespacesl(xx.ibu)
- Temporary Tablespaces(xxx.ibt)
- Redo Log(ib_logfileN)



## 存储引擎-Buffer Pool

MySQL中每个chunk的大小一般为128M，每个block对应一个page，一个chunk下面有8192个block。这样可以避免内存碎片化。

 分成多个instance，可以有效避免并发冲突。 Page id % instance num得到它属于哪个instance

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122225640.png)

当buffer pool里的页面都被使用之后，再需要换存其他页面怎么办？淘汰已有的页面 基于什么规则淘汰：淘汰那个最近一段时间最少被访问过的缓存页了，这种思想就是典型的 LRU 算法了。 

普通的LRU算法存在缺陷，考虑我们需要扫描100GB的表，而我们的buffer pool只有1GB，这样就会因为全表扫描的数据量大，需要淘汰的缓存页多，导致在淘汰的过程中，极有可能将需要频繁使用到的缓存页给淘汰了，而放进来的新数据却是使用频率很低的数据。 

MySQL 确实没有直接使用 LRU 算法，而是在 LRU 算法上进行了优化。 

**MySQL 的优化思路就是**：对数据进行冷热分离，将 LRU 链表分成两部分，一部分用来存放冷数据，也就是刚从磁盘读进来的数据，另一部分用来存放热点数据，也就是经常被访问到数据。 当从磁盘读取数据页后，会先将数据页存放到 LRU 链表冷数据区的头部，如果这些缓存页在 1 秒之后被访问，那么就将缓存页移动到热数据区的头部；如果是 1 秒之内被访问，则不会移动，缓存页仍然处于冷数据区中。 淘汰时，首先淘汰冷数据区。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122225820.png)

![](https://dev.mysql.com/doc/refman/8.0/en/images/innodb-buffer-pool-list.png)

https://dev.mysql.com/doc/refman/8.0/en/innodb-buffer-pool.html

## 存储引擎-Page

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122230927.png)

## 存储引擎-B+ Tree

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122231220.png)

前情提要：如何帮助羊老师从抖音薅一个亿的羊毛?

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122231523.png)

## 事务引擎-Atomicity与Undo Log

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122231953.png)

原子性：一个事务（transaction）中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。

事务在执行过程中发生错误，会被恢复（Rollback）到事务开始前的状态，就像这个事务从来没有执行过一样。 需要记录数据修改前的状态，一边在事务失败时进行回滚。 **undo log是逻辑日志，记录的是数据的增量变化，它的作用是保证事务的原子性和事务并发控制**。可以用于事务回滚，以及提供多版本机制（MVCC），解决读写冲突和一致性读的问题。

## 事务引擎-Isolation与锁

前情提要：羊老师从抖音抢了一个亿红包，又从头条抢了一个亿。抖音和头条都要往羊老师的账户转一个亿， 如果两个操作同时进行，发生冲突怎么办?

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122232157.png)

Isolation（隔离性）：数据库允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致。 

如果多个并发事务访问同一行记录，就需要**锁机制**来保证了。 读写是否冲突？读写互不阻塞，MVCC机制。

## 事务引擎-Isolation与MVCC

**MVCC的意义：**

- 读写互不阻塞；
- 降低死锁概率；
- 实现一致性读。

**Undo Log在MVCC的作用：**

- 每个事务有一个单增的事务ID;
- 数据页的行记录中包含了`DB_ROW_ ID`，`DB_TRX_ID`, `DB_ROLL_PTR`;
  `DB_ROLL_PTR`将数据行的所有快照记录都通过链表的结构串联了起来。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122232830.png)

脏读：事务还没提交之前，它对数据做的修改，不应该被其他人看到。 

万一抖音给我的账户转账的事务还没完成，羊老师就查到了账户上有一个亿，后来抖音发现不对，把这个事务回滚掉了。过一会羊老师发现自己账户的一个亿又没了，去找银行要个说法，结果被保安赶了出来。

## 事务引擎-Durability与Redo Log

如何保证事务结束后，对数据的修改永久的保存？

方案一：事务提交前页面写盘

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122233618.png)



方案二：WAL(Write-ahead logging)

redo log是物理日志，记录的是页面的变化，它的作用是保证事务持久化。如果数据写入磁盘前发生故障，重启MySQL后会根据redo log重做。

持久化：事务处理结束后，对数据的修改就是永久的，即便系统故障也不会丢失。 

WAL：修改并不直接写入到数据库文件中，而是写入到另外一个称为 WAL 的文件中；如果事务失败，WAL 中的记录会被忽略，撤销修改；如果事务成功，它将在随后的某个时间被写回到数据库文件中，提交修改。 

优点： 只记录增量变化，没有写放大 Append only，没有随机IO

# 04.企业实践

## 春节红包雨挑战

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122233933.png)

## 大流量-Sharding

当数据库中的数据量越来越大时，不论是读还是写，压力都会变得越来越大。虽然上面的方案可以扩展读节点，但是对于写流量增加，以及数据量的增加却没有办法。

问题背景

- 单节点写容易成为瓶颈
- 单机数据容量上限

解决方案

- 业务数据进行水平拆分
- 代理层进行分片路由

实施效果

- 数据库写入性能线性扩展
- 数据库容量线性扩展

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122234203.png)

## 流量突增-扩容

问题背景

- 活动流量上涨
- 集群性能不满足要求

解决方案

- 扩容DB物理节点数量
- 利用影子表进行压测

实施效果

- 数据库集群提供更高的吞吐
- 保证集群可以承担预期流量

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122234308.png)

## 流量突增-代理连接池

问题背景

- 突增流量导致大量建联
- 大量建联导致负载变大，延时上升

解决方案

- 业务侧预热连接池
- 代理侧预热连接池
- 代理侧支持连接队列

实施效果

- 避免DB被突增流量打死
- 避免代理和DB被大量建联打死

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122234704.png)

## 稳定性&可靠性

为什么要高可用： 

恶意事故：程序员删库跑路？哪个程序员不想执行一把rm –rf *？ 

偶然事故：如果一个机房断电？断网？ 

某施工队，施工的时候挖掘机把某游戏公司的光纤挖断了，一下午的时间，保守估计损失一个亿。



## 稳定性&可靠性-3AZ高可用

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122234935.png)

BinLog：binlog是mysql用来记录数据库表结构变更以及表数据修改的的二进制日志，它只会记录表的变更操作，但不会记录select和show这种查询操作。 

数据恢复：误删数据之后可以通过mysqlbinlog工具恢复数据 

主从复制：主库将binlog传给从库，从库接收到之后读取内容写入从库，实现主库和从库数据一致性 

审计：可以通过二进制日志中的信息进行审计，判断是否对数据库进行注入攻击

## 稳定性&可靠性-HA管理

问题背景

- db所在机器异常宕机
- db节点异常宕机

解决方案

- ha服务监管、切换宕机节点
- 代理支持配置热加载
- 代理自动屏蔽宕机读节点

实施效果

- 读节点宕机秒级恢复
- 写节点宕机30s内恢复服务

## 总结

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20230122235237.png)















参考文献：

1. https://bytedance.feishu.cn/file/boxcnXzUnOJI7nUFhvBMTW9sfBh
2. https://juejin.cn/post/7101128002909995022#heading-11





