// @Description: leet105，从前序与中序遍历序列构造二叉树
// @Author: sunding
// @Date: 2023-12-14 00:27:16
package main

import (
	"algorithm-go/base"
	"fmt"
)

type TreeNode = base.TreeNode

// @description: 递归
// @param {[]int} preArr 前序数组
// @param {[]int} inArr 中序数组
// @return {*}
func buildTree(preArr []int, inArr []int) *TreeNode {
	if len(preArr) == 0 || len(inArr) == 0 {
		return nil
	}

	var index int
	value := preArr[0]

	for i, v := range inArr {
		if v == value {
			index = i
			break
		}
	}

	leftInArr := inArr[0:index]
	rightInArr := inArr[index+1:]

	leftInArrlen := len(leftInArr)
	leftPreArr := preArr[1 : leftInArrlen+1]
	rightPreArr := preArr[leftInArrlen+1:]

	node := &TreeNode{
		Value: value,
		Left:  buildTree(leftPreArr, leftInArr),
		Right: buildTree(rightPreArr, rightInArr),
	}

	return node
}

func testBuildTree() {
	preArr := []int{1, 2, 4, 5, 3, 6, 7}
	inArr := []int{4, 2, 5, 1, 6, 3, 7}

	root := buildTree(preArr, inArr)
	fmt.Println("root:", root)
}

func main() {
	testBuildTree()
}
