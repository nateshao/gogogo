/**
 * @date Created by 邵桐杰 on 2021/6/19 21:27
 * @微信公众号 千羽的编程时光
 * @个人网站 www.nateshao.cn
 * @博客 https://nateshao.gitee.io
 * @GitHub https://github.com/nateshao
 * Describe: 5 的二进制
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func converToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsp := n % 2
		result = strconv.Itoa(lsp) + result

	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		converToBin(5),  // 101
		converToBin(13), // 1101
		converToBin(22222222221231),
		converToBin(0), // 没显示
	)
	printFile("abc.txt")
}
