package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3} //len=3,cap=3
	s2 := s1[0:2]        // [1,2]
	fmt.Println(s2)

	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)

	s := make([]int, 3)
	copy(s, s2)
	fmt.Println(s)
}
