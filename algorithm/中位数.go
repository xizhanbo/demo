// 示例 1：

// 	输入：nums1 = [1,3], nums2 = [2]
// 	输出：2.00000
// 	解释：合并数组 = [1,2,3] ，中位数 2
// 示例 2：

// 输入：nums1 = [1,2], nums2 = [3,4]
// 输出：2.50000
// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
package main

import (
	"fmt"
	"strconv"
)

// Comparable 代表可以进行比较的int和float类型
type Comparable[T any] interface {
	int | float64
}

func find(num1, num2 []int) interface{} {
	l1 := len(num1)
	l2 := len(num2)
	sum := l1 + l2

	if sum%2 == 1 { //奇数取中间
		result := getNumByIndex(num1, num2, sum/2)
		return result
	} else {
		add := float64(getNumByIndex(num1, num2, sum/2-1)) + float64(getNumByIndex(num1, num2, sum/2))
		result := fmt.Sprintf("%.1f", add/2.0)
		res, _ := strconv.ParseFloat(result, 64)
		return res
	}
}

// 该方法默认是升序排列了
func getNumByIndex(x, y []int, index int) int {
	i, x1, y1 := 0, 0, 0 //i 代表循环次数，x1/y1分别代表两个数组的当前索引位置
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
	y := []int{2, 5}
	res := find(x, y)
	fmt.Println(res)
	//v := reflect.ValueOf(res)
	//switch v.Kind() {
	//case reflect.Int:
	//	i := v.Int()
	//	fmt.Println(i)
	//case reflect.Float64:
	//	f := v.Float()
	//	fmt.Printf("%.1f\n", f)
	//}
}

//时间复杂度为O(m+n)
