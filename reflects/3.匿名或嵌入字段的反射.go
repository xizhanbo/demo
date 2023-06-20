package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Id   int
	Name string
}

type People struct {
	Student // 匿名字段
}

func main() {
	p := People{Student{Id: 1, Name: "咖啡色的羊驼"}}

	t := reflect.TypeOf(p)
	// 这里需要加一个#号，可以把struct的详情都给打印出来
	// 会发现有Anonymous:true，说明是匿名字段
	fmt.Printf("%#v\n", t.Field(0))

	// 取出这个学生的名字的详情打印出来，数组中0 代表最外层的第一个结构体，1代表外出第一个结构体的第二个参数（位置代表了哪一层，数值代表的第几个）
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

	// 获取匿名字段的值的详情
	v := reflect.ValueOf(p)
	fmt.Printf("%#v\n", v.Field(0))
}
