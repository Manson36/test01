package main

import "fmt"

func sum(param []int) int64 {
	var sum int64
	for _, v := range param {
		sum += int64(v)
	}

	return sum
}

func main() {
	var a = []int{1, 2, 3}
	value := sum(a)
	fmt.Println(value)
}

//问题：
//1.参数名称起的太随意，比如a
//2.既然是求和，要有返回值
//3.任意数量求和，参数使用切片就好，不要想着使用...
//4.int类型的求和，如果在最大范围的值，会有溢出问题，要改为int64类型求和
