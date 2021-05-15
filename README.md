[TOC]

<center><h1>Go语言基础实战</h1></center>

### 1.1.1. Go语言为并发而生

go语言（或 Golang）是Google开发的开源编程语言，诞生于2006年1月2日下午15点4分5秒，于2009年11月开源，2012年发布go稳定版。Go语言在多核并发上拥有原生的设计优势，Go语言从底层原生支持并发，无须第三方库、开发者的编程技巧和开发经验。

go是非常年轻的一门语言，它的主要目标是“兼具Python 等动态语言的开发速度和C/C++等编译型语言的性能与安全性”

很多公司，特别是中国的互联网公司，即将或者已经完成了使用 Go 语言改造旧系统的过程。经过 Go 语言重构的系统能使用更少的硬件资源获得更高的并发和I/O吞吐表现。充分挖掘硬件设备的潜力也满足当前精细化运营的市场大环境。

Go语言的并发是基于 `goroutine` 的，`goroutine` 类似于线程，但并非线程。可以将 `goroutine` 理解为一种虚拟线程。Go 语言运行时会参与调度 `goroutine`，并将 `goroutine` 合理地分配到每个 CPU 中，最大限度地使用CPU性能。开启一个goroutine的消耗非常小（大约2KB的内存），你可以轻松创建数百万个`goroutine`。

`goroutine`的特点：

1. `goroutine`具有可增长的分段堆栈。这意味着它们只在需要时才会使用更多内存。
2. `goroutine`的启动时间比线程快。
3. `goroutine`原生支持利用channel安全地进行通信。
4. `goroutine`共享数据结构时无需使用互斥锁。

### 1.1.2. Go语言简单易学

#### 语法简洁

Go 语言简单易学，学习曲线平缓，不需要像 C/C++ 语言动辄需要两到三年的学习期。Go 语言被称为“互联网时代的C语言”。Go 语言的风格类似于C语言。其语法在C语言的基础上进行了大幅的简化，去掉了不需要的表达式括号，循环也只有 for 一种表示方法，就可以实现数值、键值等各种遍历。

#### 代码风格统一

Go 语言提供了一套格式化工具——go fmt。一些 Go 语言的开发环境或者编辑器在保存时，都会使用格式化工具进行修改代码的格式化，这样就保证了不同开发者提交的代码都是统一的格式。(吐槽下：再也不用担心那些看不懂的黑魔法了…)

#### 开发效率高

![开发效率高图](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/63371a7e4dd2e70f0f1f22d6981b2245.jpeg)

Go语言实现了开发效率与执行效率的完美结合，让你像写Python代码（效率）一样编写C代码（性能）。

### 1.1.3. 使用go的公司

- Google
  - https://github.com/kubernetes/kubernetes
- Facebook
  - https://github.com/facebookgo
- 腾讯
- 百度
- 360开源日志系统
  - https://github.com/Qihoo360/poseidon

### 1.1.4. go适合做什么

- 服务端开发
- 分布式系统，微服务
- 网络编程
- 区块链开发
- 内存KV数据库，例如boltDB、levelDB
- 云平台

### 1.1.5. 学习Go语言的前景

目前Go语言已经⼴泛应用于人工智能、云计算开发、容器虚拟化、⼤数据开发、数据分析及科学计算、运维开发、爬虫开发、游戏开发等领域。

Go语言简单易学，天生支持并发，完美契合当下高并发的互联网生态。Go语言的岗位需求持续高涨，目前的Go程序员数量少，待遇好。

抓住趋势，要学会做一个领跑者而不是跟随者。

国内Go语言的需求潜力巨大，目前无论是国内大厂还是新兴互联网公司基本上都会有Go语言的岗位需求。

基本上，越来越多的大厂在招Go，未来2023-2025会过渡到一个转折点。所以，一起开始学习Go吧！

# 开发环境

### [go的安装](https://www.bookstack.cn/read/topgoer/1d3381d29db04cb4.md)

