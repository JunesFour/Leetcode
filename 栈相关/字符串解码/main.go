package main

import "fmt"

func decodeString(s string) string {
	var res string
	multi := 0
	stackMulti := make([]int, 0)
	stackRes := make([]string, 0)
	for _, v := range s {
		if v == '[' {
			stackMulti = append(stackMulti, multi)
			stackRes = append(stackRes, res)
			multi = 0
			res = ""
		} else if v == ']' {
			var tmp string
			curMulti := stackMulti[len(stackMulti)-1]
			stackMulti = stackMulti[:len(stackMulti)-1]
			for i := 0; i < curMulti; i++ {
				tmp += res
			}
			res = stackRes[len(stackRes)-1] + tmp
			stackRes = stackRes[:len(stackRes)-1]
		} else if v >= '0' && v <= '9' {
			multi = multi*10 + int(v-'0')
		} else {
			res += string(v)
		}
	}
	return res
}

func main() {
	fmt.Println(decodeString("abc3[cd]xyz"))
}
