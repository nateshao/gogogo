package main

import "fmt"

func main() {
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	// 这样写也可以，更像 While 语句形式
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	/**
	func main() {
	        sum := 0
	        for {
	            sum++ // 无限循环下去
	        }
	        fmt.Println(sum) // 无法输出
	}
	*/
}
