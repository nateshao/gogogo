---
title: day1-Go语言上手-基础语言
date: 2022-05-07 23:23:33
tags: 
- Go学习路线
- 字节跳动青训营
---

[TOC]

![img](https://cdn.jsdelivr.net/gh/nateshao/images/20220510105108.jpeg)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220514094925.png)

## day01-Go语言上手-基础语言 | 青训营笔记

这是我参与「第三届青训营-后端场」笔记创作活动的的第1篇笔记

> **《字节跳动青训营》**是面向在校大学生开放免费社区的，有前端、后端专场。基础班和进阶班需要笔试（计算机基础知识+2道编程）过了即可。

|                         基础班课程表                         |
| :----------------------------------------------------------: |
| ![img](https://cdn.jsdelivr.net/gh/nateshao/images/20220507153823.png) |
|                         进阶班课程表                         |
| ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507155043.png) |

> 今天是基础班第一天上课笔记

# 01简介

## 1.1什么是Go语言

**Go语言有什么特点呢？**

1. 高性能、高并发
2. 语法简单、学习曲线平缓
3. 丰富的标准库
4. 完善的工具链
5. 静态链接
6. 快速编译
7. 跨平台
8. 垃圾回收

**举个简单例子**

简单两句代码就可以启动一个web服务

```go
package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe("localhost:8080", nil)
}
```

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507150112.png)

---

## 1.2哪些公司用Go语言 

> Go语言有哪些公司正在使用，然后主要应用在哪些场景？

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507163013.png)

1. 字节跳动已经全面拥抱了go语言，公司内部有上万个微服务使用golang来编写，不久前也开源了GO RPC框架KiteX。

   根据拉勾的招聘数据，**腾讯、百度、美团、滴滴、深信服、平安、OPPO、知乎、去哪儿、360、 金山、微博、哔哩哔哩、七牛、**

   **PingCAP** 等公司也在大星使用Go语言。国外**Google、Facebook**等公司也在大量使用Go语言。

2. 从业务维度看过语言已经在云计算、微服务、 大数据、区块链、物联网等领域蓬勃发展。然后在云计算、微服务等领域已经有非常高的市场占有率`Docker`、`Kubernetes`、 `Istio`、 `etcd`、 `prometheus` 几乎所有的云原生组件全是用Go实现的。

## 1.3字节跳动为什么全面拥抱GO语言

1. 最初使用的Python, 由于性能问题换成了Go
2. C++ 不太适合在线Web业务
3. 早期团队非Java背景
4. 性能比较好
5. 部署简单、学习成本低
6. 内部RPC和HTTP框架的推广

我们知道字节跳动已经全面拥抱了go语言，最开始公司的后端业务主要是web后端，早期团队非Java背景，C++不太适合在线Web业务，所以最开始的服务都是python的，大概从2014年开始，随着业务体量的增长，python 遇到一些性能问题。

一些团队初步尝试使用了Go,发现**入门很简单，开发效率高，性能也比较好**。go语言的开发和部署非常简单，顺带解决了之前python带来

的很头疼的依赖库版本问题。一些业务尝到甜头之后， 后面开始公司级大力推广，诞生 了公司内部的基于golang的rpc和http框架。

随着框架的推广，越来越多的python服务 使用golang重写，至今为止， golang已经成为内部使用率最高的编程语言。

# 02入门

> 这一章主要介绍：**开发环境，基础语法和标准库。**

## 2.1开发环境-安装Golang

|                        https://go.dev                        |
| :----------------------------------------------------------: |
| ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507163032.png) |
|                   https://studygolang.com                    |
| ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507165640.png) |
|                     https://goproxy.cn/                      |
| ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507165655.png) |

## 2.1开发环境-配置集成开发环境

|                      GoLang    开发环境                      |
| :----------------------------------------------------------: |
| ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507180156.png) |
|            **Visual Studio Code **  **开发环境**             |
| ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507180333.png) |
|        https://hi-hi.cn/gitpod  **基于云的开发环境**         |
| ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507180728.png) |

## 2.2基础语法-Hello World

如何配置Go开发环境。这里不仔细介绍了，自行百度即可。接下来我们来通过一些小例子教大家快速学习一 下go源代码的一 些基础语法。 我们先来看一下go语言的里面的 helloword，helloworld 代码大概长这样子

```go
package main // package main代表这个文件属于main包的一部分，main 包也就是程序的入口包。

import (
	"fmt" // 标准库里面的FMT包。这个包主要是用来往屏幕输入输出字符串、格式化字符串。
)

func main() {
	fmt.Println("hello world") // main 函数的话里面调用了fmt.Println 输出helloword
}
```

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507181526.png)

## 2.2基础语法-变量

>Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。声明变量的一般形式是使用 var 关键字：

**第一种，指定变量类型，如果没有初始化，则变量默认为零值**(变量没有做初始化时系统默认设置的值)。

```go
package main
import "fmt"
func main() {
    // 声明一个变量并初始化
    var a = "RUNOOB"
    fmt.Println(a)

    // 没有初始化就为零值
    var b int
    fmt.Println(b)

    // bool 零值为 false
    var c bool
    fmt.Println(c)
}
/** 输出
RUNOOB
0
false
/
```

**第二种，根据值自行判定变量类型。**

```go
package main
import "fmt"
func main() {
    var d = true
    fmt.Println(d)
}
// 输出：true
```

**第三种，如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误，格式：**

```go
package main
import "fmt"
func main() {
    f := "千羽" // var f string = "Runoob"
    fmt.Println(f)
}
// 输出：千羽
```

---

```go
package main

import (
	"fmt"
	"math"
)

func main() {

	var a = "initial"

	var b, c int = 1, 2

	var d = true

	var e float64

	f := float32(e)

	g := a + "foo"
	fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
	fmt.Println(g)                // initialapple

	const s string = "constant"
	const h = 500000000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}
// 输出
// initial 1 2 true 0 0
// initialfoo                                                      
// constant 500000000 6e+11 -0.28470407323754404 0.7591753930288755
```

