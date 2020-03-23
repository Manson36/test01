package main

import (
	"fmt"
)

func main() {
	//元素类型为map的切片
	//var s1 = make([]map[int]string, 0, 10)
	//s1[0][0] = "李沁" //error:index out of range [0] 原因：make的切片的长度为0

	//var s2 = make([]map[int]string, 1, 10)
	//s2[0][0] = "白鹿" //panic:assignment to entry in nil map 原因：slice中的map没有初始化

	var s3 = make([]map[int]string, 10, 10)
	s3[0] = make(map[int]string, 10)
	s3[0][0] = "李沁"
	s3[0][1] = "白鹿"

	for i, v := range s3 {
		fmt.Println(i, v)
	}
}
