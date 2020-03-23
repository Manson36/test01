package main

import "fmt"

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//要求：f1(f2)
//分析：f1的参数是没有参数的函数，而f2有两个参数；这时就需要创建一个函数返回func(),但在函数内部通 过闭包 执行f2的操作
//进一步分析：我们要想把f2传进函数，那么就需要将f2作为参数传入f3
func f3(f func(int, int), x, y int) func() {
	return func() {
		//fmt.Println(x + y)
		f(x, y)
	}
}

func main() {
	//f1(f3(11, 22))
	f1(f3(f2, 11, 22))
}
