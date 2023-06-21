/*
中位数是有序序列最中间的那个数。如果序列的大小是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。
例如：
[2,3,4]，中位数是 3
[2,3]，中位数是 (2 + 3) / 2 = 2.5
给你一个数组 nums，有一个大小为 k 的窗口从最左端滑动到最右端。窗口中有 k 个数，每次窗口向右移动 1 位。你的任务是找出每次窗口移动后得到的新窗口中元素的中位数，并输出由它们组成的数组。

示例：

给出 nums = [1,3,-1,-3,5,3,6,7]，以及 k = 3。

窗口位置                      中位数
---------------               -----
[1  3  -1] -3  5  3  6  7       1
1 [3  -1  -3] 5  3  6  7      -1
1  3 [-1  -3  5] 3  6  7      -1
1  3  -1 [-3  5  3] 6  7       3
1  3  -1  -3 [5  3  6] 7       5
1  3  -1  -3  5 [3  6  7]      6
因此，返回该滑动窗口的中位数数组 [1,-1,-1,3,5,6]。
该函数的时间复杂度为O(nklogk)，其中n 表示数组元素个数，k 表示窗口大小
————————————————
*/
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func window(num []int) float64 {
	sort.Ints(num)
	len := len(num)
	if len%2 == 1 {
		return float64(num[(len)/2])
	} else {
		result, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", float64(num[len/2]+num[len/2-1])/float64(2)), 64)
		return result
	}
}
func main() {
	nums := [...]int{1, 3, -1, -3, 5, 3, 6, 7}
	wLen := 4 //窗口大小
	result := make([]float64, len(nums)-wLen+1)
	i := 0
	for slow := 0; slow <= len(nums)-wLen; slow++ {
		fast := slow + wLen
		win := make([]int, wLen)
		copy(win[:], nums[slow:fast])
		//这里只能是拷贝，不可以用下面这种
		//win := nums[slow:fast]
		fmt.Println(win)
		result[i] = window(win)
		i++
	}
	fmt.Println(result)
}

/*
copy(win[:], nums[slow:fast]) 和 win := nums[slow:fast] 都是将数组 nums 中下标从 slow 开始长度为 wLen 的元素截取出来，生成一个新的数组。但是，它们的实现机制和结果不同。

win := nums[slow:fast] 是使用切片语法对 nums 中下标从 slow 开始长度为 wLen 的元素进行截取，返回一个新的切片（底层引用的是 nums 数组的一部分）。这种方式直接生成了一个切片，因此在使用时很方便，但是需要注意的是，它们引用的是同一个底层数组，如果在新的切片上进行修改，会影响到原数组。造成数据的混乱

对于 copy(win[:], nums[slow:fast])，其过程是先创建了一个长度为 wLen 的新数组 win，然后使用copy()将 nums[slow:fast] 中的元素复制到 win 数组中。这种方式生成的数组与原数组 nums 无关，它们所占用的内存空间是不同的，而且可以避免因为修改 win 数组而影响到原数组 nums。copy() 函数需要指定目标数组和要复制的源数组，并返回实际复制的元素个数。

综上，二者的区别在于 win := nums[slow:fast] 直接生成了一个切片，引用的是原数组 nums 的一部分，而 copy(win[:], nums[slow:fast]) 需要新建一个数组来存储复制的内容。
*/
