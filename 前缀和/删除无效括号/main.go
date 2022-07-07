package main

func removeInvalidParentheses(s string) []string {
	lremove, rremove := 0, 0
	for _, i := range s {
		if i == '(' {
			lremove += 1
		} else if i == ')' && lremove == 0 {
			rremove += 1
		} else if i == ')' && lremove > 0 {
			lremove -= 1
		}
	}
	ans := []string{}
	dfs(&ans, s, 0, 0, 0, lremove, rremove)
	return ans
}

func dfs(ans *[]string, s string, lcount, rcount, start, lremove, rremove int) {
	if lremove == 0 && rremove == 0 {
		if valid(s) {
			*ans = append(*ans, s)
		}
		return
	}
	for i := start; i < len(s); i++ {
		if s[i] == '(' {
			lcount += 1
		}
		if s[i] == ')' {
			rcount += 1
		}
		if i > start && s[i] == s[i-1] {
			continue
		}
		if s[i] == '(' && lremove > 0 {
			dfs(ans, s[:i]+s[i+1:], lcount-1, rcount, i, lremove-1, rremove)
		}
		if s[i] == ')' && rremove > 0 {
			dfs(ans, s[:i]+s[i+1:], lcount, rcount-1, i, lremove, rremove-1)
		}
		//剪枝优化
		if rcount > lcount {
			break
		}
	}
}

func valid(s string) bool {
	cnt := 0
	for _, v := range s {
		if v == '(' {
			cnt++
		} else if v == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

func main() {

}
