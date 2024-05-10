// @Description: 二叉树先序遍历
// @Author: sunding
// @Date: 2023-11-02 14:38:05
package main

import (
	"algorithm-go/base"
	"fmt"
)

type TreeNode = base.TreeNode

// @description: 二叉树先序遍历，递归法
// @param {*TreeNode} root
// @return {*}
func traverse(root *TreeNode) (res []int) {

	var order func(node *TreeNode, arr *([]int))
	order = func(node *TreeNode, arr *([]int)) {
		if node != nil {
			res = append(res, node.Value)
			order(node.Left, arr)
			order(node.Right, arr)
		}
	}

	order(root, &res)

	return res
}

// @description: 二叉树先序遍历，迭代法
// @param {*TreeNode} root
// @return {*}
func traverse2(root *TreeNode) (res []int) {

	stack := []*TreeNode{}

	if root != nil {
		stack = append(stack, root)
	}

	for len(stack) > 0 {
		length := len(stack)
		node := stack[length-1]
		stack = stack[:length-1]

		if node == nil {
			opLength := len(stack)
			opNode := stack[opLength-1]
			res = append(res, opNode.Value)
			stack = stack[:opLength-1]
			continue
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		stack = append(stack, node, nil)
	}

	return res
}

func testTraverse() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	root := base.SortIntsToTree(arr)

	res1 := traverse(root)
	res2 := traverse2(root)
	fmt.Println("res1:", res1)
	fmt.Println("res2:", res2)
}

func main() {
	testTraverse()
}
