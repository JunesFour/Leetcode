package main

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i, _ := range dp {
		dp[i] = make([]bool, n)
	}
	ans := ""
	// 第一层遍历子串的长度
	for l := 0; l < n; l++ {
		// 第二层是子串的左边界
		for i := 0; i+l < n; i++ {
			// j是子串的右边界
			j := i + l
			if l == 0 {
				dp[i][j] = true
			} else if l == 1 {
				dp[i][j] = (s[i] == s[j])
			} else {
				dp[i][j] = (s[i] == s[j] && dp[i+1][j-1])
			}
			// 更新最大值
			if dp[i][j] && l+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}

func main() {

}