### [配置gopath](https://www.bookstack.cn/read/topgoer/aa46a6a64197bfac.md)

### [编辑器](https://www.bookstack.cn/read/topgoer/40a82b857ad55d0c.md)

### [git安装](https://www.bookstack.cn/read/topgoer/679e5d562083142a.md)

### [第一个go程序](https://www.bookstack.cn/read/topgoer/699e5c205f57701c.md)

# Go基础

### [go语言的主要特征](https://www.bookstack.cn/read/topgoer/9d74dc8391e5e525.md)

### 1.1. golang 简介

### 1.1.1. 来历

很久以前，有一个IT公司，这公司有个传统，允许员工拥有20%自由时间来开发实验性项目。在2007的某一天，公司的几个大牛，正在用c++开发一些比较繁琐但是核心的工作，主要包括庞大的分布式集群，大牛觉得很闹心，后来c++委员会来他们公司演讲，说c++将要添加大概35种新特性。这几个大牛的其中一个人，名为：Rob Pike，听后心中一万个xxx飘过，“c++特性还不够多吗？简化c++应该更有成就感吧”。于是乎，Rob Pike和其他几个大牛讨论了一下，怎么解决这个问题，过了一会，Rob Pike说要不我们自己搞个语言吧，名字叫“go”，非常简短，容易拼写。其他几位大牛就说好啊，然后他们找了块白板，在上面写下希望能有哪些功能（详见文尾）。接下来的时间里，大牛们开心的讨论设计这门语言的特性，经过漫长的岁月，他们决定，以c语言为原型，以及借鉴其他语言的一些特性，来解放程序员，解放自己，然后在2009年，go语言诞生。

### 1.1.2. 思想

Less can be more 大道至简,小而蕴真 让事情变得复杂很容易，让事情变得简单才难 深刻的工程文化

### 1.1.3. 优点

自带gc。

静态编译，编译好后，扔服务器直接运行。

简单的思想，没有继承，多态，类等。

丰富的库和详细的开发文档。

语法层支持并发，和拥有同步并发的channel类型，使并发开发变得非常方便。

简洁的语法，提高开发效率，同时提高代码的阅读性和可维护性。

超级简单的交叉编译，仅需更改环境变量。

Go 语言是谷歌 2009 年首次推出并在 2012 年正式发布的一种全新的编程语言，可以在不损失应用程序性能的情况下降低代码的复杂性。谷歌首席软件工程师罗布派克(Rob Pike)说：我们之所以开发 Go，是因为过去10多年间软件开发的难度令人沮丧。Google 对 Go 寄予厚望，其设计是让软件充分发挥多核心处理器同步多工的优点，并可解决面向对象程序设计的麻烦。它具有现代的程序语言特色，如垃圾回收，帮助开发者处理琐碎但重要的内存管理问题。Go 的速度也非常快，几乎和 C 或 C++ 程序一样快，且能够快速开发应用程序。

### 1.1.4. Go语言的主要特征：

```go
    1.自动立即回收。    2.更丰富的内置类型。    3.函数多返回值。    4.错误处理。    5.匿名函数和闭包。    6.类型和接口。    7.并发编程。    8.反射。    9.语言交互性。
```

### 1.1.5. Golang文件名：

```go
`所有的go源码都是以 ".go" 结尾。`
```

### 1.1.6. Go语言命名：

1.Go的函数、变量、常量、自定义类型、包`(package)`的命名方式遵循以下规则：

```go
    1）首字符可以是任意的Unicode字符或者下划线    2）剩余字符可以是Unicode字符、下划线、数字    3）字符长度不限
```

2.Go只有25个关键字

```go
    break        default      func         interface    select    case         defer        go           map          struct    chan         else         goto         package      switch    const        fallthrough  if           range        type    continue     for          import       return       var
```

3.Go还有37个保留字

```go
    Constants:    true  false  iota  nil    Types:    int  int8  int16  int32  int64                uint  uint8  uint16  uint32  uint64  uintptr              float32  float64  complex128  complex64              bool  byte  rune  string  error    Functions:   make  len  cap  new  append  copy  close  delete                 complex  real  imag                 panic  recover
```

