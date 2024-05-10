// @Description: 给定一个二叉树 root ，返回其最大深度。二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。
// @Author: sunding
// @Date: 2023-12-13 21:58:23
package main

import (
	"algorithm-go/base"
	"fmt"
)

type TreeNode = base.TreeNode

// @description: 层序遍历
// @param {*TreeNode} root
// @return {*}
func getMaxDepth(root *TreeNode) int {
	stack := []*TreeNode{root}
	res := [][]int{}
	for len(stack) > 0 {
		length := len(stack)
		values := []int{}
		for i := 0; i < length; i++ {
			node := stack[i]
			values = append(values, node.Value)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		res = append(res, values)
		stack = stack[length:]
	}
	return len(res)
}

func testGetMaxDepth() {
	values := []int{3, 9, 20, 0, 0, 15, 7}
	root := base.SortIntsToTree(values)
	length := getMaxDepth(root)
	fmt.Println("length:", length)
}

func main() {
	testGetMaxDepth()
}
