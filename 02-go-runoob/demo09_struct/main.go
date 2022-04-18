package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	// 创建一个新的结构体
	fmt.Println(Books{"千羽的编程时光", "千羽", "Go语言实战", 12})
	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "千羽的编程时光", author: "千羽", subject: "Go开发实战", book_id: 12})
}
