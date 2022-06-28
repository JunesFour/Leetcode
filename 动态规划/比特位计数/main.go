package main

func countBits(n int) (res []int) {
	res = make([]int, n+1)
	res[0] = 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			res[i] = res[i-1] + 1
		} else {
			res[i] = res[i/2]
		}
	}
	return
}

func main() {

}
