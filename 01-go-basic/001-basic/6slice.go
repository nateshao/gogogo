/**
 * @date Created by 邵桐杰 on 2021/7/18 21:00
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

// 切片
func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[2:6]
	fmt.Println("arrs 的值为= ", s)
}
