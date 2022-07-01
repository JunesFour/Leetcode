package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 统计从当前节点出发，满足的路径数目
func pathFrom(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	// 路径数量
	cnt := 0
	if root.Val == sum {
		cnt++
	}
	// 如果已经找到一条满足的路径，后面可以能会出现正负值相互抵消的情况，从而新增一条更长的路径
	cnt += pathFrom(root.Left, sum-root.Val)
	cnt += pathFrom(root.Right, sum-root.Val)
	return cnt
}
func pathSum(root *TreeNode, sum int) (res int) {
	if root == nil {
		return 0
	}
	return pathFrom(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func pathSum1(root *TreeNode, targetSum int) (res int) {
	// 前缀和表
	mp := make(map[int]int)
	var dfs func(root *TreeNode)
	pre := 0
	// 初始状态
	mp[0] = 1
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		pre += root.Val
		// 累加上根节点到当前节点之间刚好和等于目标值的路径
		res += mp[pre-targetSum]
		// 记录前缀和
		mp[pre] += 1
		dfs(root.Left)
		dfs(root.Right)
		// 因为是树形结构，所以要回溯
		mp[pre] -= 1
		pre -= root.Val
	}
	dfs(root)
	return
}

func main() {
	//pathSum()
}