## 2.3基础语法-if else

Go 语言提供了以下几种条件判断语句：

| 语句           | 描述                                                         |
| :------------- | :----------------------------------------------------------- |
| if 语句        | **if 语句** 由一个布尔表达式后紧跟一个或多个语句组成。       |
| if...else 语句 | **if 语句** 后可以使用可选的 **else 语句**, else 语句中的表达式在布尔表达式为 false 时执行。 |
| if 嵌套语句    | 你可以在 **if** 或 **else if** 语句中嵌入一个或多个 **if** 或 **else if** 语句。 |
| switch 语句    | **switch** 语句用于基于不同条件执行不同动作。                |
| select 语句    | **select** 语句类似于 **switch** 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。 |

**注意：Go 没有三目运算符，所以不支持 ?: 形式的条件判断。**

go语言里面的if else写法和C或者C++类似。不同点是if后面没有括号。如果你写括号的话，那么在保存的时候你的编辑器会自动把你去掉。第二个不同点是Golang里面的if，它必须后面接大括号，就是你不能像C或者C++ 一样，直接把if里面的语句同一行。

```go
package main

import "fmt"

func main() {
	// if 后面没有括号
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
```

## 2.4基础语法-循环

go语言里面的循环，**在go里面没有while循环、do while循环，只有唯一的一种for循环**。最简单的for循环就是在for后面什么都不写，代表一个死循环。循环途中你可以用break跳出。在循环里面，你可以用break或者continue来跳出或者继续循环，

Go 语言提供了以下几种类型循环处理语句：

| 循环类型 | 描述                                 |
| :------- | :----------------------------------- |
| for 循环 | 重复执行语句块                       |
| 循环嵌套 | 在 for 循环中嵌套一个或多个 for 循环 |

**循环控制语句**

> 循环控制语句可以控制循环体内语句的执行过程。

GO 语言支持以下几种循环控制语句：

| 控制语句      | 描述                                             |
| :------------ | :----------------------------------------------- |
| break 语句    | 经常用于中断当前 for 循环或跳出 switch 语句      |
| continue 语句 | 跳过当前循环的剩余语句，然后继续进行下一轮循环。 |
| goto 语句     | 将控制转移到被标记的语句。                       |

```go
package main

import "fmt"

func main() {

	i := 1
	for {  // for里面什么都不写代表死循环
		fmt.Println("loop")
		break
	}
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
}
```

## 2.5基础语法-switch

> **switch** 语句用于基于不同条件执行不同动作。

go语言里面的switch分支结构。看起来也C或者C++比较类似。同样的在switch后面的那个变量名，并不是要括号。

这里有个很大的一点不同的是，在c++里面，switch case如果不显示加break的话会然后会继续往下跑完所有的case,在go语言里面的话是不需要加break的。

相比C或者C++，go语言里面的switch功能更强大。可以使用任意的变量类型，甚至可以用来取代任意的if else语句。你可以在switch后面不加任何的变量，然后在case里面写条件分支。这样代码相比你用多个if else代码逻辑会更为清晰。

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	a := 2
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5:
		fmt.Println("four or five")
	default:
		fmt.Println("other")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
```

```go
// 输出：
two
It's after noon
```

## 2.6基础语法-数组

> Go 语言提供了数组类型的数据结构。数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整型、字符串或者自定义类型。

数组就是一个具有编号且长度固定的元素序列。比如这里的话是一个可以存放5个int元素
对于一个数组，可以很方便地取特定索引的值或者往特定索引取存储值，然后也能够直接去打印一个数组。不过，在真实业务代码里面，**我们很少直接使用数组，因为它长度是固定的，我们用的更多的是切片。**

```go
package main

import "fmt"

func main() {

	var a [5]int
	a[4] = 100
	fmt.Println("get:", a[2])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
```

---

```go
// 输出
get: 0
len: 5                
[1 2 3 4 5]           
2d:  [[0 1 2] [1 2 3]]
```

## 2.7基础语法-切片
> Go 语言切片是对数组的抽象。Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

我们可以用make来创建一个切片， 可以像数组一样去取值，使用append来追加元素。

注意append的用法的话，你必须把append的结果赋值为原数组。
因为slice的原理实际上是它有一个它存储了一个长度和一个容量，加一个指向一个数组的指针，在你执行append操作的时候，如果容量不够的话，会扩容并且返回新的slice。

slice 初始化的时候也可以指定长度。slice 拥有像python 一样的切片操作，比如这个代表取出第二个到第五个位置的元素， 不包括第五个元素。不过不同于python，这里不支持负数索引。

```go
package main

import "fmt"

func main() {

	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s[2])   // c
	fmt.Println("len:", len(s)) // 3

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s) // [a b c d e f]

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c) // [a b c d e f]

	fmt.Println(s[2:5]) // [c d e]
	fmt.Println(s[:5])  // [a b c d e]
	fmt.Println(s[2:])  // [c d e f]

	good := []string{"g", "o", "o", "d"}
	fmt.Println(good) // [g o o d]
}
```

## 2.8基础语法-map

> Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
>
> Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

map，在其他编程语言里面， 它可能可以叫做哈希或者字典。map 是实际使用过程中最频繁用到的数据结构。我们可以用make来创建一个空 map,这里会需要两个类型，第一个是那个key的类型，这里是string另一个是value 的类型，这里是int。可以从里面去存储或者取出键值对。可以用delete从里面删除键值对。

golang的map是完全无序的，遍历的时候不会按照字母顺序，也不会按照插入顺序输出，而是随机顺序。

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)           // map[one:1 two:2]
	fmt.Println(len(m))      // 2
	fmt.Println(m["one"])    // 1
	fmt.Println(m["unknow"]) // 0

	r, ok := m["unknow"]
	fmt.Println(r, ok) // 0 false

	delete(m, "one")

	m2 := map[string]int{"one": 1, "two": 2}
	var m3 = map[string]int{"one": 1, "two": 2}
	fmt.Println(m2, m3)
}
```

