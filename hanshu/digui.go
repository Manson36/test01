package main

import "fmt"

//计算n的阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}

	return n * f(n-1)
}

func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

func main() {
	//ret := f(7)
	//fmt.Println(ret)

	ret := taijie(5)
	fmt.Println(ret)
}
