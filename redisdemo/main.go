package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	//1.连接redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("conn redis err, errmsg:", err.Error())
		return
	}
	//fmt.Println("connection success",conn)
	//2.通过go向redis写入数据 string [key-value]
	_, err = conn.Do("set", "name1", "tom") //Do传入参数是cmd，args...,返回interface{}, err
	_, err = conn.Do("mset", "name2", "john", "name3", "marry")
	_, err = conn.Do("expire", "name1", 10) //设置过期时间：10秒
	if err != nil {
		fmt.Println("set err, errmsg:", err.Error())
		return
	}
	defer conn.Close() //关闭。。
	//3.通过go从redis读取数据
	r1, err := redis.String(conn.Do("get", "name1"))
	r2, err := redis.Strings(conn.Do("mget", "name2", "name3")) //返回多个结果，必须使用Strings
	if err != nil {
		fmt.Println("get err, err msg:", err.Error())
		return
	}
	for i, v := range r2 {
		fmt.Printf("r2[%d]:%s\n", i, v)
	}
	fmt.Println("get读取的结果是：", r1, r2)

	//useHash(conn)
	useList(conn)
}

//操作Hash结构
func useHash(conn redis.Conn) {
	//写入数据
	_, err := conn.Do("Hset", "messages", "name", "xiaozhou")
	_, err = conn.Do("Hset", "messages", "age", "20")
	_, err = conn.Do("Hmset", "user01", "age", "20", "name", "tom")
	if err != nil {
		println("err msg:", err.Error())
	}

	//读取数据
	r1, err := redis.String(conn.Do("hget", "messages", "name"))
	r2, err := redis.Int(conn.Do("hget", "messages", "age"))
	//r2, err := redis.String(conn.Do("hget","messages", "age")) 使用String得到的结果也是对的

	r3, err := redis.Strings(conn.Do("hmget", "user01", "name", "age"))
	if err != nil {
		fmt.Println("hash get err:", err.Error())
	}

	fmt.Printf("hget name=%s, age=%v, user01:%v", r1, r2, r3)
	for _, v := range r3 {
		fmt.Println(v)
	}
}

//操作List数据
func useList(conn redis.Conn) {
	//写入数据
	_, err := conn.Do("lpush", "herolist", "lujunyi", "wusong")
	if err != nil {
		fmt.Println("list push err:", err.Error())
		return
	}
	//读取数据
	r, err := redis.String(conn.Do("rpop", "herolist"))
	if err != nil {
		fmt.Println("list pop err:", err.Error())
		return
	}

	fmt.Println("list 结果:", r)
}
