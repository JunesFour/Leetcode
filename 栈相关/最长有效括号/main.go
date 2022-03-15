package main

// 栈内存放括号的位置
// 栈基存放的是最后一个多余右括号的位置，方便计数，初始时压入-1
// 遇到右括号时必须先弹出一个值
// 如果弹出后栈内元素为空，那么更新最后一个多余右括号的位置
// 如果不为空，那就更新最长有效括号的长度，方法是当前下标减去栈顶元素的下标
func longestValidParentheses(s string) int {
	maxLength := 0
	stack := []int{-1}
	for i := 0; i < len(s); i++ {
		// 如果是左括号，直接将位置入栈
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			// 不管怎么样，先弹出一个元素
			stack = stack[:len(stack)-1]
			// 弹出的是栈基元素，那么更新最后一个多余右括号的下标
			if len(stack) == 0 {
				stack = append(stack, i)
				// 弹出的不是栈基，更新最大值
			} else {
				if maxLength < i-stack[len(stack)-1] {
					maxLength = i - stack[len(stack)-1]
				}
			}
		}
	}
	return maxLength
}

func main() {

}
