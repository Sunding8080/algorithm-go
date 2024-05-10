// @Description: 二叉树深度遍历，相关leetcode解析
// @Author: sunding
// @Date: 2023-11-02 16:26:41
package main

import (
	"algorithm-go/base"
	"algorithm-go/summary/tree/breadth-traverse/breadth"
	"fmt"
)

type TreeNode = base.TreeNode

// @description: leetcode102：给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
// @param {breadth.TreeNode} root
// @return {*}
func testLeet102() (res [][]int) {
	arr := []int{1, 2, 3, -1, -1, 6, 7}
	root := base.SortIntsToTree(arr)
	res = breadth.Traverse(root)

	fmt.Println("testLeet102:", res)
	return res
}

// @description: leetcode107：给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
// @return {*}
func testLeet107() (res [][]int) {
	arr := []int{1, 2, 3, -1, -1, 6, 7}
	root := base.SortIntsToTree(arr)
	res2 := breadth.Traverse2(root)
	length := len(res2)

	for i := length - 1; i >= 0; i-- {
		res = append(res, res2[i])
	}

	fmt.Println("testLeet107:", res)
	return res
}

// @description: leetcode199：给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
// @return {*}
func testLeet199() (res []int) {
	arr := []int{1, 2, 3, 0, 5, 0, 4}
	root := base.SortIntsToTree(arr)
	res2 := breadth.Traverse2(root)
	length := len(res2)

	for i := 0; i < length; i++ {
		arr := res2[i]
		res = append(res, arr[len(arr)-1])
	}

	fmt.Println("testLeet199:", res)
	return res
}

// @description: leetcode637：给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。
// @return {*}
func testLeet637() (res []float32) {
	arr := []int{3, 9, 20, 0, 0, 15, 7}
	root := base.SortIntsToTree(arr)
	res2 := breadth.Traverse2(root)

	for _, v1 := range res2 {
		sum := 0

		for _, v2 := range v1 {
			sum += v2
		}
		res = append(res, float32(sum)/float32(len(v1)))
	}

	fmt.Println("testLeet637:", res)
	return res
}

// @description: leetcode104：给定一个二叉树，找出其最大深度。二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。说明: 叶子节点是指没有子节点的节点。
// @return {*}
func testLeet104() int {
	arr := []int{3, 9, 20, 0, 0, 15, 7}
	root := base.SortIntsToTree(arr)
	res := breadth.Traverse2(root)

	return len(res)
}

// @description: 获取二叉树的最小深度
// @param {*TreeNode} root
// @return {*}
func getMinDeepth(root *TreeNode) (deepth int) {

	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	deepth = 0

	for len(queue) > 0 {
		length := len(queue)

		for i := 0; i < length; i++ {
			node := queue[i]
			if node.Left == nil && node.Right == nil {
				return deepth + 1
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		deepth++
		queue = queue[length:]

	}

	return deepth
}

// @description: leetcode111:二叉树的最小深度
// @return {*}
func testLeet111() int {
	arr := []int{3, 9, 20, 0, 0, 15, 7}
	root := base.SortIntsToTree(arr)
	deepth := getMinDeepth(root)

	fmt.Println("testLeet111:", deepth)
	return deepth
}

func main() {
	// 以下都是一个系列
	testLeet102()
	testLeet107()
	testLeet199()
	testLeet637()
	testLeet104()

	testLeet111()
}
