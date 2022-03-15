package main

func trap(height []int) (ans int) {
	left, right := 0, len(height)-1
	if right < 2 {
		return
	}
	for left < right {
		if height[left] < height[right] {
			if height[left] > height[left+1] {
				ans += height[left] - height[left+1]
				height[left+1] = height[left]
			}
			left++
		} else {
			if height[right] > height[right-1] {
				ans += height[right] - height[right-1]
				height[right-1] = height[right]
			}
			right--
		}
	}
	return
}

func main() {

}