## 2.9基础语法-range

> Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

用range来快速遍历，这样代码能够更加简洁。 range遍历的时候，对于数组会返回两个值，第一个是索引，第二个是对应位置的值。如果我们不需要索引的话，我们可以用下划线来忽略。

```go
package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num) // index: 0 num: 2
		}
	}
	fmt.Println(sum) // 9

	m := map[string]string{"a": "A", "b": "B"}
	for k, v := range m {
		fmt.Println(k, v) // b 8; a A
	}
	for k := range m {
		fmt.Println("key", k) // key a; key b
	}
}
```

## 2.10基础语法-函数

Go 语言最少有个 main() 函数(java叫做方法)。你可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务。函数声明告诉了编译器函数的名称，返回类型，和参数。

Go 语言标准库提供了多种可动用的内置的函数。例如，len() 函数可以接受不同类型参数并返回该类型的长度。如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。

```go
package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

func main() {
	res := add(1, 2)
	fmt.Println(res) // 3

	v, ok := exists(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok) // A True
}
```

这个是Golang里面一个简单的实现两个变量相加的函数。Golang 和其他很多语言不一样的是，变量类型是后置的。
Golang里面的函数原生支持返回多个值。在实际的业务逻辑代码里面几乎所有的函数都返回两个值，**第一个是真正的返回结果，第二个值是一个错误信息。**

## 2.11基础语法-指针

我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。

Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。

```go
package main

import "fmt"

func add2(n int) {
	n += 2
}

func add2ptr(n *int) {
	*n += 2
}

func main() {
	n := 5
	add2(n)
	fmt.Println(n) // 5
	add2ptr(&n)
	fmt.Println(n) // 7
}
```

go里面也支持指针。**指针的一个主要用途就是对于传入参数进行修改。**
我们来看这个函数。这个函数试图把一个变量+2。 但是单纯像上面这种写法其实是无效的。因为传入函数的参数实际上是一个拷贝， 那也说这个+2,是对那个拷贝进行了+2，并不起作用。 如果我们需要起作用的话，那么我们需要把那个类型写成指针类型，那么为了类型匹配，调用的时候会加一个&符号。

## 2.12基础语法-结构体
> Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
>
> 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。

```go
package main

import "fmt"

type user struct {
	name     string
	password string
}

func main() {
	a := user{name: "wang", password: "1024"}
	b := user{"wang", "1024"}
	c := user{name: "wang"}
	c.password = "1024"
	var d user
	d.name = "wang"
	d.password = "1024"

	fmt.Println(a, b, c, d)                 // {wang 1024} {wang 1024} {wang 1024} {wang 1024}
	fmt.Println(checkPassword(a, "haha"))   // false
	fmt.Println(checkPassword2(&a, "haha")) // false
}

func checkPassword(u user, password string) bool {
	return u.password == password
}

func checkPassword2(u *user, password string) bool {
	return u.password == password
}
```

结构体的话是带类型的字段的集合。
比如这里user结构包含了两个字段，name和password。我们可以用结构体的名称去初始化一个结构体变量，构造的时候需要传入每个字段的初始值。也可以用这种键值对的方式去指定初始值，这样可以只对一部分字段进行初始化。同样的结构体我们也能支持指针，这样能够实现对于结构体的修改，也可以在某些情况下避免一些大结构体的拷贝开销。

## 2.13基础语法-结构体方法

在Golang里面可以为结构体去定义一些方法。 会有一点类似其他语言里面的类成员函数。 如这里，我们把上面一个例子的checkPassword的实现，从一个普通函数，改成了结构体方法。这样用户可以像a.checkPassword("xx”) 这样去调用。具体的代码修改，就是把第一个参数， 加上括号，写到函数名称前面。

在实现结构体的方法的时候也有两种写法，一种是带指针， 一种是不带指针。这个它们的区别的话是说如果你带指针的话，那你那么你就可以对这个结构体去做修改。如果你不带指针的话，那你实际上操作的是一个拷贝， 你就无法对结构体进行修改。

```go
package main

import "fmt"

type user struct {
	name     string
	password string
}

func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

func main() {
	a := user{name: "wang", password: "1024"}
	a.resetPassword("2048")
	fmt.Println(a.checkPassword("2048")) // true
}
```

## 2.14基础语法-错误处理

Go 语言通过内置的错误接口提供了非常简单的错误处理机制。

error类型是一个接口类型，这是它的定义：

```go
type error interface {
    Error() string
}
```

我们可以在编码中通过实现 error 接口类型来生成错误信息。

函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息：

```go
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // 实现
}
```

在下面的例子中，我们在调用Sqrt的时候传递的一个负数，然后就得到了non-nil的error对象，将此对象与nil比较，结果为true，所以fmt.Println(fmt包在处理error时会调用Error方法)被调用，以输出错误，请看下面调用的示例代码：

```go
result, err:= Sqrt(-1)

if err != nil {
   fmt.Println(err)
}
```

---

错误处理在go语言里面符合语言习惯的做法就是使用一个单独的返回值来传递错误信息。

不同于Java自家使用的异常。go语言的处理方式，能够很清晰地知道哪个函数返回了错误，并且能用简单的if else来处理错误。
在函数里面，我们可以在那个函数的返回值类型里面，后面加一个error,就代表这个函数可能会返回错误。

那么在函数实现的时候，return 需要同时return 两个值，要么就是如果出现错误的话，那么可以return nil 和一个error。如果没有的话，那么返回原本的结果和nil

```go
package main

import (
	"errors"
	"fmt"
)

type user struct {
	name     string
	password string
}

func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

func main() {
	u, err := findUser([]user{{"wang", "1024"}}, "wang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name) // wang

	if u, err := findUser([]user{{"wang", "1024"}}, "li"); err != nil {
		fmt.Println(err) // not found
		return
	} else {
		fmt.Println(u.name)
	}
}
```

## 2.15基础语法-字符串操作

下面我们来看go语言里面的字符串操作。在标准库string包里面有很多常用的字符串工具函数，比如contains判断个字符串里面是否有包含另个字符串 ，count 字符审计数，index 查找某个字符审的位置。join 连接多个字符串

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"
	fmt.Println(strings.Contains(a, "ll"))                // true
	fmt.Println(strings.Count(a, "l"))                    // 2   统计
	fmt.Println(strings.HasPrefix(a, "he"))               // true
	fmt.Println(strings.HasSuffix(a, "llo"))              // true
	fmt.Println(strings.Index(a, "ll"))                   // 2
	fmt.Println(strings.Join([]string{"he", "llo"}, "-")) // he-llo  拼接
	fmt.Println(strings.Repeat(a, 2))                     // hellohello  
	fmt.Println(strings.Replace(a, "e", "E", -1))         // hEllo  
	fmt.Println(strings.Split("a-b-c", "-"))              // [a b c]  分割
	fmt.Println(strings.ToLower(a))                       // hello  转小写
	fmt.Println(strings.ToUpper(a))                       // HELLO  转大写
	fmt.Println(len(a))                                   // 5      长度
	b := "你好"
	fmt.Println(len(b)) // 6
}
```

## 2.16基础语法-字符串格式化

标准库的FMT包里面有很多的字符串格式相关的方法，比如prinf这个类似于C语言里面的printf 函数。不同的是，在go语言里面的话, 你可以很轻松地用%v来打印任意类型的变量，而不需要区分数字符串。也可以用%+v打印详细结果，%#v则更详细。

```go
package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	s := "hello"
	n := 123
	p := point{1, 2}
	fmt.Println(s, n) // hello 123
	fmt.Println(p)    // {1 2}

	fmt.Printf("s=%v\n", s)  // s=hello
	fmt.Printf("n=%v\n", n)  // n=123
	fmt.Printf("p=%v\n", p)  // p={1 2}
	fmt.Printf("p=%+v\n", p) // p={x:1 y:2}
	fmt.Printf("p=%#v\n", p) // p=main.point{x:1, y:2}

	f := 3.141592653
	fmt.Println(f)          // 3.141592653
	fmt.Printf("%.2f\n", f) // 3.14
}

