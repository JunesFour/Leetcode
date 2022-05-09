package main

func moveZeroes(nums []int) {
	zeroIndex := -1
	for i := 0; i < len(nums); i++ {
		// 第一次遍历到0
		if zeroIndex == -1 && nums[i] == 0 {
			zeroIndex = i
			continue
		}
		// 不是第一次遍历到0
		if zeroIndex != -1 && nums[i] == 0 {
			continue
		}
		// 遍历到数字但是zeroIndex等于-1
		if nums[i] != 0 && zeroIndex == -1 {
			continue
		}
		if nums[i] != 0 && zeroIndex != -1 && zeroIndex < i {
			nums[i], nums[zeroIndex] = nums[zeroIndex], nums[i]
			for j := zeroIndex + 1; j < len(nums); j++ {
				if nums[j] == 0 {
					zeroIndex = j
					break
				}
			}
		}
	}
}

func moveZeroes1(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func main() {

}
