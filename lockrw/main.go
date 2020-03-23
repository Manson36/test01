package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      = 0
	wg     = sync.WaitGroup{}
	lock   = sync.Mutex{}
	rwlock = sync.RWMutex{}
)

func read() {
	defer wg.Done()
	rwlock.RLock()
	fmt.Println("x 的值是：", x)
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
}

func write() {
	defer wg.Done()
	rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 20)
	rwlock.Unlock()
}

func main() {
	start := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	//time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}

//第一点：sync.WaitGroup的作用：没有sync.WaitGroup，没有输出x的值。因为在主程序结束的时候，程序就结束了
//第二点：读写互斥锁。尤其适用于读的操作远远大于写的时候，会大大节省时间。
//第三点：使用读写互斥锁会发现读取的速度太快，输出的全是1，可以加time.sleep暂时解决
