/**
 * @date Created by 邵桐杰 on 2021/7/19 21:13
 * @微信公众号 千羽的编程时光
 * @个人网站 www.nateshao.cn
 * @博客 https://nateshao.gitee.io
 * @GitHub https://github.com/nateshao
 * Describe:
 */
package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"name":     "ccmouse",
		"course":   "golang",
		"username": "zhangsan",
		"site":     "cainiao",
	}
	fmt.Println(m)
	m2 := make(map[string]int)
	var m3 map[string]int

	fmt.Println(m, m2, m3)
	for k, v := range m { // 遍历输出map m1---打印key和value
		fmt.Println("key =", k, " value = ", v)
	}

	for k := range m {
		fmt.Println(k) // 只打印key
	}
}
