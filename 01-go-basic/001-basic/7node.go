/**
 * @date Created by 邵桐杰 on 2021/8/19 22:08
 * @微信公众号 千羽的编程时光
 * @个人网站 www.nateshao.cn
 * @博客 https://nateshao.gitee.io
 * @GitHub https://github.com/nateshao
 * Describe:
 */
package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Print(node.value)
}
func createTreeNode(value int) *treeNode {
	return &treeNode{value: value}
}
func main() {
	var root treeNode
	root.left = &treeNode{value: 3}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.print()
	fmt.Println()
}
