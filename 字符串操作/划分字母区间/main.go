package main

import "fmt"

func partitionLabels(s string) []int {
	hashTable := make(map[byte]int)
	res := make([]int, 0)
	for i, v := range s {
		hashTable[byte(v)] = i
	}
	index := 0
	pre := -1
	for i, v := range s {
		if index < hashTable[byte(v)] {
			index = hashTable[byte(v)]
		}
		if i == index {
			index = 0
			res = append(res, i-pre)
			pre = i
		}
	}
	return res
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
