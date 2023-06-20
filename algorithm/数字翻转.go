package main

import "math"

func reverse(x int) (rev int) {
	for x != 0 {
		if rev > math.MaxInt32/10 || rev < math.MinInt32/10 {
			return 0
		}
		yushu := x % 10
		x /= 10
		rev = rev*10 + yushu
	}
	return rev
}

func main() {
	println(reverse(123456))
}
