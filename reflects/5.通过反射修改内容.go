package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	s := &Student{Id: 1, Name: "咖啡色的羊驼"}

	v := reflect.ValueOf(s)

	// 修改值必须是指针类型否则不可行
	if v.Kind() != reflect.Ptr {
		fmt.Println("不是指针类型，没法进行修改操作")
		return
	}

	// 获取指针所指向的元素
	v = v.Elem()

	// 获取目标key的Value的封装
	name := v.FieldByName("Name")

	if name.Kind() == reflect.String {
		name.SetString("小学生")
	}

	fmt.Printf("%#v \n", *s)

	// 如果是整型的话
	test := 888
	testV := reflect.ValueOf(&test)
	testV.Elem().SetInt(666)
	fmt.Println(test)
}
