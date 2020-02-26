package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //表示和数据库的最大链接数，0表示没有限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化链接代码，链接哪个ip
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
func main() {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("set", "cat1", "tom")
	if err != nil {
		fmt.Println("set err:", err.Error())
		return
	}
	r, err := redis.String(conn.Do("get", "cat1"))
	if err != nil {
		fmt.Println("get err:", err.Error())
		return
	}
	fmt.Println("r=", r)
}
