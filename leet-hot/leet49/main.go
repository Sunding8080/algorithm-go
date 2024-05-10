// @Description: leet49,给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
// @Author: sunding
// @Date: 2023-12-12 20:21:06

package main

import "fmt"

func groupAnagrams(arr []string) [][]string {
	hash := map[[26]int][]string{}

	for _, v := range arr {
		count := [26]int{}
		for _, c := range v {
			count[c-'a']++
		}
		hash[count] = append(hash[count], v)
	}

	values := [][]string{}
	for _, v := range hash {
		values = append(values, v)
	}

	return values
}

func testGroupAnagrams() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func main() {
	testGroupAnagrams()
}
