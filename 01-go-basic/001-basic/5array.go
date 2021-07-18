/**
 * @date Created by 邵桐杰 on 2021/6/24 23:35
 * @微信公众号 千羽的编程时光
 * @个人网站 www.nateshao.cn
 * @博客 https://nateshao.gitee.io
 * @GitHub https://github.com/nateshao
 * Describe:
 */
package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 32, 532, 23, 21, 2, 3}
	fmt.Println(arr1, arr2, arr3)
	for i, v := range arr3 { // 遍历
		fmt.Println(i, v)
	}
}
