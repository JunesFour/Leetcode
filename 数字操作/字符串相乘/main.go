package main

func multiply(num1 string, num2 string) string {
	a := make([]int, 110)
	b := make([]int, 110)
	c := make([]int, 110)

	len1 := len(num1)
	len2 := len(num2)
	for i := 0; i < len1; i++ {
		a[len1-i] = int(num1[i] - '0')
	}
	for i := 0; i < len2; i++ {
		b[len2-i] = int(num2[i] - '0')
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			c[i+j-1] += a[i] * b[j]
			if c[i+j-1] >= 10 {
				c[i+j] += c[i+j-1] / 10
				c[i+j-1] %= 10
			}
		}
	}
	var p int
	if len1 < len2 {
		p = len1
	} else {
		p = len2
	}
	for i := p; i <= len1+len2; i++ {
		if c[i] > 0 {
			p = i
		}
	}
	str := ""
	for i := p; i > 0; i-- {
		str += string(c[i] + '0')
	}
	return str
}

func main() {

}
