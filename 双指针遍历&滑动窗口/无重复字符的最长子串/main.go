package main

import "math"

func lengthOfLongestSubstring(s string) (ret int) {
	l, n, cache := 0, len(s), make([]int, 128)
	for i := 0; i < n; i++ {
		idx := s[i]
		l = int(math.Max(float64(l), float64(cache[idx])))
		ret = int(math.Max(float64(ret), float64(i-l+1)))
		cache[idx] = i + 1
	}
	return
}

func main() {

}
