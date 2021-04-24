package main

// 单个包可以这样
// import "fmt"
// 多个包
import (
	"fmt"
	"time"
)

func main() {
	//fmt.Print()// 不换行
	//fmt.Println()// 换行
	//fmt.Printf()// 格式化输出
	// fmt.Println("你好，世界！");
	// fmt.Println("hello world!");
	fmt.Println("hello world")
	time.Sleep(100)
	fmt.Println("你好，世界！")
}
