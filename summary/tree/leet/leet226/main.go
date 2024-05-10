// @Description: leetcode226, 翻转一棵二叉树
// @Author: sunding
// @Date: 2023-11-03 21:51:12
package main

/*
* 只需要将左右节点转换就可以实现
* 解法：
* 1 递归深度遍历(前后)
* 2 迭代深度遍历（前中后）
* 3 广度遍历（层序遍历）
 */

import (
	"algorithm-go/base"
	"fmt"
)

type TreeNode = base.TreeNode

// @description: 先序遍历，递归法, （后序递归法也行， 中序递归法不行）
// @param {*TreeNode} root 根节点， []int 前序遍历结果
// @return {*}
func traverse(root *TreeNode) (*TreeNode, []int) {

	res := []int{}
	var order func(node *TreeNode)
	order = func(node *TreeNode) {
		if node == nil {
			return
		}

		node.Left, node.Right = node.Right, node.Left

		res = append(res, node.Value)

		order(node.Left)
		order(node.Right)
	}

	order(root)

	return root, res
}

// @description: 广度遍历，层序遍历
// @param {*TreeNode} root 根节点， []int 广度遍历结果
// @return {*}
func traverse2(root *TreeNode) (*TreeNode, []int) {
	res := []int{}
	queue := []*TreeNode{}

	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		length := len(queue)

		for i := 0; i < length; i++ {
			node := queue[i]
			node.Left, node.Right = node.Right, node.Left

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			res = append(res, node.Value)
		}

		queue = queue[length:]

	}

	return root, res
}

// @description: 前序遍历
// @return {*}
func testTraverse() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	root1 := base.SortIntsToTree(arr)
	fmt.Println("root1:", root1)
	antiRoot1, res1 := traverse(root1)
	fmt.Println("antiRoot1:", antiRoot1, ", res1:", res1)

	root2 := base.SortIntsToTree(arr)
	fmt.Println("root2:", root2)
	antiRoot2, res2 := traverse2(root2)
	fmt.Println("antiRoot2:", antiRoot2, ", res2:", res2)
}

func main() {
	testTraverse()
}
