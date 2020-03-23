package main

import (
	"fmt"
	"sync"
)

//没有加锁，所以存在数据竞争问题，得到的结果不是固定的
//var x = 0
//var wg sync.WaitGroup
//
//func add() {
//	for i := 0; i < 50000; i++ {
//		x = x + 1
//	}
//	wg.Done()
//}
//
//func main() {
//	wg.Add(3)
//	go add()
//	go add()
//	go add()
//	wg.Wait()
//	fmt.Println(x)
//}

//互斥锁，sunc.Mutex
var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}

}

func main() {
	wg.Add(3)
	go add()
	go add()
	go add()

	wg.Wait()
	fmt.Println(x)
}
