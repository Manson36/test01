package main

import "sync"

var (
	wg   sync.WaitGroup
	once sync.Once
)

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}

		ch2 <- x * x
	}

	//Once.Do()确保某个操作只执行一次
	//Once.Do()中的参数是没有参数的匿名函数，这里使用闭包的方法传参ch2进去
	f := func() { close(ch2) }
	once.Do(f)
}

func main() {

}
