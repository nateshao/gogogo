package main

import "fmt"

type myint int

func main() {

	var a myint
	fmt.Println("a = ", a)
	fmt.Printf("a type is %T\n", a)
}
