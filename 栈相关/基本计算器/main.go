package main

func calculate(s string) int {
	stack, res, num, sign := []int{}, 0, 0, 1
	for i := 0; i < len(stack); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		} else if s[i] == '+' {
			res += sign * num
			num = 0
			sign = 1
		} else if s[i] == '-' {
			res += sign * num
			num = 0
			sign = -1
		} else if s[i] == '(' {
			// 将前一个结果和符号压入栈中
			stack = append(stack, res)
			stack = append(stack, sign)
			// 将结果置为0， 只需要在括号内计算新结果即可
			res = 0
			sign = 1
		} else if s[i] == ')' {
			res += sign * num
			num = 0
			res *= stack[len(stack)-1]
			res += stack[len(stack)-2]
			stack = stack[:len(stack)-2]
		}
	}
	// 最后不是以右括号结尾就是以一个数字结尾
	if num != 0 {
		res += sign * num
	}
	return res
}

func main() {

}
