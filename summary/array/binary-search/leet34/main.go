// @Description: 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。如果数组中不存在目标值 target，返回 [-1, -1]。
// @Author: sunding
// @Date: 2023-11-01 16:31:08
package main

import "fmt"

func findLeft(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func findLeft2(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2

		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func findRight(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

// func findRight2(nums []int, target int) int {

// }

// func findRange(nums []int, target int) [2]int {

// }

func testSearch() {
	arr := []int{16, 28, 34, 34, 34, 34, 34, 60, 60, 70, 70, 100}
	fmt.Println(findLeft(arr, 34))
	// fmt.Println(findLeft2(arr, 70))
}

func main() {
	testSearch()
}
