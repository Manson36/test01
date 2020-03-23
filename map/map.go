package main

import (
	"fmt"
)

func main() {
	var m map[string]int
	fmt.Println(m == nil)        //map只定义没有初始化，在内存中没有分配空间，为nil
	m = make(map[string]int, 10) //要估算好该map的容量，避免在运行期间进行动态扩容
	m["李沁"] = 90                 //使用key-value的方式去访问
	m["白鹿"] = 94
	m["杨紫"] = 92
	fmt.Println(m)

	value, ok := m["小紫"] //获取map的值
	if !ok {
		fmt.Println("查无此人")
	} else {
		fmt.Println(value)
	}

	for k, v := range m { //map的遍历
		fmt.Println(k, v)
	}

	delete(m, "杨紫") //map的删除
	fmt.Println(m)
}
