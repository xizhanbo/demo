package main

// N*log2N
func QuickSort(array []int) []int {
	// 基线条件：数组为空或仅有一个元素时直接返回
	if len(array) < 2 {
		return array
	}
	// 选择第一个元素作为主元素
	pivot := array[0]
	var (
		less    []int
		greater []int
		equal   []int
	)
	// 分区操作，将数组分为小于、等于、大于主元素的三部分
	for _, num := range array {
		switch {
		case num < pivot:
			less = append(less, num)
		case num > pivot:
			greater = append(greater, num)
		default:
			equal = append(equal, num)
		}
	}
	// 递归排序
	less, greater = QuickSort(less), QuickSort(greater)
	// 将排好序的三部分组合起来
	result := append(less, equal...)
	result = append(result, greater...)
	return result
}
