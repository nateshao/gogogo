package main

import (
	"fmt"
)

func main() {

	var myMap1 map[string]string

	if myMap1 == nil {
		fmt.Println("myMap1 is nil..")
	}
	// 开辟空间
	myMap1 = make(map[string]string, 10)
	// 赋值
	myMap1["one"] = "java"
	myMap1["two"] = "python"
	myMap1["three"] = "Go"
	fmt.Println(myMap1)

	fmt.Println("=========第二种声明方式==========")
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "rust"
	fmt.Println(myMap2)

}
