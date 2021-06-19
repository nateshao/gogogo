# Google资深工程师深度讲解Go语言-笔记

#### Description

```go
%d          十进制整数
%x, %o, %b  十六进制，八进制，二进制整数。
%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
%t          布尔：true或false
%c          字符（rune） (Unicode码点)
%s          字符串
%q          带双引号的字符串"abc"或带单引号的字符'c'
%v          变量的自然形式（natural format）
%T          变量的类型
%%          字面上的百分号标志（无操作数）
```
## 第2章 基础语法
### 2-1 变量定义

> go语言Print、Println 、Printf 、Sprintf 、Fprintf的区别
```go
Print：输出内容，结尾没有换行4ab
Println：输出内容，结尾自带回车换行4ab
Fprintf：格式化字符串并输出，如：fmt.Sprintf("是字符串 %s ","string")dc5ba5a
Sprintf ：格式化字符串并返回，可以用来赋值。不会输出。2e58
 s := fmt.Sprintf("是字符串 %s ","string") 757ad
```
- Go语言中，你定义了变量，没使用的话，会报错
```go
/**
 * @date Created by 邵桐杰 on 2021/6/19 11:44
 * @微信公众号 千羽的编程时光
 * @个人网站 www.nateshao.cn
 * @博客 https://nateshao.gitee.io
 * @GitHub https://github.com/nateshao
 */
package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4  // int 这里可以省略 因为Go语言会识别到
	var s string = "abc" // string 这里可以省略 因为Go语言会识别到
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()
	consts()
	enums()
}
```

### 2-2 内建变量类型
```go
reflect.TypeOf()可以判断类型
派生类型:包括：
(a) 指针类型（Pointer）
(b) 数组类型
© 结构化类型(struct)
(d) Channel 类型
(e) 函数类型
(f) 切片类型
(g) 接口类型（interface）
(h) Map 类型
```

| 数字类型                                                     | 其他类型                                      | 浮点型                                 |      |      |
| ------------------------------------------------------------ | --------------------------------------------- | -------------------------------------- | ---- | ---- |
| 1	uint8<br/>无符号 8 位整型 (0 到 255)                    | 1	byte<br/>类似 uint8                      | 1	float32<br/>IEEE-754 32位浮点型数 |      |      |
| 2	uint16<br/>无符号 16 位整型 (0 到 65535)                | 2	rune<br/>类似 int32                      | 2	float64<br/>IEEE-754 64位浮点型数 |      |      |
| 3	uint32<br/>无符号 32 位整型 (0 到 4294967295)           | 3	uint<br/>32 或 64 位                     | 3	complex64<br/>32 位实数和虚数     |      |      |
| 4	uint64<br/>无符号 64 位整型 (0 到 18446744073709551615) | 4	int<br/>与 uint 一样大小                 | 4	complex128<br/>64 位实数和虚数    |      |      |
| 5	int8<br/>有符号 8 位整型 (-128 到 127)                  | 5	uintptr<br/>无符号整型，用于存放一个指针 |                                        |      |      |
| 6	int16<br/>有符号 16 位整型 (-32768 到 32767)            |                                               |                                        |      |      |
| 7	int32<br/>有符号 32 位整型 (-2147483648 到 2147483647)  |                                               |                                        |      |      |
| 8	int64<br/>有符号 64 位整型 (-9223372036854775808 到 9223372036854775807) |                                               |                                        |      |      |

### 2-3 常量与枚举

定义常量使用关键字const，go语言常量不用大写
```go
func contst()  {
    const name  = "abc.txt" // 常量一般是小写
    const a,b = 3,4
    var c int
    c = int(math.Sqrt(a*a+b*b))
    fmt.Println(c)
}

```
枚举 特殊的常量类型枚举
```go
func enums()  {
    const (
    java = 0
    golang = 1
    python =2
    scala = 3
)
    fmt.Println(java,golang,python,scala)
}
```

- Go语言没有char，只有rune

## 2-4 条件语句













