```

## 2.17基础语法-JSON处理

下面我们来看一下JSON操作，go语言里面的JSON操作非常简单，对于一个已有的结构体，我们可以什么都不做，只要保证每个字段的第一个字母是大写，也就是是公开字段。那么这个结构体就能用**JSON.marshaler去序列化**，变成一个JSON的字符串。

序列化之后的字符串也能够用**JSON.Unmarshal去反序列化**到一个空的变量里面。

这样默认序列化出来的字符串的话，它的风格是大写字母开头，而不是下划线。我们可以在后面用json tag等语法来去修改输出JSON结果里面的字段名。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}

func main() {
	a := userInfo{Name: "wang", Age: 18, Hobby: []string{"Golang", "TypeScript"}}
	buf, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf)         // [123 34 78 97...]
	fmt.Println(string(buf)) // {"Name":"wang","age":18,"Hobby":["Golang","TypeScript"]}

	buf, err = json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	var b userInfo
	err = json.Unmarshal(buf, &b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b) // main.userInfo{Name:"wang", Age:18, Hobby:[]string{"Golang", "TypeScript"}}
}
```

## 2.18基础语法-时间处理

在go语言里面最常用的就是time.Now()来获取当前时间，然后你也可以用time.Date()去构造一个带时区的时间， 构造完的时间。上面有很多方法来获取这个时间点的年月日 小时分钟秒，然后也能用点sub去对两个时间进行减在和某些系统交互的时候，我们经常会用到时间戳。可以用.UNIX来获取时间戳time.format()、 time.Parse()

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now) // 2022-03-27 18:04:59.433297 +0800 CST m=+0.000087933
	t := time.Date(2022, 3, 27, 1, 25, 36, 0, time.UTC)
	t2 := time.Date(2022, 3, 27, 2, 30, 36, 0, time.UTC)
	fmt.Println(t)                                                  // 2022-03-27 01:25:36 +0000 UTC
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()) // 2022 March 27 1 25
	fmt.Println(t.Format("2006-01-02 15:04:05"))                    // 2022-03-27 01:25:36
	diff := t2.Sub(t)
	fmt.Println(diff)                           // 1h5m0s
	fmt.Println(diff.Minutes(), diff.Seconds()) // 65 3900
	t3, err := time.Parse("2006-01-02 15:04:05", "2022-03-27 01:25:36")
	if err != nil {
		panic(err)
	}
	fmt.Println(t3 == t)    // true
	fmt.Println(now.Unix()) // 1648738080
}
```

## 2.19基础语法-数字解析

字符串和数字之间的转换。在go语言当中，关于字符串和数字类型之间的转换都在strconv这个包下，这个包是string convert这两个单词的缩写。
可以用strconv.ParseInt()或者strconv.ParseFloat()来解析一个字符串。可以用strconv.Atoi("AAA")把一个十进制字符串转成数字。可以用strconv.Itoa()把数字转成字符串。如果输入不合法，那么这些函数都会返回error

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f) // 1.234

	n, _ := strconv.ParseInt("111", 10, 64)
	fmt.Println(n) // 111

	n, _ = strconv.ParseInt("0x1000", 0, 64)
	fmt.Println(n) // 4096

	n2, _ := strconv.Atoi("123")
	fmt.Println(n2) // 123
	n2, err := strconv.Atoi("AAA")
	fmt.Println(n2, err) // 0 strconv.Atoi: parsing "AAA": invalid syntax
	n3 := strconv.Itoa(123)
	fmt.Println(n3) // 123

}
```

