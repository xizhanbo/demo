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
	s := Student{Id: 1, Name: "咖啡色的羊驼"}
	t := reflect.TypeOf(s)

	// 通过.Kind()来判断对比的值是否是struct类型
	if k := t.Kind(); k == reflect.Struct {
		fmt.Println("bingo")
	}

	num := 1
	numType := reflect.TypeOf(num)
	if k := numType.Kind(); k == reflect.Int {
		fmt.Println("bingo")
	}
}
