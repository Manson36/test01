package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//将map中的key进行排序，输出
func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%2d", i)
		value := rand.Intn(100) //生成0-99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片中
	var keys = make([]string, 0, 200)
	for key := range scoreMap { //map的key的遍历
		keys = append(keys, key)
	}

	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, k := range keys {
		fmt.Println(k, scoreMap[k])
	}
}
