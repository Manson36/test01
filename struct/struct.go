package main

import "fmt"

type myInt int
type yourInt = int

func main() {
	var m myInt
	m = 100
	fmt.Println(m)
	fmt.Printf("%T %d\n", m, m)
	var y yourInt
	y = 100
	fmt.Printf("m的类型：%T；y的类型：%T", m, y)
}
