package main

import (
	"fmt"
)

type address struct {
	province string
	city     string
}

type person3 struct {
	name   string
	gender string
	add    address
}

type company struct {
	name    string
	address //这里使用的是匿名嵌套
}

func main() {
	p1 := person3{
		name:   "李沁",
		gender: "女",
		add: address{
			province: "山东",
			city:     "烟台",
		},
	}
	fmt.Println(p1, p1.name, p1.add.city) //这里不能直接使用p1.city

	c1 := company{
		name: "京",
		address: address{
			province: "北京",
			city:     "北京",
		},
	}
	fmt.Println(c1, c1.name, c1.city) //因为是匿名嵌套，这里可以直接使用c1.city
}
