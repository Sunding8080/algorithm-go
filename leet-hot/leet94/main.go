// @Description: leetcode94，给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
// @Author: sunding
// @Date: 2023-12-13 17:07:35
package main

import (
	"algorithm-go/base"
	"fmt"
)

type TreeNode = base.TreeNode

func traverse(root *TreeNode) (res []int) {
	var order func(node *TreeNode, res *([]int))
	order = func(node *TreeNode, res *([]int)) {
		if node != nil {
			order(node.Left, res)
			*res = append(*res, node.Value)
			order(node.Right, res)
		}
	}

	return
}

func testTraverse() {
	values := []int{3, 9, 20, 0, 0, 15, 7}
	root := base.SortIntsToTree(values)
	res := traverse(root)
	fmt.Println("res:", res)
}

func main() {
	testTraverse()
}
