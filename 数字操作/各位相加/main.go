package main

func addDigits(num int) int {
	for num >= 10 {
		num = num%10 + num/10
	}
	return num
}

func main() {

}
