package main // package main代表这个文件属于main包的一部分，main 包也就是程序的入口包。

import (
	"fmt" // 标准库里面的FMT包。这个包主要是用来往屏幕输入输出字符串、格式化字符串。
)

func main() {
	fmt.Println("hello world") // main 函数的话里面调用了fmt.Println 输出helloword
}
