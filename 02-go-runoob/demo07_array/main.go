package main

import "fmt"

func main() {
	var arr [10]int
	fmt.Println(arr)
	var i, j int

	for i = 0; i < 10; i++ {
		arr[i] = i + 100

	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, arr[j])
	}
}
