package main

import "fmt"

/*
示例:
输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
说明:动态规划方法。O(n)
*/
func jump(nums []int) int {
	length := len(nums)
	end := 0         //在这段代码中，end代表上一次可以跳跃到索引位置。
	maxPosition := 0 //每次跳跃的最大步长
	steps := 0
	for i := 0; i < length-1; i++ {
		maxPosition = getMax3(maxPosition, i+nums[i])
		if i == end { //如果当前位置i等于end时，说明跳跃已经到了最远索引位置，需要再次跳跃，并将end更新为这次跳跃的最远距离。同时，记录跳跃次数steps。
			end = maxPosition
			steps++
		}
	}
	return steps
}
func getMax3(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func main() {
	arr := []int{2, 3, 1, 1, 4}
	step := jump(arr)
	fmt.Println(step)
}
