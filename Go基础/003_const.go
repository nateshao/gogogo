package main

import "fmt"

// const 定义枚举类型
// iota 只能够配合const()一起使用，iota只有在const进行累加效果
const (
	// MAOMING  = iota //默认为0
	MAOMING  = 10 * iota //默认为0
	SHANGHAI             // 默认为1
	NEIJIN               // 默认为2
	SHENZHEN             // 默认为3
)

func main() {
	// 常量(只读属性)
	const length int = 10
	fmt.Println("length = ", length)
	fmt.Println("MAOMING = ", MAOMING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("NEIJIN = ", NEIJIN)
	fmt.Println("SHENZHEN = ", SHENZHEN)
}
