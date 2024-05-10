// @Description: 二叉树深度遍历，层序遍历模板, 递归遍历遍历模板
// @Author: sunding
// @Date: 2023-11-02 16:07:39
package breadth

import "algorithm-go/base"

type TreeNode = base.TreeNode

// @description: 二叉树深度遍历,层序遍历模板
// @param {*TreeNode} root
// @return {*} 返回层序遍历的切片数组
func Traverse(root *TreeNode) (res [][]int) {

	nodeQueue := []*TreeNode{}
	if root != nil {
		nodeQueue = append(nodeQueue, root)
	}

	for len(nodeQueue) > 0 {
		length := len(nodeQueue)
		curValues := []int{}

		for i := 0; i < length; i++ {
			node := nodeQueue[0]
			curValues = append(curValues, node.Value)

			if node.Left != nil {
				nodeQueue = append(nodeQueue, node.Left)
			}
			if node.Right != nil {
				nodeQueue = append(nodeQueue, node.Right)
			}
		}
		nodeQueue = nodeQueue[length:]

		res = append(res, curValues)
	}

	return res
}

// @description: 二叉树深度遍历,递归遍历模板, 前中后序遍历都可以，因为左右顺序在前中后序遍历中都是一样的
// @param {*TreeNode} root
// @return {*}
func Traverse2(root *TreeNode) (res [][]int) {

	var order func(node *TreeNode, deepth int)
	order = func(node *TreeNode, deepth int) {
		if node == nil {
			return
		}

		if len(res) == deepth {
			res = append(res, []int{})
		}

		res[deepth] = append(res[deepth], node.Value)

		order(node.Left, deepth+1)
		order(node.Right, deepth+1)
	}

	order(root, 0)
	return res
}
