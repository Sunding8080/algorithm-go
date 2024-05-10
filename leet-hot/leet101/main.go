// @Description: leet101给你一个二叉树的根节点 root ， 检查它是否轴对称。
// @Author: sunding
// @Date: 2023-12-13 22:57:01
package main

import (
	"algorithm-go/base"
	"fmt"
)

type TreeNode = base.TreeNode

func compare(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right != nil {
		return false
	} else if left != nil && right == nil {
		return false
	} else if left.Value != right.Value {
		return false
	} else {
		return compare(left.Left, right.Right) && compare(left.Right, right.Left)
	}
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

func testIsSymmetric() {
	values := []int{3, 9, 20, 0, 0, 15, 7}
	root := base.SortIntsToTree(values)
	res := isSymmetric(root)
	fmt.Println("res:", res)
}

func main() {
	testIsSymmetric()
}
