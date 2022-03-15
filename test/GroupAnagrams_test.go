package main

import (
	"fmt"
	"testing"
)

func groupAnagrams(strs []string) (res [][]string) {
	mp := make(map[[26]int][]string)
	// 遍历字符串数组，将字母异位词放在map的同一个位置
	for _, str := range strs {
		var arr [26]int
		// 统计一个单词的字母个数
		for _, b := range str {
			arr[b-'a']++
		}
		// 通过统计相同的字母个数实现字母异位词的key
		mp[arr] = append(mp[arr], str)
	}
	// 将map中的字母异位词拼接到res二维数组中
	for _, str := range mp {
		res = append(res, str)
	}
	return
}

func TestGroupAnagrams(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
