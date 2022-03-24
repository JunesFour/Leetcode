package main

// 完全背包
func wordBreak(s string, wordDict []string) bool {
	hashTable := make(map[string]bool)
	for _, v := range wordDict {
		hashTable[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && hashTable[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func main() {

}
