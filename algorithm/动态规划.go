// https://blog.csdn.net/baidu_28312631/article/details/47418773
package main

import (
	"fmt"
)

const MAX = 101

// O(n)
func main() {
	var D [MAX][MAX]int
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Scan(&D[i][j])
		}
	}
	maxSum := D[n][:]
	for i := n - 1; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			maxSum[j] = getMax(maxSum[j], maxSum[j+1]) + D[i][j]
		}
	}
	fmt.Println(maxSum[1])
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
