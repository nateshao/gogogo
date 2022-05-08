package main

import "fmt"

func main() {

	/**
	2 到 100 间的素数
	*/
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break //如果发现因子则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d 是素数\n", i)
		}
	}
}
