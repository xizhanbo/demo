package main

//dlv 测试用例
//dlv debug 文件名 ： 开启debug
//开启之后
//b 文件名:行号， 定位行号
//n 逐行执行    步过，遇到函数不进入函数
//s 是执行step 当遇到函数时，进入函数
//p 是打印
//c 执行到最后或者下一个断点
//q 退出
import "fmt"

func add(c *int) *int {
	*c++
	return c
}

func main() {
	a := new(int)
	*a = 10
	b := add(a)
	fmt.Println(*b)
	b = add(a)
	fmt.Println(*b)
	b = add(a)
	fmt.Println(*b)
	b = add(a)
	fmt.Println(*b)
	b = add(a)
	fmt.Println(*b)
}
