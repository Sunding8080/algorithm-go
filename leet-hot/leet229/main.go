// @Description: leet229,给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
// @Author: sunding
// @Date: 2023-12-13 22:40:52
package main

import (
	"algorithm-go/base"
)

type TreeNode = base.TreeNode

// @description: 递归法处理
// @param {*TreeNode} node
// @return {*}
func invertTree(node *TreeNode) *TreeNode {
	if node != nil {
		return node
	}

	node.Left, node.Right = node.Right, node.Left

	if node.Left != nil {
		invertTree(node.Left)
	}

	if node.Right != nil {
		invertTree(node.Right)
	}

	return node
}

func testInvertTree() {
	values := []int{3, 9, 20, 0, 0, 15, 7}
	root := base.SortIntsToTree(values)
	invertTree(root)
}

func main() {
	testInvertTree()
}
