// @Description: leet283,给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。请注意 ，必须在不复制数组的情况下原地对数组进行操作。
// @Author: sunding
// @Date: 2023-12-12 22:21:47
package main

import "fmt"

// @description: 双指针法：遍历过程
// @param {*[]int} p
// @return {*}
func moveZeroes(p *[]int) {
	nums := *p

	left, right, len := 0, 0, len(nums)
	for right < len {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func testMoveZeroes() {
	// arr := []int{0, 1, 0, 3, 12}
	arr := []int{1, 2, 0, 3, 0, 4, 0, 5}

	moveZeroes(&arr)
	fmt.Println("arr:", arr)
}

func main() {
	testMoveZeroes()
}
