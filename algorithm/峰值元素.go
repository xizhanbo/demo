package main

import "fmt"

func findPeakElement(nums []int) int {
	return findPeakElementHelper(nums, 0, len(nums)-1)
}

// 二分法 log2N
func findPeakElementHelper(nums []int, left, right int) int {
	if left == right { // 如果数组中只有一个元素，则它就是峰值元素
		return left
	}
	mid := (left + right) / 2
	if nums[mid] > nums[mid+1] { // 说明峰值元素可能在左侧，递归查找左侧数组
		return findPeakElementHelper(nums, left, mid)
	}
	// 说明峰值元素可能在右侧，或者当前元素就是峰值元素，递归查找右侧数组
	return findPeakElementHelper(nums, mid+1, right)
}

func main() {
	arr := []int{1, 3, 5, 2, 7, 10, 9}
	fmt.Println(findPeakElement(arr))
}
