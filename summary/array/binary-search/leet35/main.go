// @Description: leetcode35(二分查找): 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// @Author: sunding
// @Date: 2023-11-01 14:12:09
package main

import "fmt"

// @description:  二分查找, 左闭右开，完整思路
// @param {[]int} nums 排序的数组
// @param {int} target 目标元素
// @param {int} index 目标元素数组下标
// @return {*}
func search1(nums []int, target int) (index int) {
	length := len(nums)
	left := 0
	right := length - 1
	index = -1
	mid := -1

	for left < right {
		mid = (left + right) / 2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid
		} else {
			// 找到目标
			index = mid
			break
		}
	}

	// 没有找到, nums[mid]和target比较
	// 最后一次判断时 [left, right), [mid, mid + 1), mid = left, right = left + 1
	// 可以简化成函数search2
	if index == -1 {
		if target > nums[mid] {
			index = mid + 1 // 此时,right的值是mid + 1, 所以index = right， index = left
		}
		if target < nums[mid] {
			index = mid // 此时,right的值是mid， 所以index = right, index = left
		}
	}

	return index
}

// @description:  二分查找, 左闭右开，在Search1的基础上简化, 合并部分逻辑
// @param {[]int} nums 排序的数组
// @param {int} target 目标元素
// @param {int} index 目标元素数组下标
// @return {*}
func search2(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := length - 1

	for left < right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid
		} else {
			// 找到目标
			return mid
		}
	}

	// 这里返回left、right都是可以的，根据推算都符合
	// return left
	return right
}

// @description:  二分查找, 左闭右闭, 详细流程
// @param {[]int} nums 排序的数组
// @param {int} target 目标元素
// @param {int} index 目标元素数组下标
// @return {*}
func search3(nums []int, target int) (index int) {
	left := 0
	right := len(nums)
	mid := -1
	index = -1

	for left <= right {
		mid = (left + right) / 2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			index = mid
			break
		}
	}

	// 没有找到, nums[mid]和target比较
	//最后一次判断时 [left, right], [mid, mid], mid = left = right
	// 可以简化成函数search4
	if index == -1 {
		if target < nums[mid] {
			index = mid // 此时,right的值是mid - 1, index = right + 1，index = left
		}
		if target > nums[mid] {
			index = mid + 1 // 此时,right的值是mid, index = right + 1，index = left
		}
	}

	return index
}

// @description:  二分查找, 左闭右闭, 在Search3的基础上简化, 合并部分逻辑
// @param {[]int} nums 排序的数组
// @param {int} target 目标元素
// @param {int} index 目标元素数组下标
// @return {*}
func search4(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := length - 1

	for left <= right {
		mid := (left + right) / 2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}

	// return left
	return right + 1
}

func testSearch() {
	arr := []int{16, 28, 35, 49, 55, 67, 72, 81, 96, 109}

	index1 := search1(arr, 20)
	fmt.Println("search1(arr, 20):", index1)
	index2 := search1(arr, 30)
	fmt.Println("search1(arr, 30):", index2)

	index3 := search2(arr, 40)
	fmt.Println("search2(arr, 40):", index3)
	index4 := search2(arr, 50)
	fmt.Println("search2(arr, 50):", index4)

	index5 := search3(arr, 60)
	fmt.Println("search3(arr, 60):", index5)
	index6 := search3(arr, 73)
	fmt.Println("search3(arr, 70):", index6)

	index7 := search4(arr, 81)
	fmt.Println("search4(arr, 80):", index7)
	index8 := search4(arr, 90)
	fmt.Println("search4(arr, 99):", index8)
}

func main() {
	testSearch()
}
