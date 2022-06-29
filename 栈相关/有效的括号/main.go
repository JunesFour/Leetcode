package main

func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		}
		if len(stack) == 0 {
			return false
		}
		if s[i] == ')' {
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		if s[i] == ']' {
			if stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		if s[i] == '}' {
			if stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {

}