## 2.20基础语法-进程信息
在go里面，我们能够用os.Args来得到程序执行的时候的指定的命令行参数。比如我们编译的一一个二进制文件，command。 后面接 abcd来启动，输出就是os.Args会是一个长度为 5的slice ,第一个成员代表二 进制自身的名字。我们可以用os.Getenv("PATH")来读取环境变量。

```go
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// go run example/20-env/main.go a b c d
	fmt.Println(os.Args) // [/var/folders/8p/n34xxfnx38dg8bv_x8l62t_m0000gn/T/go-build3406981276/b001/exe/main a b c d]
	fmt.Println("-----------")
	fmt.Println(os.Getenv("PATH")) // /usr/local/go/bin...
	fmt.Println("-----------")
	fmt.Println(os.Setenv("AA", "BB"))
	fmt.Println("-----------")

	buf, err := exec.Command("grep", "127.0.0.1", "/etc/hosts").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf)) // 127.0.0.1       localhost
}
```

# 03 实战

> 这里主要介绍**猜谜游戏，在线词典和SOCKS5代理。**

## 3.1猜谜游戏介绍

这里用Golang来构建一个猜数字游戏。 在这个游戏里面，程序首先会生成一个介于 1到100之间的随机整数，然后提示玩家进行猜测。玩家每次输入一个数字，程序会告诉玩家这个猜测的值是高于还是低于那个秘密的值。如果猜对了，就告诉玩家胜利并且退出程序。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507203547.png)

## 3.1.1猜谜游戏-生成随机数
```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxNum := 100
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
}
```

当程序运行的时候会生成一个0到100之间的随机数字。 我们先来生成这个随机数。为了生成随机数，我们需要用到math/rand包。我们的第一个版本的代码是这样子的，我们先导入fmt包和math/rand包， 定义一个变量，maxNum是100。下面用rand.Intn来生成一个随机数， 再打印出这个随机数。

## 3.1.2猜谜游戏-生成随机数效果
![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507204235.png)

我们发现每次都会打印相同的数字到屏幕上。这个不是我们想要的，为什么呢?

## 3.1.2猜谜游戏-生成随机数V2
我们用time.Now().UnixNano()来初始化随机种子。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507204526.png)

## 3.1.3猜谜游戏-读取用户输入

实现用户输入输出，并成数字。
我们可以用它的ReadString 方法来读取一行。如果失败了的话，我们会打印错误并能退出。ReadString 返回的结果包含结尾的换行符，我们把它去掉，再转换成数字。如果转换失败，我们同样打印错误，退出。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507205304.png)

## 1.4猜谜游戏-实现判断逻辑

现在我们有了一个秘密的值，然后也从用户的输入里面读到了一个值， 我们来比较这两个值的大小。如果是用户输入的值比那个秘密的值要大的话，就告诉用户你猜的值太大了，请再试一次。如果是小了也同理，如果是相等的话，那么我们就告诉用户你赢了。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507205651.png)

## 3.1.5猜谜游戏-实现游戏循环

此时我们的程序大致可以正常工作了，但是玩家只能输入一次猜测，无论猜测是否正确，程序都会突退出。为了改变这种行为，让游戏可以正常玩下去，我们需要加一个循环。我们把刚刚的代码挪到一个for循环里面，再**把return改成continue以便于在出错的时候能够继续循环**。在用户输入正确的时候break，这样才能够在用户胜利的时候退出游戏。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507210000.png)

就这样，我们已经成功地在Golang里面构建了一个猜谜游戏。在这个过程中，我们复习了之前的很多概念，比如**变量循环、函数控制流和错误处理。**

## 3.2在线词典介绍

实现一个命令行排版的词典

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507210357.png)

用户可以在命令行里面查询一个单词。我们能通过调用第三方的API查询到单词的翻译并打印出来。
这个例子里面，我们会学习如何用go语言来来发送HTTP请求、解析json 过来，还会学习如何使用代码生成来提高开发效率。

## 3.2.1在线词典-抓包

我们先来看一下我们要用到的API，以彩云科技提供的在线翻译为例。先请打开彩云翻译的网页，然后右键F12检查打开浏览器的开发者工具。

|           彩云小译：https://fanyi.caiyunapp.com/#/           |
| :----------------------------------------------------------: |
| ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507211033.png) |
| ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507211307.png) |
| ![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507211259.png) |

此时我们点一下翻译按钮，浏览器会发送系列请求， 我们能很轻松地找到那个用来查询单词的请求。

这是一个HTTP的post的请求，请求的header的相当的复杂，有十来个。然后请求头是一个 json里面有两个字段，一个是代表你要你是从什么语言转化成什么语言，source 就是你要查询的单词。API 的返回结果里面会有Wiki和dictionary两个字段。我们需要用的结果主要在dictionary Explanations字段里面。其他有些字段里面还包括音标等信息。

## 3.2.2在线词典-代码生成

我们需要在Golang里面去发送这个请求。因为这个请求比较复杂，用代码构造很麻烦，实际上我们有一种非常简单的方式来生成代码，我们可以右键浏览器里面的copy as cur。copy完成之 后大家可以在终端粘贴一下curl命令，应该可以成功

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507211602.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507211658.png)

## 3.2.2在线词典-代码生成

然后我们打开一个网站https://curlconverter.com/#go，粘贴curl请求，在右边的语言里面选Golang就能够看到一串很长的代码，我们直接把它copy到我们的编辑器里面。有几个header比较复杂， 生成代码有转义导致的编译错误，删掉这几行即可。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507211926.png)

## 3.2.2在线词典-生成代码解读

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	var data = strings.NewReader(`{"trans_type":"en2zh","source":"good"}`)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("app-name", "xy")
	req.Header.Set("os-type", "web")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
