// 示例 1：

// 	输入：nums1 = [1,3], nums2 = [2]
// 	输出：2.00000
// 	解释：合并数组 = [1,2,3] ，中位数 2
// 示例 2：

// 	输入：nums1 = [1,2], nums2 = [3,4]
// 	输出：2.50000
// 	解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
package main

func find(num1, num2 []int) float64 {
	l1 := len(num1)
	l2 := len(num2)
	sum := l1 + l2
	if sum%2 == 1 { //奇数去中间
		return float64(getNumByIndex(num1, num2, sum/2))
	} else {

		return (float64(getNumByIndex(num1, num2, sum/2-1)) + float64(getNumByIndex(num1, num2, sum/2))) / 2.0
	}
}

func getNumByIndex(x, y []int, index int) int {
	i, x1, y1 := 0, 0, 0
	res := 0
	l1 := len(x)
	l2 := len(y)
	for i <= index {
		if l1-1 < x1 {
			res = y[y1]
			y1++
			i++
			continue
		}
		if l2-1 < y1 {
			res = x[x1]
			x1++
			i++
			continue
		}
		if x[x1] < y[y1] {
			res = x[x1]
			x1++
		} else {
			res = y[y1]
			y1++
		}
		i++

	}
	return res
}

func main() {
	x := []int{1, 3}
	y := []int{2}
	println(find(x, y))
	// x2 := []int{1, 2}
	// y2 := []int{3, 4}
	// println(find(x2, y2))
}
