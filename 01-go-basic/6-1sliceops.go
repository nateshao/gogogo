/**
 * @date Created by 邵桐杰 on 2021/7/18 23:28
 * @微信公众号 千羽的编程时光
 * @个人网站 www.nateshao.cn
 * @博客 https://nateshao.gitee.io
 * @GitHub https://github.com/nateshao
 * Describe:
 */
package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n",
		s, len(s), cap(s))
}

func main() {
	fmt.Println("Creating slice")
	var s []int // Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	fmt.Println("------------ make s2 --------------", s2, len(s2), cap(s2))
	fmt.Println("------------ make s3 --------------", s3, "len = ", len(s3), "cap =", cap(s3))
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	printSlice(s2)
}
