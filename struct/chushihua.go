package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	//方法一：直接赋值
	var p1 person
	p1.name = "李沁"
	p1.age = 29
	fmt.Println(p1)
	var p2 = person{
		name: "白鹿",
		age:  25,
	}
	fmt.Println(p2)
	var p3 = person{
		"小紫",
		26,
	}
	fmt.Println(p3)
}
