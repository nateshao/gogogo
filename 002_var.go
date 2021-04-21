package main

/*
	四种变量的声明方式
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
	fmt.Println("gC = ", gC)
}
