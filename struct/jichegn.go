package main

import (
	"fmt"
)

type animal struct {
	name string
}

//给animal实现了一个move的方法：
func (a animal) move() {
	fmt.Println(a.name, "会移动。")
}

type dog struct {
	feet int8
	animal
}

//给dog实现了一个wang的方法
func (d dog) wang() {
	fmt.Println(d.name, "在汪汪！")
}

func main() {
	d1 := dog{feet: 4, animal: animal{name: "周林"}}
	fmt.Println(d1)
	d1.move()
	d1.wang()
}