```

我们来看一下这生成的代码：

1. 首先第12行我们创建了一个 HTTP client,创建的时候可以指定很多参数，包括比如请求的超时是否使用cookie等。接下来是构造一个 HTTP请求，这是一个 post请求。
2. 然后会用到HTTP .NewRequest，第一个参数是http方法POST，第二个参数是URL，最后一个参数是 body，body因为可能很大，为了支持流式发送，是一个只读流。 我们用了strings.NewReader来把字符串转换成一个流。这样我们就成功构造了一个HTTP request,接下来我们需要对这个HTTP request来设置一堆header.
3. 接下来我们把我们调用client.do request，就能得到response如果请求失败的话，那么这个error会返回非nil,会打印错误并且退出进程。response 有它的HTTP状态码，response header和body。body同样是一个流， 在golang里面， 为了避免资源泄露，你需要加一个**defer来手动关闭这个流**，这个defer会在这个函数运行结束之后去执行。接下来我们是用ioutil.Readll来读取这个流，能得到整个body。我们再用print打印出来。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507212624.png)

我们来运行生成的代码，能看到我们已经能够成功地发出请求，把返回的JSON打印出来。但是现在那个输入是固定的，我们是要从一个变量来输入，我们需要用到JSON列化

![image-20220507212929130](https://cdn.jsdelivr.net/gh/nateshao/images/20220507212929.png)

## 3.2.3在线词典-生成request body

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507213249.png)

## 3.2.4在线词典-解析response body

接下来我们要做的是把这个response body来解析出来。
在js/Python这些脚本语言里面，body 是一个字典或者map的结构，可以直接从里面取值。 但是golang是 个强类型语言，这种做法并不是最佳实践。
更常用的方式是和request的一样，写一个结构体，把返回的JSON反序列化到结构体里面。但是我们在浏览器里面可以看到这个API返回的结构非常复杂，如果要一定义结构体字段，非常繁琐并且容易出错。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507213449.png)

此时有一个小技巧的是，网上有对应的代码生成工具，我们可以打开这个网站https://oktools.net/json2go，把json字符串粘贴进去，这样我们就能够生成对应结构体。在某些时刻，我们如果不需要对这个返回结果，做很多精细的操作，我们可以选择转换嵌套，能让生成的代码更加紧凑。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507213626.png)

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507213531.png)



这样我们就得到了一个response结构体。接下来我们修改代码，我们先定一个 response结构体的对象，然后我们用JSON.unmarshal把body反序列化到这个结构体里面，再试图打印出来

![image-20220507214011259](https://cdn.jsdelivr.net/gh/nateshao/images/20220507214011.png)

现在我们再运行一下，这里打印的时候使用了%#v，这样可以让打印出来的结果比较容易读。我们现在离最终版本已经很近了，接下来我们需要修改代码为打印response里面的特定字段。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507214145.png)

## 3.2.5在线词典-打印结果
观察那个json可以看出我们需要的结果是在Dictionary.explanations.我们用for range循环来迭代它，然后直接打印结构，参照一些词典的显示方式， 我们可以在那个前面打印出这个单词和它的音标。这里有英式音标和美式音标。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507214508.png)

## 3.2.6在线词典-完善代码

现在我们的程序的输入还是写死的。我们把代码的主体改成个query函数，查询的单词作为参数传递进来。然后我们写一个简单的 main函数，这个main函数首先判断下命令和参 数的个数，如果它不是两个，那么我们就打印出错误信息，退出程序。否则就获取到用户输入的单词， 然后执行query函数。

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}

type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func query(word string) {
	client := &http.Client{}
	request := DictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Set("app-name", "xy")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("device-id", "")
	req.Header.Set("os-type", "web")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "_ym_uid=16456948721020430059; _ym_d=1645694872")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD example: simpleDict hello`)
		os.Exit(1)
	}
	word := os.Args[1]
	query(word)
}
```

这样子我们的命令行词典就算完成了，我们可以简单地试一下。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507210357.png)

## 3.3 SOCKS5代理介绍

我们来写一个socks5代理服务器，对于大家来说，一提到代理服务器， 第一想到的是翻墙。 不过很遗憾的是，socks5 协议它虽然是代理协议，但它并不能用来翻墙，它的协议都是明文传输。

这个协议历史比较久远，诞生于互联网早期。它的用途是，比如某 些企业的内网为了确保安全性，有很严格的防火墙策略，但是带来的副作用就是访问某些资源会很麻烦。

socks5相当于在防火墙开了个口子，让授权的用户可以通过单个端口去访问内部的所有资源。实际上很多翻墙软件，最终暴露的也是一个 socks5协议的端口。如果有同学开发过爬虫的话，就知道，在爬取过程中很容易会遇到IP访问频率超过限制。这个时候很多人就会去网上找一些代理IP池，这些代理IP池里面的很多代理的协议就是socks5。

> go-by-example\proxy\v4> go run .\main.go
>
> curl --socks5 127.0.0.1:1080 -v http://www.qq.com

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507215348.png)

## 3.3 SOCKS5代理-原理

接下来我们来了解一下 socks5协议的工作原理。正常浏览器访问一个网站，如果不经过代理服务器的话，就是先和对方的网站建立TCP连接，然后三次握手，握手完之后发起HTTP请求，然后服务返回HTTP响应。如果设置代理服务器之后，流程会变得复杂一些

1. 首先是浏览器和socks5代理建立TCP连接，代理再和真正的服务器建立TCP连接。这里可以分成四个阶段，握手阶段、认证阶段、请求阶段、relay 阶段。
2. 第一个握手阶段，浏览器会向socks5代理发送请求，包的内容包括一个协议的版本号 ，还有支持的认证的种类，socks5 服务器会选中一个认证方式，返回给浏览器。如果返回的是00的话就代表不需要认证，返回其他类型的话会开始认证流程，这里我们就不对认证流程进行概述了。
3. 第三个阶段是请求阶段，认证通过之后浏览器会socks5服务器发起请求。主要信息包括版本号，请求的类型，一般主要是connection请求，就代表代理服务器要和某个域名或者某个IP地址某个端口建立TCP连接。代理服务器收到响应之后，会真正和后端服务器建立连接，然后返回一个响应
4. 第四个阶段是relay阶段。此时浏览器会发送正常发送请求，然后代理服务器接收到请求之后，会直接把请求转换到真正的服务器上。然后如果真正的服务器以后返回响应的话，那么也会把请求转发到浏览器这边。然后实际上代理服务器并不关心流量的细节，可以是HTTP流量，也可以是其它 TCP流星。这个就是 socks5协议的工作原理，接下来我们就会试图去简单地实现它。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507215517.png)

## 3.3.1 SOCKS5代理-TCP echo server

第一步，我们先在go里面写一个简单的TCP echo server。为了方便测试，server 的工作逻辑很简单，你给他发送啥，他就回复啥，大概代码会长这样子:

1. 首先我们在main函数里面先用net.listen去监听一个端口，会返回一个server,然后在一个死循环里面，每次去accept一个请求，成功就会返回一个连接。接下来的话我们在一个process函数里面去处理这个连接。注意这前面会有个go关键字，这个代表启动一个goroutinue,可以暂时类比为其他语言里面的启动一 个子线程。只是这里的goroutinue的开销会比子线程要小很多，可以很轻松地处理上万的并发。
2. 接下来是这个process函数的实现。首先第一步的话会先加一个 defer connection.close()，defer 是Golang里面的一个语法，这一行的含义就是代表在这个函数退出的时候要把这个连接关掉，否则会有资源的泄露。
3. 接下来的话我们会用bufio.NewReader来创建一个带缓冲的只读流，这个在前面的猜谜游戏里面也有用到，带缓冲的流的作用是，可以减少底层系统调用的次数，比如这里为了方便是一个字节一 个字节的读取，但是底层可能合并成几次大的读取操作。并且带缓冲的流会有更多的一些工具数用来读取数据。我们可以简单地调用那个readbyte函数来读取单个字节。再把这一个字节写进去连接。

```go
package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process(client)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		_, err = conn.Write([]byte{b})
		if err != nil {
			break
		}
	}
}
```

---

我们来简单测试一下我们的第一 个TCP服务器，然后测试会需要用到nc命令（window需要安装），我们用nc 127.0.0.1 1080，输入Hello然后服务器就会给你返回Hello。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507222330.png)

## 3.3.2 SOCKS5代理- auth

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

const socks5Ver = 0x05
const cmdBind = 0x01
const atypIPV4 = 0x01
const atypeHOST = 0x03
const atypeIPV6 = 0x04

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process(client)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	err := auth(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err)
		return
	}
	log.Println("auth success")
}

func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+----------+----------+
	// |VER | NMETHODS | METHODS  |
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 |
	// +----+----------+----------+
	// VER: 协议版本，socks5为0x05
	// NMETHODS: 支持认证的方法数量
	// METHODS: 对应NMETHODS，NMETHODS的值为多少，METHODS就有多少个字节。RFC预定义了一些值的含义，内容如下:
	// X’00’ NO AUTHENTICATION REQUIRED
	// X’02’ USERNAME/PASSWORD

	ver, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read ver failed:%w", err)
	}
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}
	methodSize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read methodSize failed:%w", err)
	}
	method := make([]byte, methodSize)
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("read method failed:%w", err)
	}
	log.Println("ver", ver, "method", method)
	// +----+--------+
	// |VER | METHOD |
	// +----+--------+
	// | 1  |   1    |
	// +----+--------+
	_, err = conn.Write([]byte{socks5Ver, 0x00})
	if err != nil {
		return fmt.Errorf("write failed:%w", err)
	}
	return nil
}
```

