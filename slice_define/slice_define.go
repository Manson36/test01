package main

import "fmt"

func main() {
	var s1 []int //没有分配内存
	fmt.Println(s1)
	fmt.Println(s1 == nil)

	s1 = []int{1, 2, 3}
	fmt.Println(s1)

	s2 := make([]int, 2, 4)
	fmt.Println(s2)
	s3 := make([]int, 0, 4) //make初始化，分配内存
	fmt.Println(s3)
	fmt.Println(s3 == nil)
}
