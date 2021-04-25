package main

import (
	"fmt"
)

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called..")
	fmt.Println(arg)

	// 给interface{}提供一个类的断言的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)
		fmt.Printf("value type is %T\n", value)
	}

}

type Book struct {
	auth string
}

func main() {

	book := Book{"Golang"}
	myFunc(book)
	myFunc(100)
	myFunc("zbc")
	myFunc(3.1415)
}
