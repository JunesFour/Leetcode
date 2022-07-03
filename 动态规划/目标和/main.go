package main

// 减法对应总和：neg
// 加法对应总和：sum - neg
// target = (sum - neg) - neg
// neg = (sum - target) / 2
// 问题就转化为01背包，nums中的数字之和凑成(sum - target) / 2有几种方法
// dp[i]表示凑成i有dp[i]种方法
// dp[j] += dp[j - nums[i]]
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for j := neg; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[neg]
}

func main() {

}
