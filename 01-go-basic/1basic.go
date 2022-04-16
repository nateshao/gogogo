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
	fmt.Printf("%d %q\n", a, s) // // 格式化输出 0 ""
}

func variableInitialValue() {
	var a, b int = 3, 4  // int 这里可以省略 因为Go语言会识别到
	var s string = "abc" // string 这里可以省略 因为Go语言会识别到
	fmt.Println(a, b, s) // 3 4 abc
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s) // 3 4 true def
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s) // 3 5 true def
}

func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1) // 3 kkk true
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b)) // (0.000+0.000i) 5
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
	fmt.Println(filename, c) // abc.txt 5
}

func enums() {
	const ( // iota 自增值
		cpp        = iota // 0
		_                 // 1
		python            //	2
		golang            // 3
		javascript        // 4
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang) // 0 4 2 3
	fmt.Println(b, kb, mb, gb, tb, pb)           // 1 1024 1048576 1073741824 1099511627776 1125899906842624
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
