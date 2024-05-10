// @Description: leet437,给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
// @Author: sunding
// @Date: 2023-12-14 01:55:39
package main

import (
	"algorithm-go/base"
	"fmt"
)

type TreeNode = base.TreeNode

func traverse(node *TreeNode, target int, result *[][]int, path []int) {

	if node == nil {
		return
	}

	target = target - node.Value
	path = append(path, node.Value)

	if (node.Left == nil) && (node.Right == nil) && (target == 0) {
		*result = append(*result, path)
		return
	}

	traverse(node.Left, target, result, path)
	traverse(node.Right, target, result, path)
}

func pathSum(root *TreeNode, target int) [][]int {
	result := [][]int{}
	traverse(root, target, &result, []int{})
	return result
}

func testPathSum() {
	values := []int{1, 3, 4, 3, 2, 1, -1, -1, -1, 1}
	root := base.SortIntsToTree(values)
	result := pathSum(root, 7)
	fmt.Println("result:", result)
}

func main() {
	testPathSum()
}