4.可见性：

```go
    1）声明在函数内部，是函数的本地值，类似private    2）声明在函数外部，是对当前包可见(包内所有.go文件都可见)的全局值，类似protect    3）声明在函数外部且首字母大写是所有包可见的全局值,类似public
```

### 1.1.7. Go语言声明：

有四种主要声明方式：

```go
    var（声明变量）, const（声明常量）, type（声明类型） ,func（声明函数）。
```

Go的程序是保存在多个.go文件中，文件的第一行就是package XXX声明，用来说明该文件属于哪个包(package)，package声明下来就是import声明，再下来是类型，变量，常量，函数的声明。

### 1.1.8. Go项目构建及编译

一个Go工程中主要包含以下三个目录：

```
    src：源代码文件    pkg：包文件    bin：相关bin文件
```

1: 建立工程文件夹 goproject

2: 在工程文件夹中建立src,pkg,bin文件夹

3: 在GOPATH中添加projiect路径 例 e:/goproject

4: 如工程中有自己的包examplepackage，那在src文件夹下建立以包名命名的文件夹 例 examplepackage

5：在src文件夹下编写主程序代码代码 goproject.go

6：在examplepackage文件夹中编写 examplepackage.go 和 包测试文件 examplepackage_test.go

7：编译调试包

go build examplepackage

go test examplepackage

go install examplepackage

这时在pkg文件夹中可以发现会有一个相应的操作系统文件夹如windows_386z, 在这个文件夹中会有examplepackage文件夹，在该文件中有examplepackage.a文件

8：编译主程序

go build goproject.go

成功后会生成goproject.exe文件

至此一个Go工程编辑成功。

```
1.建立工程文件夹 go$ pwd/Users/***/Desktop/go2: 在工程文件夹中建立src,pkg,bin文件夹$ lsbin        conf    pkg        src3: 在GOPATH中添加projiect路径$ go envGOPATH="/Users/liupengjie/Desktop/go"4: 那在src文件夹下建立以自己的包 example 文件夹$ cd src/$ mkdir example5：在src文件夹下编写主程序代码代码 goproject.go6：在example文件夹中编写 example.go 和 包测试文件 example_test.go    example.go 写入如下代码：    package example    func add(a, b int) int {        return a + b    }    func sub(a, b int) int {        return a - b    }    example_test.go 写入如下代码：    package example    import (        "testing"    )    func TestAdd(t *testing.T) {        r := add(2, 4)        if r != 6 {            t.Fatalf("add(2, 4) error, expect:%d, actual:%d", 6, r)        }        t.Logf("test add succ")    }7：编译调试包    $ go build example    $ go test example    ok      example    0.013s    $ go install example$ ls /Users/***/Desktop/go/pkg/darwin_amd64$ ls /Users/***/Desktop/go/pkg/darwin_amd64/example.a    8：编译主程序    oproject.go 写入如下代码：    package main     import (        "fmt"    )    func main(){        fmt.Println("go project test")    }    $ go build goproject.go    $ ls    example        goproject.go    goproject       成功后会生成goproject文件    至此一个Go工程编辑成功。       运行该文件：    $ ./goproject    go project test
```

### 1.1.9. go 编译问题

golang的编译使用命令 go build , go install;除非仅写一个main函数，否则还是准备好目录结构； GOPATH=工程根目录；其下应创建src，pkg，bin目录，bin目录中用于生成可执行文件，pkg目录中用于生成.a文件； golang中的import name，实际是到GOPATH中去寻找name.a, 使用时是该name.a的源码中生命的package 名字；这个在前面已经介绍过了。

注意点：

