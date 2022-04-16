package main

import "fmt"

func main() {
	strings := []string{"千羽", "千寻"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	nums := [6]int{1, 2, 3, 6}
	for i, x := range nums {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)

	}
}
