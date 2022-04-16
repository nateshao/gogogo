package main

import "fmt"

func main() {
	a, b := swap("千羽", "千寻")
	fmt.Println(a, b)
}
func swap(x, y string) (string, string) {
	return y, x
}
