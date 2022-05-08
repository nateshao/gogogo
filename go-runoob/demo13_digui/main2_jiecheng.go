package main

import "fmt"

func Factor(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factor(n-1)
		return result
	}
	return 1
}
func main() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factor(uint64(i)))
}
