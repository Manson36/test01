package main

import (
	"encoding/json"
	"fmt"
)

type person4 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person4{
		Name: "李沁",
		Age:  29,
	}
	//序列化：把go语言的结构体变量 --> json结构的字符串
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("json marshal err:", err)
		return
	}
	fmt.Printf("%v\n", string(b))
	//反序列化：json格式的字符串 -->go语言能够识别的结构体变量
	str := `{"name":"李沁", "age":29}`
	var p2 person4
	json.Unmarshal([]byte(str), &p2) //传指针是为了能在json.Unmarshal内部修改p2的值
	fmt.Printf("%v\n", p2)
}
