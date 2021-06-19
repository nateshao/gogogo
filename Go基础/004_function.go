// package main

// import "fmt"

// /*	函数名（a 类型，b类型）									*/
// func fool(a string, b int) int {
// 	fmt.Println("a = ", a)
// 	fmt.Println("b = ", b)

// 	c := 100 // 初始化并赋值100
// 	return c
// }

// func main() {
// 	fmt.Println("突突突，起飞")
// 	c := fool("abc123", 23) //调用函数并初始化赋值给c
// 	fmt.Println("c = ", c)
// }
/*********************************************************************************8888*/
/* package main

import "fmt"

func bool(a string, b int) int {
	fmt.Println("a = ", a, "b = ", b)
	c := 123
	return c
}

func main() {
	fmt.Println("aaa")
	c := bool("qianyu", 335)
	fmt.Println("c = ", c)
} */

/******************************************************/

package main

import "fmt"

// 返回多个返回值，匿名的
func bool2(a string, b int) (int, int) {

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 66, 77

}

func bool3(a string, b int) (r1 int, r2 int) {
	fmt.Println("bool3--------------------")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r1 = 1000
	r2 = 2000
	return
}

func main() {
	fmt.Println("aaaaaa")
	ret1, ret2 := bool2("a1234", 32)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)
	ret1, ret2 = bool3("bool33333", 33333)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)
	fmt.Println(ret1 + ret2)
}
