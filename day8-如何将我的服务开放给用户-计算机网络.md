---
title: day8-如何将我的服务开放给用户-计算机网络
date: 2022-05-21 22:02:43
tags: 
- Go学习路线
- 字节跳动青训营
---

![如何将我的服务开放给用户](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220712155538.jpg)

这是我参与「第三届青训营 -后端场」笔记创作活动的的第8篇笔记。*计算机网络真的是太庞大了*

## 「如何将我的服务开放给用户」 第三届字节跳动青训营 - 后端专场

同时这也是课表的第8天课程《如何将我的服务开放给用户》

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220712154741.png)

PC端阅读效果更佳，点击文末：**阅读原文**即可。

## 这节课可以学到什么?

- 系统的熟悉和学习到企业级网络接入核心组件及基本原理
- 当面试时，别人问到你从输入网页到内容加载出来，可以泛泛而谈
- 可以自己从零到一搭建属于自己的网站/博客(网络基础设施)
- 当访问服务出现问题时，可以针对性地进行故障分析及解决

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220712152326.png)


## 1.1问题引入

经典问题：浏览器输入网站域名www.toutiao.com到网页加载出来，都经历了哪些过程?
①域名解析
②TCP建连
③SSL握手
...

可以通过**浏览器抓包**区解析**根源或者本质**  ：www.toutiao.com

