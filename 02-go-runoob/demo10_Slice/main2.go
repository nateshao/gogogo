package main

import "fmt"

func main() {
	var numbers []int

	printSlice2(numbers)

	if numbers == nil {
		fmt.Printf("切片是空的")
	}
}

func printSlice2(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
