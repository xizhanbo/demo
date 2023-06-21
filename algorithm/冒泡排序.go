package main

func Sort(arr *[]int) error {
	length := len(*arr)
	for i := 0; i < length-1; i++ { //控制魂环轮数
		for j := 0; j < length-i-1; j++ { //每轮两两比较
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j+1], (*arr)[j] = (*arr)[j], (*arr)[j+1]
			}
		}
	}
	return nil
}

func main() {
	arr := []int{6, 9, 3, 2, 1, 5, 7}
	err := Sort(&arr)
	if err != nil {
		panic(err)
	}
	for i := 1; i < len(arr); i++ {
		println(arr[i])
	}
}

//时间复杂度O(n平方)
