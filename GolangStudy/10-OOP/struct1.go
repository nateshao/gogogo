package main

import "fmt"

//type myint int
type Book1 struct {
	title string
	name  string
	auth  string
	id    int
}

func changeBook1(book Book) {
	book.title = "gogogo.."
}

func main() {

	/*	var a myint
		fmt.Println("a = ", a)
		fmt.Printf("a type is %T\n", a)
	*/

	// 定义book1变量
	var book1 Book1
	// 赋值
	book1.id = 1
	book1.name = "Go圣经"
	book1.auth = "张三"
	book1.title = "Golang 实战"
	fmt.Println(book1.title, book1.auth, book1.name, book1.id)
	fmt.Printf("%v\n", book1)
	changeBook1(book1)
	fmt.Printf("%v\n", book1)

}