就这样我们就已经完成了一个能够返回你输入信息的一个TCP server,接下来我们是要开始实现协议的第一步， 认证阶段，从这一部分开始会变得比较复杂。
我们实现一个空的auth函数，在process函数里面调用，再来编写auth函数的代码。我们回忆一下认证阶段的逻辑，

首先第一步的话， 浏览器会给代理服务器发送一个包， 然后这个包有三个字段，第一个字段，version 也就是协议版本号，固定是5第二个字段 methods,认证的方法数目第三个字段每个method的编码，0代表 不需要认证，2 代表用户名密码认证

我们先用read bytes来把版本号读出来，然后如果版本号不是socket 5的话直接返回报错，接下来我们再读取method size，也是一个字节。 然后我们需要我们去make一个相应长度的一个slice，用io.ReadFull把它去填充进去。写到这里，我们把获取到的版本号和认证方式打印一下。此时，代理服务器还需要返回一个response，返回包包括两个字段，一个是version一个是method,也就是我们选中的鉴传方式，我们当前只准备实现不需要鉴传的方式，也就是00。我们用curl命令测试一下当前版本的效果

## 3.3.3 SOCKS5代理-请求阶段

接下来我们开始做第三步，实现请求阶段，我们试图读取到携带URL或者IP地址+端口的包，然后把它打印出来。我们实现一个和 auth函数类似的connect 函数，同样在process里面去调用。再来实现connect函数的代码。

我们来回忆一下请求阶段的逻辑。浏览器会发送一个包， 包里面包含如下6个字段，version 版本号，还是5。command ，代表请求的类型，我们只支持connection请求，也就是让代理服务建立新的TCP连接。RSV 保留字段，不理会。atype 就是目标地址类型，可能是IPV 4 IPV 6或者域名下面是addr,这个地址的长度是根据atype的类型而不同的port 端口号，两个字节，我们需要逐个去读取这些字段。
面这四个字段总共四个字节，我们可以一次性把它读出来。我们定一个长度为4的buffer然后把它读满。读满之后，

然后第0个、 第1个、第3个、分别是version cmd和type,version需要判断是socket5，cmd 需要判断是1。下面的 type,可能是ipv4，ipv6， 或者是host。如果IPV 4的话，我们再次读满这个buffer，因为这个buffer长度刚好也是4个字节， 然后逐个字节打印成IP地址的格式保存到addr变量。

