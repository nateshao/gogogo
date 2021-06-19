package main

import "fmt"

func main() {
	//写入defer关键字
	defer fmt.Println("main end1")
	defer fmt.Println("main end2") // 栈的形式

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2") // return执行早于defer
}
