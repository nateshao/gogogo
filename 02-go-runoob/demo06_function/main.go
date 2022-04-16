package main

import "fmt"

func main() {
	/* 调用函数并返回最大值 */
	fmt.Println(max(10, 2))
}

func max(num1, num2 int) int {
	/* 定义局部变量 */
	var res int
	if num1 > num2 {
		res = num1
	} else {
		res = num2
	}
	return res
}
