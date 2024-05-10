// @Description: leetcode704，给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
// @Author: sunding
// @Date: 2023-10-27 13:20:19
package main

import "fmt"

// @description:  二分查找, 左闭右开
// @param {[]int} nums 排序的数组
// @param {int} target 目标元素
// @param {int} index 目标元素数组下标
// @return {*}
func search1(nums []int, target int) (index int) {
	length := len(nums)
	left := 0
	right := length - 1

	// 处于中间
	for left < right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid
		} else if target == nums[mid] {
			return mid
		}
	}

	return -1
}

// @description:  二分查找, 左闭右开
// @param {[]int} nums 排序的数组
// @param {int} target 目标元素
// @param {int} index 目标元素数组下标
// @return {*}
func search2(nums []int, target int) (index int) {
	length := len(nums)
	left := 0
	right := length - 1

	// 处于中间位置
	for left <= right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else if target == nums[mid] {
			return mid
		}
	}

	return -1
}

// @description: 二分查找
// @return {*}
func testBinarySearch() {
	arr := []int{16, 28, 35, 49, 55, 67, 72, 81, 99, 109}
	index := search1(arr, 55)
	fmt.Println("res:", index)
	index2 := search2(arr, 99)
	fmt.Println("res2:", index2)
}

func main() {
	testBinarySearch()
}
