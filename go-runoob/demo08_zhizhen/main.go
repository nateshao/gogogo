package main

import "fmt"

func main() {
	var a int = 10

	fmt.Printf("变量的地址: %x\n", &a)
	fmt.Printf("a的值：%d\n", a)

}
