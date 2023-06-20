package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "雪白的羊驼"
	t := reflect.TypeOf(name)
	v := reflect.ValueOf(name)
	fmt.Println("type:", t)
	fmt.Println("value:", v)
}