![在这里插入图片描述](https://img-blog.csdnimg.cn/be9ddb01fc4f4249953fd0c008542536.png)

## 1.2字节接入框架

A life of a request


![在这里插入图片描述](https://img-blog.csdnimg.cn/31818a6f83454fe1a58d26b38f5724f4.png)
# 02 企业接入升级打怪之路

- 域名系统
- 自建DNS服务器
- HTTPS协议
- 接入全站加速
- 四层负载均衡
- 七层负载均衡

## 2.1使用域名系统
### 2.1.1 Host管理
![在这里插入图片描述](https://img-blog.csdnimg.cn/8e740a6e42bc4733998ad410d4acaa52.png)


随着example公司业务规模和员工数量的增长，使用该方式面临诸多问题: 

- 流量和负载:用户规模指数级增长，文件大小越来越大，统一分 发引起较大的网络流量和cpu负载
- 名称冲突：无法保证主机名称的唯一性， 同名主机添加导致服务故障
- 时效性：分发靠人工上传，时效性太差


### 2.1.2 使用域名系统
使用域名系统替换hosts文件

![](https://img-blog.csdnimg.cn/874ff63543de478aba982d707af64508.png)
关于域名空间:

- 域名空间被组织成树形结构
- 域名空间通过划分zone的方式进行分层授权管理
- 全球公共域名空间仅对应一棵树
- 根域名服务器:查询起点
- 域名组成格式: [a-zA-Z0-9 _-], 以点划分label

顶级域gTLD: general Top-level Domains: gov政府 .edu教育.com商业.mil军事.org非盈利组织

**城名报文格式**

![](https://img-blog.csdnimg.cn/d5a7650c45c543e69a06bfe770704d7b.png)

### 2.1.3 域名购买与配置迁移
首先是域名购买

![](https://img-blog.csdnimg.cn/704f40bc8e734c68901d3423a68961fc.png)
![在这里插入图片描述](https://img-blog.csdnimg.cn/f029e74585e44c8fbf956c1a68692f5e.png)
购买二级域名: example.com
![在这里插入图片描述](https://img-blog.csdnimg.cn/22480bd0921441b59035940394bf04ee.png)
域名备案：防止在网上从事非法的网站经营活动，打击不良互联网信息的传播，一般在云厂商处即可进行实名认证并备案
修改配置：清空/etc/hosts
配置/etc/resolv.conf中nameservers为公共DNS
迁移原配置，通过控制台添加解析记录即可

###  2.1.4 如何开放外部用户访问
>如何建设外部网站，提升公司外部影响力?

方案:租赁一个外网ip,专用于外部用户访问户网站，将www.example.com解析到外网ip
100.1.2.3,将该ip绑定到一台物理机上，并发布公网route,用于外部用户访问。

![](https://img-blog.csdnimg.cn/10abbae6fcff4084a71d13a4b7b037bd.png)
## 2.2 自建DNS服务器
###  2.2.1问题背景
内网域名的解析也得出公网去获取，效率低下

外部用户看到内网ip地址，容易被hacker攻击

云厂商权威DNS容易出故障，影响用户体验

持续扩大公司品牌技术影响力，使用自己的DNS系统

## 2.2.2 DNS查询过程

![](https://img-blog.csdnimg.cn/829b3d13167c415897d5699a262a084c.png)


dig {$domain} +trace

dig 命令主要用来从 DNS 域名服务器查询主机地址信息,可以用来测试域名系统工作是否正常。

参数：

主机：指定要查询域名主机；
查询类型：指定DNS查询的类型；
查询类：指定查询DNS的class；
查询选项：指定查询选项。

@<服务器地址>：指定进行域名解析的域名服务器；
-b<ip地址>：当主机具有多个IP地址，指定使用本机的哪个IP地址向域名服务器发送域名查询请求；
-f<文件名称>：指定dig以批处理的方式运行，指定的文件中保存着需要批处理查询的DNS任务信息；
-P：指定域名服务器所使用端口号；
-t<类型>：指定要查询的DNS数据类型；
-x<IP地址>：执行逆向域名查询；
-4：使用IPv4；
-6：使用IPv6；
-h：显示指令帮助信息。
dig @8.8.8.8 www.baidu.com #指定DNS服务器解析,比如
```sh
[root@jenkins ~]# dig @223.5.5.5 www.baidu.com

; <<>> DiG 9.11.4-P2-RedHat-9.11.4-9.P2.el7 <<>> @223.5.5.5 www.baidu.com
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 18719
;; flags: qr rd ra; QUERY: 1, ANSWER: 3, AUTHORITY: 0, ADDITIONAL: 0

;; QUESTION SECTION:
;www.baidu.com.			IN	A

;; ANSWER SECTION:
www.baidu.com.		15	IN	CNAME	www.a.shifen.com.
www.a.shifen.com.	15	IN	A	220.181.38.150
www.a.shifen.com.	15	IN	A	220.181.38.149

;; Query time: 6 msec
;; SERVER: 223.5.5.5#53(223.5.5.5)
;; WHEN: Tue Oct 22 02:20:06 EDT 2019
;; MSG SIZE  rcvd: 90
dig www.baidu.com +short #简洁查询
dig +nocmd +noall +answer +ttlid www.baidu.com #更简洁查询
dig +trac www.baidu.com   #DNS请求的递归查询过程
dig +trace -t A @8.8.8.8 taobao.com #指定A记录解析
dig www.baidu.com CNAME #查询域名的CNAME
dig -x 8.8.8.8 #反查询，通过IP查询域名
dig -x 8.8.8.8 +short 
```

类型	目的
A	地址记录，用来指定域名的 IPv4 地址，如果需要将域名指向一个 IP 地址，就需要添加 A 记录。
AAAA	用来指定主机名(或域名)对应的 IPv6 地址记录。
CNAME	如果需要将域名指向另一个域名，再由另一个域名提供 ip 地址，就需要添加 CNAME 记录。
MX	如果需要设置邮箱，让邮箱能够收到邮件，需要添加 MX 记录。
NS	域名服务器记录，如果需要把子域名交给其他 DNS 服务器解析，就需要添加 NS 记录。
SOA	SOA 这种记录是所有区域性文件中的强制性记录。它必须是一个文件中的第一个记录。
TXT	可以写任何东西，长度限制为 255。绝大多数的 TXT记录是用来做 SPF 记录(反垃圾邮件)。

### 2.2.3 DNS记录类型
A/AAAA: IP指向记录，用于指向IP，前者为IPv4记录，后者为IPv6记录
CNAME:别名记录，配置值为别名或主机名，客户端根据别名继续解析以提取IP地址
TXT:文本记录，购买证书时需要
MX:邮件交换记录，用于指向邮件交换服务器
NS:解析服务器记录，用于指定哪台服务器对于该域名解析

SOA记录:起始授权机构记录，每个zone有 且仅有唯一-的一 条SOA记录，SOA是描述zone属性以及主要权威服务器的记录
![](https://img-blog.csdnimg.cn/c3aaac0039624c70968c5660561ba055.png)
![在这里插入图片描述](https://img-blog.csdnimg.cn/6b09abe51d904d2695a798906e53c4c2.png)


### 2.2.4权威DNS系统架构
思考:站在企业角度思考，我们需要的是哪种DNS服务器?

答案：权威DNS，LocalDNS(可选)
常见的开源DNS: bind、 nsd、 knot、coredns
![](https://img-blog.csdnimg.cn/89b9791ae4c9478aa894703e8df6fc89.png)

DNS Query
DNS Response
DNS Update
DNS Notify
DNS XFR

---

经过研发人员的不断努力，example 公司有了自己的权威DNS系统。
![](https://img-blog.csdnimg.cn/50587fe406de4de396f31839e3f6de0e.png)


## 2.3 接入HTTPS协议
### 2.3.1问题背景

- 页面出现白页/出现某些奇怪的东西
- 返回了403的页面
- 搜索不了东西
- 搜索问题带了小尾巴，页面总要闪几次
- 页面弹窗广告
- 搜索个汽车就有人给我打电话推销4s店和保险什么的
.....




![](https://img-blog.csdnimg.cn/61d7f2e8a6a04957bffa0481fd6b4c2c.png)
### 2.3.2对称加密和非对称加密

常见的加密算法
对称加密：一份秘钥
![](https://img-blog.csdnimg.cn/5876788d40a04d22b5838b6088471134.png)


非对称加密:公钥和私钥

![](https://img-blog.csdnimg.cn/57af077b1bce4b6b9d6fea1055b20182.png)
### 2.3.3 SSL的通信过程
- client random
- server random
- premaster secret
- 加密算法协商 --> 对称秘钥session key
![](https://img-blog.csdnimg.cn/c903ec6967dc4e3eb1b9dc5d623e253b.png)


### 2.3.4证书链
公钥确定是可信的吗?会不会被劫持?

Server端发送是带签名的证书链(下图)

![](https://img-blog.csdnimg.cn/4dac4c897a3b45a2bb0f07cd3e955cf3.png)

Client收到会仍然需要验证:
- 是否是可信机构颁布
- 域名是否与实际访问一致
- 检查数字签名是否一致
- 检查证书的有效期
- 检查证书的撤回状态

![](https://img-blog.csdnimg.cn/c24a341c09734b9baeda4d905847b84b.png)

### 2.3.5使用https
![](https://img-blog.csdnimg.cn/df41cbaf491d4847b88bc2f91a29b445.png)


## 2.4接入全站加速

### 2.4.1问题背景
**外网用户访问站点，一定是一帆风顺的吗?可能出现的问题有哪些?**
![](https://img-blog.csdnimg.cn/d553d0958c3e4699bc4169866b6d437f.png)
源站容量低，可承载的并发请求数低，容易被打垮
报文经过的网络设备越多，出问题的概率越大，丢包、劫持、mtu问题
自主选路网络链路长，时延高

**响应慢、卡顿**
如果请求一个网页响应超过3秒，80%的用户就会找其他产品进行替代。就算做的再好也没用
极大的流失了大部分的用户群体，NPS留存率数据不乐观。


### 2.4.2解决方案
源站容量问题：增加后端机器扩容;静态内容，使用静态加速缓存
网络传输问题：动态加速DCDN
全站加速：静态加速+动态加速

### 2.4.3静态加速CDN

针对静态文件传输，网络优化方式?

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220706184914.png)

加缓存（CDN）
![](https://img-blog.csdnimg.cn/146b8945f5b7444196fffd5071116052.png)

### 2.4.3静态加速CDN
解决服务器端的“第一公里”问题

缓解甚至消除了不同运营商之间互联的瓶颈造成的影响

减轻了各省的出口带宽压力


优化了网上热点内容的分布
### 2.4.4动态加速DCDN
针对POST等非静态请求等不能在用户边缘缓存的业务，基于智能选路技术，从众多回源线路中择优选择一条线路进行传输 。

- 用户发起动态请求
- 智能选择性能与稳定性最优路径
- 动态请求通过最优路径快速回源


![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220712165901.gif)
### 2.4.5 DCDN原理
RTT示例:
- 用户到核心: 35ms
- 用户到边缘: 20ms
- 边缘到汇聚: 10ms
- 汇聚到核心: 10ms
![](https://img-blog.csdnimg.cn/3059b24dab2d46bd93fdd5c70bad6e56.png)


常规请求耗时计算:
Via DCDN: 100ms

20(TCP)+20*2(TLS)+20+ 10+ 10(routine)

Direct: 140ms

35(TCP)+35*2(TLS)+35(routine)

### 2.4.6使用全站加速
请区分下列场景使用的加速类型
1. 用户首次登录抖音，注册用户名手机号等用户信息  **动态加速DCDN**

2. 抖音用户点开某个特定的短视频加载后观看 **静态加速CDN**
3. 用户打开头条官网进行网页浏览       **静态加速CDN+动态加速DCDN**

### 2.4.6 使用全站加速
![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220712165247.png)

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/20220712165351.png)

## 2.5  4层负载均衡

### 2.5.1问题背景

>提问:在运营商处租用的100.1.2.3的公网IP，如何在企业内部使用最合理?

现状：直接找一个物理机，ifconfig将网卡配上这个IP，起server监听即可

应用多，起多个server监听不同的端口即可
租多个公网ip (数量有限)
>怎样尽可能充分的利用和管理有限的公网IP资源?

### 2.5.2什么是4层负载均衡?

基于IP+端口，利用某种算法将报文转发给某个后端服务器，实现负载均衡地落到后端服务器上。

![](https://img-blog.csdnimg.cn/bdcd53bb6fd14ec285cdbdcb39b26baf.png)

### 2.5.3常见的调度算法原理
- RR轮询: Round Robin,将所有的请求平均分配给每个真实服务器RS
- 加权RR轮询：给每个后端服务器一个权值比例， 将请求按照比例分配

- 最小连接：把新的连接请求分配到当前连接数最小的服务器



- 五元组hash：根据sip、 sport、 proto、 dip、 dport对静态分配的服务器做散列取模

- 缺点：当后端某个服务器故障后，所有连接都重新计算，影响整个hash环
- 一致性hash: 只影响故障服务器上的连接session,其余服务器上的连接不受影响

### 2.5.4常见的实现方式FULL .NAT

![](https://img-blog.csdnimg.cn/1cfa52632c9e4667a8d2d950c0c7622f.png)
>提问：RS怎么知道真实的CIP?
>
>回答：通过TCP option字段传递然后通过特殊的内核模块反解

### 2.5.5 4层负载均衡特点
大部分都是通过dpdk技术实现，技术成熟，大厂都在用
纯用户态协议栈，kernel bypass,消除协议栈瓶颈
无缓存，零拷贝，大页内存(减少cache miss)
仅针对4层数据包转发，小包转发可达到限速，可承受高cps

### 2.5.6使用4层负载均衡
![](https://img-blog.csdnimg.cn/23400b86a1024a0b90bbdb64a3f0cb05.png)

# 2.6  7层负载均衡
### 2.6.1问题背景

>提问：四层负载对100.1.2.3只能bind一个80端口，而有多个外部站点需要使用，该如何解决?

>换个问法:有一些7层相关的配置需求，该怎么做?

SSL卸载:业务侧是http服务，用户需要用https访问

请求重定向:浏览器访问toutiao.com自动跳转www.toutiao.com

路由添加匹配策略:完全、前缀、正则
Header编辑
跨域支持
协议支持: websocket、 grpc、 quic

### 2.6.2 Nginx简介
**最灵活的高性能WEB SERVER,应用最广的7层反向代理。**

![](https://img-blog.csdnimg.cn/289b319456c24c4db7a587ff6f3a17ea.png)
### 2.6.3 Nginx和Apache性能对比

  

![](https://img-blog.csdnimg.cn/db0770d61af44c3bb1941e6c7c22fae8.png)

- 模块化设计，较好的扩展性和可靠性
- 基于master/worker架构设计
- 支持热部署;可在线升级
- 不停机更新配置文件、更换日志文件、更新服务器二进制
- 较低的内存消耗: 1万个keep- -alive 连接模式下的非活动连接仅消耗2.5M内存
- 事件驱动:异步非阻塞模型、支持aio, mmap (内存映射)

### 2.6.4 Nginx反向代理示意图
代理服务器功能
- Keepalive
- 访问日志
- ourl rewrite重写
- 路径别名
- 基于ip的用户的访问控制
- 限速及并发连接数控制

![](https://img-blog.csdnimg.cn/b4bb1be88f2b410680e5bb2ea5b349e8.png)

### 2.6.5 Nginx内部架构
![](https://img-blog.csdnimg.cn/ccf864b5d2f849dc9c89a9c30ed34691.png)


### 2.6.6 事件驱动模型
![](https://img-blog.csdnimg.cn/eff29ee6ee894b89b874cfaafb7f3741.png)

### 2.6.7 异步非阻塞

传统服务器：一个进程/线程处理一个连接/请求。阻塞模型、依赖OS实现并发
Nginx: 个进程/线程处理多个连接/请求。异步非阻塞模型、减少OS进程切换
![](https://img-blog.csdnimg.cn/accd089f2f8747288565c5c29ccb1ef8.png)


### 2.6.8 Nginx简单调优
![](https://img-blog.csdnimg.cn/b3f1600ddab24a43908c90a04d81c244.png)

### 2.6.8 别让OS限制了Nginx的性能
优化内核网络参数


fs.filemax= 999999
net.ipv4.tcp_ tw_ reuse = 1
net.ipv4.tcp_ keepalive_ _time = 600
net.ipv4.tcp_ fin_ timeout = 30
net.ipv4.tcp_ max_tw_ buckets = 5000
net.ipv4.ip_ local_port_ range = 1024 61000
net.ipv4.tcp_ max_ syn.backlog= 1024
net.ipv4.tcp_ syncookies = 1


### 2.6.8提升CPU使用效率
![](https://img-blog.csdnimg.cn/c670bf2ea4054bd9baaeb405a20fa1f8.png)

### 2.6.9提升网络效率
![](https://img-blog.csdnimg.cn/a7f5cdeeef6f4949993909a356754bf0.png)

### 2.6.10使用7层负载均衡
![](https://img-blog.csdnimg.cn/a24564e06b9441e6a499febebe2d2355.png)
## 03.动手实践
#### 3.1 DNS服务器搭建
![](https://img-blog.csdnimg.cn/abcad42a023d4f21befea7a09bcb266a.png)
![在这里插入图片描述](https://img-blog.csdnimg.cn/652c1d0d60f643aeb4d98107eef856a8.png)

### 3.2四层负载均衡实验
开源的解决方案: LVS+keepalived
LVS: linux virtual server, linux虚拟服务器， 根据目标地址和目标端口实现用户请求转发，本身不产生流量，只
做用户请求转发，详见http://www.linuxvirtualserver.org/

![](https://img-blog.csdnimg.cn/dc64a16cd0e84eeb966be93c2444378e.png)![在这里插入图片描述](https://img-blog.csdnimg.cn/98f49327ad1745288916a6f34380bba0.png)
![](https://img-blog.csdnimg.cn/b105e1e2c62d409aa2ea5bfa18648589.png)


### 3.3  7层负载均衡实验
![](https://img-blog.csdnimg.cn/d74d48b8f304460e93cebfa46be7278f.png)


### 3.4 SSL自签证书实验
![](https://img-blog.csdnimg.cn/fc61286b061c4900867750e9af39deea.png)


### 3.5如何将本地服务开放外网访问
>提问:服务开发前期，如何低成本的让别人访问自己的服务?


回答: Ngrok, Expose your localhost to the web
使用条件:使用github账户授权登录，即可使用，详见https://dashboard.ngrok.com/get-started/setup

![](https://img-blog.csdnimg.cn/dca8a5bbf6904430a2585e9555aa9dd3.png)

### 3.5如何将本地服务开放外网访问
命令: ./ngrok http example.com:8082

![](https://img-blog.csdnimg.cn/9aff3f79e19e4576a4c0f0f51b26658f.png)
再看接入架构
![](https://img-blog.csdnimg.cn/2ee99ca75a654263bfb5b7848153d871.png)

### 总结
![](https://img-blog.csdnimg.cn/b2582827465042a997839ba3e6d604fc.png)
### 大作业
范本: https://strikefreedom.top/
搭建一个专属的个人网站:奇闻轶事、个人杜撰、美食、宠物、技术博客。

1. https://www.volcengine.com/
2.  https://www.aliyun.com/
3.  https://www.cnblogs.com/ximu-xin/p/8726815.html
4.  https://segmentfault.com/a/ 1190000021494676
5.  https://blog.cloudflare.com/keyless -ssl-the nitty gritty- technical-details/
6.  https://www.jianshu.com/p/1dae6e1680ff
7.  https://www.aliyun.com/
8.  https://w3techs.com/technologies/overview/web_server
9.  https://help.dreamhost.com/hc/en-us/articles/215945987-Web-server-performance comparison
10.  http://www.4k8k.xyz/article/mss359681091/85318254
11.  https://wwwjianshu.com/p/3d9e45082d27
12.  https://help.aliyun.com/document_detail/27544.html

