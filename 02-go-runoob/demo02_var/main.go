package main

import (
	"fmt"
)

func main() {
	// 声明一个变量并初始化
	var name = "邵桐杰"
	fmt.Println(name)

	var b, c int = 2, 3
	fmt.Println(b, c)
	// 没有初始化就为零值
	var d int
	fmt.Println(d)
	// bool 零值为 false
	var f bool
	fmt.Println(f)

	fmt.Println("-----------------------")
	var i int
	var f2 float64
	var b2 bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f2, b2, s)

	z := "邵桐杰"
	fmt.Println(z)
}
