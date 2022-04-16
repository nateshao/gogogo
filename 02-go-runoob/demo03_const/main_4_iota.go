package main

import "fmt"

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

/**
iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），即：i=1, j=6，这没问题，
关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3。

简单表述:

i=1：左移 0 位，不变仍为 1。
j=3：左移 1 位，变为二进制 110，即 6。
k=3：左移 2 位，变为二进制 1100，即 12。
l=3：左移 3 位，变为二进制 11000，即 24。
注：<<n==*(2^n)。
*/
func main() {
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}
