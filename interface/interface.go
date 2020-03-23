package main

import "fmt"

//类型断言1
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("他不是字符串")
	} else {
		fmt.Println("传进来的字符串是：", str)
	}
}

//类型断言2
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch v := a.(type) {
	case string:
		fmt.Println("是一个字符串：", v)
	case int:
		fmt.Println("是一个int", v)
	case bool:
		fmt.Println("是一个bool类型", v)
	}
}

func main() {
	assign(100)
	assign2(false)
}
