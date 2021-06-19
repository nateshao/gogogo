package main

import (
	"fmt"
)

func main() {
	// 定义数组
	var myArray1 [10]int
	myArray2 := [4]int{1, 2, 3, 4}
	myArray3 := [10]int{111, 1112, 2223, 444}

	// 遍历数组
	for i := 0; i < len(myArray2); i++ {
		fmt.Println(myArray2[i])
	}

	// 数组遍历二
	for index, value := range myArray2 {
		fmt.Println("index = ", index, "value = ", value)
	}

	// 查看数组的类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

	// 遍历数组三
	for index, value := range myArray3 {
		fmt.Println("index = ", index, "value = ", value)

	}

}
