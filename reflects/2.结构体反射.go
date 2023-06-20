package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Id   int
	Name string
}

func (s student) Hello() {
	fmt.Println("hello! i am %v ,a student", s.Name)
}

func main() {
	s := student{
		Id:   1,
		Name: "羊驼",
	}
	// 获取目标对象
	t := reflect.TypeOf(s)
	// .Name()可以获取去这个类型的名称
	fmt.Println("这个类型的名称是:", t.Name())
	// 获取目标对象的值类型
	v := reflect.ValueOf(s)
	// .NumField()来获取其包含的字段的总数
	for i := 0; i < t.NumField(); i++ {
		// 从0开始获取Student所包含的key
		key := t.Field(i)
		// 通过interface方法来获取key所对应的值
		value := v.Field(i).Interface()
		fmt.Printf("第%d个字段是：%s:%v = %v \n", i+1, key.Name, key.Type, value)
	}
	// 通过.NumMethod()来获取Student里头的方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("第%d个方法是：%s:%v\n", i+1, m.Name, m.Type)
	}
}