如果是个host的话，需要先读它的长度，再make 一个相应长度的buf填充它。再转换成字符串保存到 addr变量。IPV 6用得比较少，我们就暂时先不支持。 最后还有两个字节那个是port，我们读取它，然后按协议规定的大端字节序转换成数字。由于上面的buffer已经不会被其他变量使用了，我们可以直接复用之前的内存，建立一个临时的 slice，长度是2用于读取，这样的话最多会只读两个字节回来。接下来我们把这个地址和端口打印出来用于调试。收到浏览器的这个请求包之后，我们需要返回一个包，这个包有很多字段，但其实大部分都不会使用。

第一个是版本号还是socket5。第 二个， 就是返回的类型，这里是成功就返回0第三个是保留字段填0第四个atype地址类型填1第五个，第六个暂时用不到，都填成0。一 共4+4+2个字节，后面6个字节都是0填充。

```go
package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

const socks5Ver = 0x05
const cmdBind = 0x01
const atypIPV4 = 0x01
const atypeHOST = 0x03
const atypeIPV6 = 0x04

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process(client)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	err := auth(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err)
		return
	}
	err = connect(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err)
		return
	}
}

func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+----------+----------+
	// |VER | NMETHODS | METHODS  |
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 |
	// +----+----------+----------+
	// VER: 协议版本，socks5为0x05
	// NMETHODS: 支持认证的方法数量
	// METHODS: 对应NMETHODS，NMETHODS的值为多少，METHODS就有多少个字节。RFC预定义了一些值的含义，内容如下:
	// X’00’ NO AUTHENTICATION REQUIRED
	// X’02’ USERNAME/PASSWORD

	ver, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read ver failed:%w", err)
	}
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}
	methodSize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read methodSize failed:%w", err)
	}
	method := make([]byte, methodSize)
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("read method failed:%w", err)
	}

	// +----+--------+
	// |VER | METHOD |
	// +----+--------+
	// | 1  |   1    |
	// +----+--------+
	_, err = conn.Write([]byte{socks5Ver, 0x00})
	if err != nil {
		return fmt.Errorf("write failed:%w", err)
	}
	return nil
}

func connect(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+-----+-------+------+----------+----------+
	// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER 版本号，socks5的值为0x05
	// CMD 0x01表示CONNECT请求
	// RSV 保留字段，值为0x00
	// ATYP 目标地址类型，DST.ADDR的数据对应这个字段的类型。
	//   0x01表示IPv4地址，DST.ADDR为4个字节
	//   0x03表示域名，DST.ADDR是一个可变长度的域名
	// DST.ADDR 一个可变长度的值
	// DST.PORT 目标端口，固定2个字节

	buf := make([]byte, 4)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return fmt.Errorf("read header failed:%w", err)
	}
	ver, cmd, atyp := buf[0], buf[1], buf[3]
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}
	if cmd != cmdBind {
		return fmt.Errorf("not supported cmd:%v", ver)
	}
	addr := ""
	switch atyp {
	case atypIPV4:
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			return fmt.Errorf("read atyp failed:%w", err)
		}
		addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
	case atypeHOST:
		hostSize, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("read hostSize failed:%w", err)
		}
		host := make([]byte, hostSize)
		_, err = io.ReadFull(reader, host)
		if err != nil {
			return fmt.Errorf("read host failed:%w", err)
		}
		addr = string(host)
	case atypeIPV6:
		return errors.New("IPv6: no supported yet")
	default:
		return errors.New("invalid atyp")
	}
	_, err = io.ReadFull(reader, buf[:2])
	if err != nil {
		return fmt.Errorf("read port failed:%w", err)
	}
	port := binary.BigEndian.Uint16(buf[:2])

	log.Println("dial", addr, port)

	// +----+-----+-------+------+----------+----------+
	// |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER socks版本，这里为0x05
	// REP Relay field,内容取值如下 X’00’ succeeded
	// RSV 保留字段
	// ATYPE 地址类型
	// BND.ADDR 服务绑定的地址
	// BND.PORT 服务绑定的端口DST.PORT
	_, err = conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0})
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	return nil
}
```

## 3.3.4 SOCKS5 代理- relay 阶段

我们直接用net.dial建立一个TCP连接

建立完连接之后，我们同样要加一个defer来关闭连接。接下来需要建立浏览器和下游服务器的双向数据转发。
标准库的io.copy可以实现-个单向数据转发，双向转发的话，需要启动两个goroutinue。

```go
	_, err = io.ReadFull(reader, buf[:2])
	if err != nil {
		return fmt.Errorf("read port failed:%w", err)
	}
	port := binary.BigEndian.Uint16(buf[:2])

	dest, err := net.Dial("tcp", fmt.Sprintf("%v:%v", addr, port))
	if err != nil {
		return fmt.Errorf("dial dst failed:%w", err)
	}
	defer dest.Close()
	log.Println("dial", addr, port)
```

## 3.3.4 SOCKS5 代理- relay 阶段

我们可以试着在浏览器里面再测试一下， 在浏览器里面测试代理需要安装这个SwitchyOmega插件，然后里面新建一个情景模式，代理服务器选socks5,端口1080，保存并启用。
此时你应该还能够正常地访问网站，代理服务器这边会显示出浏览器版本的域名和端口。

![](https://cdn.jsdelivr.net/gh/nateshao/images/20220507224018.png)

## 总结

1. 首先我们学习：Go 语言学习背景介绍

2. 然后学习：Go 语言基础语言详细讲解
   - 开发环境
   - 基础语法
   - 标准库
   - Go 语言实战

3. 做项目
   - 项目一：猜谜游戏
   - 项目二：命令行词典
   - 项目三：SOCKS5 代理

源码已经上传到这里：

> 源码：https://github.com/nateshao/gogogo/tree/master/day01-05-07



参考链接：

1. 【Go 语言原理与实践学习资料】第三届字节跳动青训营-后端专场：https://juejin.cn/post/7093721879462019102
2. 《Go 语言教程》：https://www.runoob.com/go/go-tutorial.html
