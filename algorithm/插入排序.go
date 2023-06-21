package main

// O(N平方）
func insertSort(arr *[]int) error {
	length := len(*arr)
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0 && (*arr)[j] > (*arr)[j+1]; j-- {
			(*arr)[j+1], (*arr)[j] = (*arr)[j], (*arr)[j+1]
		}
	}
	return nil
}

func main() {
	arr := []int{6, 9, 3, 2, 1, 5, 7}
	err := insertSort(&arr)
	if err != nil {
		panic(err)
	}
	for i := 1; i < len(arr); i++ {
		println(arr[i])
	}
}
