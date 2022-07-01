package main

import "fmt"

func compare(mp1, mp2 map[rune]int) bool {
	for k, v := range mp1 {
		if v != mp2[k] {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) (res []int) {
	if len(p) > len(s) {
		return
	}
	// 滑动窗口的大小
	n := len(p)
	// 滑动窗口中每种元素的数量
	targetMap := make(map[rune]int)
	for _, v := range p {
		targetMap[v]++
	}
	// 构建滑动窗口
	window := make(map[rune]int)
	for i := 0; i < n; i++ {
		window[rune(s[i])]++
	}
	if compare(targetMap, window) {
		res = append(res, 0)
	}
	for i := n; i < len(s); i++ {
		window[rune(s[i])]++
		window[rune(s[i-n])]--
		if compare(targetMap, window) {
			res = append(res, i-n+1)
		}
	}
	return
}

func findAnagrams1(s string, p string) (ans []int) {
	if len(s) < len(p) {
		return nil
	}
	cntP, cntS := [26]int{}, [26]int{}
	for i := range p {
		cntP[p[i]-'a']++
		cntS[s[i]-'a']++
	}
	if cntP == cntS {
		ans = append(ans, 0)
	}
	for i := len(p); i < len(s); i++ {
		cntS[s[i]-'a']++
		cntS[s[i-len(p)]-'a']--
		if cntS == cntP {
			ans = append(ans, i-len(p)+1)
		}
	}
	return ans
}

func main() {
	//mp1 := make(map[rune]int)
	//mp2 := make(map[rune]int)
	//mp1['a'] = 1
	//mp1['c'] = 2
	//mp2['a'] = 1
	//mp2['c'] = 1
	//fmt.Println(compare(mp1, mp2))
	fmt.Println(findAnagrams("abab", "ab"))
}
