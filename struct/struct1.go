package main

import "fmt"

type person1 struct {
	name, gender string
}

func f(x person1) {
	x.gender = "女"
}
func f2(x *person1) {
	x.gender = "女" //语法糖，自动根据指针找对应的变量，结构体中出现
}

func main() {
	var p person1
	p.name = "张三"
	p.gender = "男"
	f(p)                  //因为是值类型，所以函数只是拷贝了 p 的一个副本
	fmt.Println(p.gender) //这里输出的还是 男。
	f2(&p)
	fmt.Println(p.gender)
}
