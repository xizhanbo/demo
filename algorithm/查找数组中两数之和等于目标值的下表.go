package main

import "fmt"

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

// 最优解法hashtable 时间复杂度O(n)

func search(arr []int, target int) []int {
	hashtable := map[int]int{}
	for i, x := range arr {
		if p, ok := hashtable[target-x]; ok {
			return []int{p, i}
		}
		hashtable[x] = i
	}
	return nil
}

func main() {
	arr := []int{2, 7, 17, 32, 33}
	target := 9
	result := search(arr, target)
	for _, index := range result {
		fmt.Println(index)
	}
}
