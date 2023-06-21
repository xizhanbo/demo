package main

import "fmt"

// O(n)
func findPeakElements(nums []int) []int {
	n := len(nums)
	if n == 1 { // 特殊情况：数组中只有一个元素
		return []int{0}
	}
	peaks := []int{}
	if nums[0] > nums[1] { // 处理数组的左侧边界
		peaks = append(peaks, 0)
	}
	for i := 1; i < n-1; i++ { // 处理数组中间的元素
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			peaks = append(peaks, i)
		}
	}
	if nums[n-1] > nums[n-2] { // 处理数组的右侧边界
		peaks = append(peaks, n-1)
	}
	return peaks
}

func main() {
	arr := []int{1, 3, 5, 2, 7, 10, 9}
	fmt.Println(findPeakElements(arr))
}
