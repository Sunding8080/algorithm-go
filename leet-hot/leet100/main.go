// @Description: leet100,给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// @Author: sunding
// @Date: 2023-12-12 19:16:42
package main

import "fmt"

// @description: 哈希法解题
// @param {int} target
// @param {[]int} arr
// @return {*}
func twoSum(target int, arr []int) []int {
	hash := map[int]int{}

	for i, v := range arr {
		if _, ok := hash[target-v]; ok {
			return []int{v, i}
		}
		hash[v] = i
	}

	return nil
}

func testTwoSum() {

	target := 6
	arr := []int{3, 2, 4, 1, 9}

	fmt.Println("target:", target, "arr:", arr, "result:", twoSum(target, arr))
}

func main() {
	testTwoSum()
}
