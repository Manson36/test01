package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5
}
func f6() (x int) {
	defer func(x *int) {
		*x++
	}(&x)
	return 5
}

func main() {
	fmt.Println(f1()) //5 修改的x不是返回值
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5 //返回值是y，修改的是x
	fmt.Println(f4()) //5 改变的知识函数内x的副本
	fmt.Println(f5()) //5 defer中返回的值别没有接受者
	fmt.Println(f6()) //6
}
