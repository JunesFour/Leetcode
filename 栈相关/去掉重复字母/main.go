package main

func removeDuplicateLetters(s string) string {
	var count [26]int
	var exist [26]bool
	var stack []rune
	for _, v := range s {
		count[v-'a']++
	}
	for _, v := range s {
		if exist[v-'a'] {
			count[v-'a']--
			continue
		}
		// 达到条件应该出栈
		// 1.栈里有元素
		// 2.栈顶元素大于当前元素
		// 3.栈顶元素在后续出现过
		for len(stack) > 0 && stack[len(stack)-1] > v && count[stack[len(stack)-1]-'a'] > 0 {
			exist[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
		count[v-'a']--
		exist[v-'a'] = true
	}
	return string(stack)
}

func main() {

}
