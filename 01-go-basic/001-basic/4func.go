/**
 * @date Created by 邵桐杰 on 2021/6/19 22:56
 * @微信公众号 千羽的编程时光
 * @个人网站 www.nateshao.cn
 * @博客 https://nateshao.gitee.io
 * @GitHub https://github.com/nateshao
 * Describe:
 */
package main

import "fmt"

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation:" + op)
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	fmt.Println(eval(3, 4, "-"))
	fmt.Println(eval(3, 4, "+"))
	fmt.Println(eval(3, 4, "/"))

	fmt.Println(div(13, 3)) // 4 1
}