```
 复制代码    1.系统编译时 go install abc_name时，系统会到GOPATH的src目录中寻找abc_name目录，然后编译其下的go文件；    2.同一个目录中所有的go文件的package声明必须相同，所以main方法要单独放一个文件，否则在eclipse和liteide中都会报错；    编译报错如下：（假设test目录中有个main.go 和mymath.go,其中main.go声明package为main，mymath.go声明packag 为test);        $ go install test        can't load package: package test: found packages main (main.go) and test (mymath.go) in /home/wanjm/go/src/test        报错说 不能加载package test（这是命令行的参数），因为发现了两个package，分别时main.go 和 mymath.go;    3.对于main方法，只能在bin目录下运行 go build path_tomain.go; 可以用-o参数指出输出文件名；    4.可以添加参数 go build -gcflags "-N -l" ****,可以更好的便于gdb；详细参见 http://golang.org/doc/gdb    5.gdb全局变量主一点。 如有全局变量 a；则应写为 p 'main.a'；注意但引号不可少；
```

# [Golang内置类型和函数](https://www.bookstack.cn/read/topgoer/e7224bac95ecb9c4.md)

[init函数和main函数](https://www.bookstack.cn/read/topgoer/afdebff8cf12daeb.md)

# [命令](https://www.bookstack.cn/read/topgoer/419d423b7a2349cb.md)

假如你已安装了golang环境，你可以在命令行执行go命令查看相关的Go语言命令：

```go
$ goGo is a tool for managing Go source code.Usage:   
go command [arguments]The commands are:    
build       compile packages and dependencies    
clean       remove object files    
doc         show documentation for package or symbol    
env         print Go environment information    
bug         start a bug report    fix         run go tool fix on packages    fmt         run gofmt on package sources    generate    generate Go files by processing source    get         download and install packages and dependencies    install     compile and install packages and dependencies    list        list packages    run         compile and run Go program    test        test packages    tool        run specified go tool    version     print Go version    vet         run go tool vet on packagesUse "go help [command]" for more information about a command.Additional help topics:    c           calling between Go and C    buildmode   description of build modes    filetype    file types    gopath      GOPATH environment variable    environment environment variables    importpath  import path syntax    packages    description of package lists    testflag    description of testing flags    testfunc    description of testing functionsUse "go help [topic]" for more information about that topic.
```

go env用于打印Go语言的环境信息。

go run命令可以编译并运行命令源码文件。

go get可以根据要求和实际情况从互联网上下载或更新指定的代码包及其依赖包，并对它们进行编译和安装。

go build命令用于编译我们指定的源码文件或代码包以及它们的依赖包。

go install用于编译并安装指定的代码包及它们的依赖包。

go clean命令会删除掉执行其它命令时产生的一些文件和目录。

go doc命令可以打印附于Go语言程序实体上的文档。我们可以通过把程序实体的标识符作为该命令的参数来达到查看其文档的目的。

go test命令用于对Go语言编写的程序进行测试。

go list命令的作用是列出指定的代码包的信息。

go fix会把指定代码包的所有Go语言源码文件中的旧版本代码修正为新版本的代码。

go vet是一个用于检查Go语言源码中静态错误的简单工具。

go tool pprof命令来交互式的访问概要文件的内容。

# [运算符](https://www.bookstack.cn/read/topgoer/b888afdd94ed0c4d.md)

Go 语言内置的运算符有：

```go
    算术运算符    关系运算符    逻辑运算符    位运算符    赋值运算符
```

### 1.1.1. 算数运算符

| 运算符 | 描述 |
| :----: | :--: |
|   +    | 相加 |
|   -    | 相减 |
|   *    | 相乘 |
|   /    | 相除 |
|   %    | 求余 |

注意： ++（自增）和—（自减）在Go语言中是单独的语句，并不是运算符。

### 1.1.2. 关系运算符

| 运算符 | 描述                                                         |
| :----- | :----------------------------------------------------------- |
| ==     | 检查两个值是否相等，如果相等返回 True 否则返回 False。       |
| !=     | 检查两个值是否不相等，如果不相等返回 True 否则返回 False。   |
| >      | 检查左边值是否大于右边值，如果是返回 True 否则返回 False。   |
| >=     | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 |
| <      | 检查左边值是否小于右边值，如果是返回 True 否则返回 False。   |
| <=     | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 |

### 1.1.3. 逻辑运算符

| 运算符 | 描述                                                         |
| :----- | :----------------------------------------------------------- |
| &&     | 逻辑 AND 运算符。 如果两边的操作数都是 True，则为 True，否则为 False。 |
| ll     | 逻辑 OR 运算符。 如果两边的操作数有一个 True，则为 True，否则为 False。 |
| !      | 逻辑 NOT 运算符。 如果条件为 True，则为 False，否则为 True。 |

### 1.1.4. 位运算符

位运算符对整数在内存中的二进制位进行操作。

| 运算符 | 描述                                                         |
| :----- | :----------------------------------------------------------- |
| &      | 参与运算的两数各对应的二进位相与。（两位均为1才为1）         |
| l      | 参与运算的两数各对应的二进位相或。（两位有一个为1就为1）     |
| ^      | 参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。（两位不一样则为1） |
| <<     | 左移n位就是乘以2的n次方。“a<<b”是把a的各二进位全部左移b位，高位丢弃，低位补0。 |
| >>     | 右移n位就是除以2的n次方。“a>>b”是把a的各二进位全部右移b位。  |

### 1.1.5. 赋值运算符

| 运算符 | 描述                                           |
| :----- | :--------------------------------------------- |
| =      | 简单的赋值运算符，将一个表达式的值赋给一个左值 |
| +=     | 相加后再赋值                                   |
| -=     | 相减后再赋值                                   |
| *=     | 相乘后再赋值                                   |
| /=     | 相除后再赋值                                   |
| %=     | 求余后再赋值                                   |
| <<=    | 左移后赋值                                     |
| >>=    | 右移后赋值                                     |
| &=     | 按位与后赋值                                   |
| l=     | 按位或后赋值                                   |
| ^=     | 按位异或后赋值                                 |



[下划线_](https://www.bookstack.cn/read/topgoer/33b665c029b860e7.md)

[变量和常量](https://www.bookstack.cn/read/topgoer/1197f69f184b79a5.md)

[基本类型](https://www.bookstack.cn/read/topgoer/00a95e3bc412b83c.md)

[数组array](https://www.bookstack.cn/read/topgoer/dbe7937cbfb0004b.md)

[切片slice](https://www.bookstack.cn/read/topgoer/ae99b7333deacee2.md)

[指针](https://www.bookstack.cn/read/topgoer/b0730e156e319023.md)

[map](https://www.bookstack.cn/read/topgoer/69b4f2adbd6bfe84.md)

[结构体](https://www.bookstack.cn/read/topgoer/d433c0989c52707c.md)



































​	

## 介绍

记录千羽学Go编程时光

#### 2. 四种变量的声明方式

> 变量声明以关键字`var`开头，变量类型放在变量的后面，行尾无需分号。 举个例子：
>
> 1. `var name string`
> 2. `    var age int`
> 3.  var isOk bool

```go
package main

import (
	"fmt"
)

//声明全局变量 方法一、方法二、方法三是可以的
var gA int = 100
var gB = 200

//用方法四来声明全局变量
// := 只能够用在 函数体内来声明
//gC := 200

func main() {
	//方法一：声明一个变量 默认的值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	//方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	//方法三：在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	//方法四：(常用的方法) 省去var关键字，直接自动匹配
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := "abcd"
	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	// =====
	fmt.Println("gA = ", gA, ", gB = ", gB)
	//fmt.Println("gC = ", gC)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, ", yy = ", yy)
	var kk, ll = 100, "Aceld"
	fmt.Println("kk = ", kk, ", ll = ", ll)

	//多行的多变量声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, ", jj = ", jj)
}
```

#### 3. 常量

```go
package main

import "fmt"

//const 来定义枚举类型
const (
	//可以在const() 添加一个关键字 iota， 每行的iota都会累加1, 第一行的iota的默认值是0
	BEIJING = 10*iota	 //iota = 0
	SHANGHAI 		  //iota = 1
	SHENZHEN          //iota = 2
)

const (
	a, b = iota+1, iota+2 // iota = 0, a = iota + 1, b = iota + 2, a = 1, b = 2
	c, d				  // iota = 1, c = iota + 1, d = iota + 2, c = 2, d = 3
	e, f				  // iota = 2, e = iota + 1, f = iota + 2, e = 3, f = 4

	g, h = iota * 2, iota *3  // iota = 3, g = iota * 2, h = iota * 3, g = 6, h = 9 
	i, k					   // iota = 4, i = iota * 2, k = iota * 3 , i = 8, k = 12
)

func main() {
	//常量(只读属性)
	const length int = 10

	fmt.Println("length = ", length)

	//length = 100 //常量是不允许修改的。

	fmt.Println("BEIJIGN = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)

	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "k = ", k)

	// iota 只能够配合const() 一起使用， iota只有在const进行累加效果。
	//var a int = iota 
}
```

#### 4. 函数

```go
package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

//返回多个返回值，匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 666, 777
}

//返回多个返回值， 有形参名称的
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("---- foo3 ----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//r1 r2 属于foo3的形参，  初始化默认的值是0
	//r1 r2 作用域空间 是foo3 整个函数体的{}空间
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("---- foo4 ----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)


	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	return
}

func main() {
	c := foo1("abc", 555)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("haha", 999)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)

	ret1, ret2 = foo3("foo3", 333)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)

	ret1, ret2 = foo4("foo4", 444)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)
}
```

#### 5. 指针

```go
package main  // 和C语言类似

import "fmt"

/*
func swap(a int ,b int) {
	var temp int
	temp = a
	a = b
	b = temp
}
*/

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa //temp = main::a
	*pa = *pb  // main::a = main::b
	*pb = temp // main::b = temp
}


func main() {
	var a int = 10
	var b int = 20

	swap(&a, &b)

	fmt.Println("a = ", a, " b = ", b)

	var p *int
	p = &a

	fmt.Println(&a)
	fmt.Println(p)

	var pp **int //二级指针
	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)
}
```



#### 6. defer

```go
package main

import "fmt"

func main() {
	//写入defer关键字
	defer fmt.Println("main end1")
	defer fmt.Println("main end2") // 栈的形式

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2") // return执行早于defer
}
控制台结果：
main::hello go 1
main::hello go 2
main end2
main end1

```



#### 7. 数组 Array

```go
package main

import "fmt"

func printArray(myArray [4]int) {
	//值拷贝

	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}
	myArray[0] = 111
}

func main() {
	//固定长度的数组
	var myArray1 [10]int

	myArray2 := [10]int{1,2,3,4}
	myArray3 := [4]int{11,22,33,44}

	//for i := 0; i < 10; i++ {
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])	
	}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	//查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

	printArray(myArray3)
	fmt.Println(" ------ ")
	for index, value := range myArray3 {
		fmt.Println("index = ", index, ", value = ", value)
	}
}
```

#### 8. 切片 slice

```go
package main

import "fmt"

func printArray(myArray []int) {
	//引用传递
	// _ 表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 100
}

func main() {
	myArray := []int{1,2,3,4} // 动态数组，切片 slice

	fmt.Printf("myArray type is %T\n", myArray)

	printArray(myArray)

	fmt.Println(" ==== ")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}
----------------------------- 切片2 ------------------------------------
package main

import "fmt"

func printArray(myArray []int) {
	//引用传递
	// _ 表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 100
}

func main() {
	myArray := []int{1,2,3,4} // 动态数组，切片 slice

	fmt.Printf("myArray type is %T\n", myArray)

	printArray(myArray)

	fmt.Println(" ==== ")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}
----------------------------- 切片3 ------------------------------------
package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//向numbers切片追加一个元素1, numbers len = 4， [0,0,0,1], cap = 5
	numbers = append(numbers, 1)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//向numbers切片追加一个元素2, numbers len = 5， [0,0,0,1,2], cap = 5
	numbers = append(numbers, 2)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//向一个容量cap已经满的slice 追加元素，
	numbers = append(numbers, 3)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("-=-------")
	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
}
----------------------------- 切片3 ------------------------------------
package main

import "fmt"

func main() {
	s := []int{1, 2, 3} //len = 3, cap = 3, [1,2,3]

	//[0, 2)
	s1 := s[0:2] // [1, 2]

	fmt.Println(s1)

	s1[0] = 100

	fmt.Println(s)
	fmt.Println(s1)

	//copy 可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3) //s2 = [0,0,0]

	//将s中的值 依次拷贝到s2中
	copy(s2, s)
	fmt.Println(s2)

}
```

#### 9. map

```go
package main

import "fmt"

func main() {
	//===> 第一种声明方式

	//声明myMap1是一种map类型 key是string， value是string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")
	}

	//在使用map前， 需要先用make给map分配数据空间
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"

	fmt.Println(myMap1)

	//===> 第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"

	fmt.Println(myMap2)

	//===> 第三种声明方式
	myMap3 := map[string]string{
		"one":   "php",
		"two":   "c++",
		"three": "python",
	}
	fmt.Println(myMap3)
}
-------------------- map2 ----------------------------
package main

import "fmt"

func printMap(cityMap map[string]string) {
	//cityMap 是一个引用传递
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	//遍历
	printMap(cityMap)

	//删除
	delete(cityMap, "China")

	//修改
	cityMap["USA"] = "DC"
	ChangeValue(cityMap)

	fmt.Println("-------")

	//遍历
	printMap(cityMap)
}


```

#### 10. 结构体 struct

```go
package main

import "fmt"

//声明一种行的数据类型 myint， 是int的一个别名
type myint int

//定义一个结构体
type Book struct {
   title string
   auth  string
}

func changeBook(book Book) {
   //传递一个book的副本
   book.auth = "666"
}

func changeBook2(book *Book) {
   //指针传递
   book.auth = "777"
}

func main() {
   /*
      var a myint = 10
      fmt.Println("a = ", a)
      fmt.Printf("type of a = %T\n", a)
   */

   var book1 Book
   book1.title = "Golang"
   book1.auth = "zhang3"

   fmt.Printf("%v\n", book1)

   changeBook(book1)

   fmt.Printf("%v\n", book1)

   changeBook2(&book1)

   fmt.Printf("%v\n", book1)
}
```

#### 11. 类class

```go
package main

import "fmt"

//如果类名首字母大写，表示其他包也能够访问
type Hero struct {
	//如果说类的属性首字母大写, 表示该属性是对外能够访问的，否则的话只能够类的内部访问
	Name  string
	Ad    int
	level int
}

/*
func (this Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this Hero) GetName() string {
	return this.Name
}

func (this Hero) SetName(newName string) {
	//this 是调用该方法的对象的一个副本（拷贝）
	this.Name = newName
}
*/
func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.level)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	//this 是调用该方法的对象的一个副本（拷贝）
	this.Name = newName
}

func main() {
	//创建一个对象
	hero := Hero{Name: "zhang3", Ad: 100}

	hero.Show()

	hero.SetName("li4")

	hero.Show()
}
---------------------------------- 继承 -------------------------------------------
package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

//=================

type SuperMan struct {
	Human //SuperMan类继承了Human类的方法

	level int
}

//重定义父类的方法Eat()
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

//子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}

func main() {
	h := Human{"zhang3", "female"}

	h.Eat()
	h.Walk()

	//定义一个子类对象
	//s := SuperMan{Human{"li4", "female"}, 88}
	var s SuperMan
	s.name = "li4"
	s.sex = "male"
	s.level = 88

	s.Walk() //父类的方法
	s.Eat()  //子类的方法
	s.Fly()  //子类的方法

	s.Print()
}

```

