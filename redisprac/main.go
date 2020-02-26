package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
		return
	}
	defer conn.Close()

	var arr = []string{"name", "age", "skill"}
	var list [][]string

	_, err = conn.Do("hmset", "monster1", "name", "xiaoli", "age", "20", "skill", "pao")
	_, err = conn.Do("hmset", "monster2", "name", "xiaozhang", "age", "21", "skill", "跳")
	_, err = conn.Do("hmset", "monster3", "name", "xiaozhou", "age", "22", "skill", "飞")
	if err != nil {
		fmt.Println("hmset err:", err.Error())
		return
	}

	r1, err := redis.Strings(conn.Do("hmget", "monster1", "name", "age", "skill"))
	r2, err := redis.Strings(conn.Do("hmget", "monster2", "name", "age", "skill"))
	r3, err := redis.Strings(conn.Do("hmget", "monster3", "name", "age", "skill"))
	if err != nil {
		fmt.Println("hmget err:", err.Error())
	}

	list = append(list, r1)
	list = append(list, r2)
	list = append(list, r3)

	//for i, v := range r1 {
	//	fmt.Printf("monster1 %s:%s\n", arr[i], v)
	//}
	//for i, v := range r2 {
	//	fmt.Printf("monster1 %s:%s\n", arr[i], v)
	//}
	//for i, v := range r3 {
	//	fmt.Printf("monster1 %s:%s\n", arr[i], v)
	//}
	for i, v1 := range list {
		for j, v2 := range v1 {
			fmt.Printf("monster[%d] %s:%s\n", i+1, arr[j], v2)
		}
	}
}
