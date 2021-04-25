package main

/*
	四种变量的声明方式
	s := ""
	var s string
	var s = ""
	var s string = ""
	第一种形式，是一条短变量声明，最简洁，但只能用在函数内部，而不能用于包变量。
	第二种形式依赖于字符串的默认初始化零值机制，被初始化为""。
	第三种形式用得很少，除非同时声明多个变量。
	第四种形式显式地标明变量的类型，当变量类型与初值类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了。实践中一般使用前两种形式中的某个，初始值重要的话就显式地指定变量的类型，否则使用隐式初始化。
*/

import "fmt"

// 声明全局变量 方法一，方法二，方法三是可以的
var gA int = 100
var gB = 324

//gC := 321
// := 只能在函数体内声明

func main() {
	// fmt.Println("hello world")
	// 法1：声明一个变量 默认值为0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	// 法2：声明一个变量 默认值为0
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "bbbc"
	fmt.Printf("bb = %s,type of bb = %T\n", bb, bb)

	// 法3：在初始华的时候，可以省区数据类型通过值自动匹配的变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	// 法4：(常用的方法)省去var关键字，直接自动匹配
	d := 100 // d初始化并赋值100
	fmt.Println("d = ", d)
	fmt.Printf("type of d = %T\n", d)

	e := "eeee"
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	g := 3.44
	fmt.Println("g = ", g)

	fmt.Println("=======================================")

	fmt.Println("gA = ", gA)
	fmt.Println("gB = ", gB)
	// fmt.Println("gC = ", gC)

	// 声明多个变量
	var xx, yy int = 123, 456
	fmt.Println("xx = ", xx, "yy = ", yy)

	var aa, ss = 12, "fkh"
	fmt.Println("aa = ", aa, "ss = ", ss)

	// 多行变量的声明
	var (
		vv int  = 110
		jj bool = true
	)
	fmt.Println("vv = ", vv, "jj = ", jj)
	qq := 0.1
	ww := 0.2
	sum := qq + ww // 只要使用浮点数，0.1就不能在内存中完全表示，所以我们知道这个值通常会出现在0.10000000000000004之间。

	fmt.Println("qq + ww = ", qq+ww) // sum =  0.30000000000000004
	fmt.Println("sum = ", sum)       // sum =  0.30000000000000004
	fmt.Printf("%.2f\n", 0.1+0.2)    // 0.30
	fmt.Printf("%.64f\n", 0.1+0.2)   // 0.2999999999999999888977697537484345957636833190917968750000000000
	fmt.Println(0.1 + 0.2)           // 0.3
}
