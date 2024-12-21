package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//chanDemo()
	raceDemo()
}

// channel用于在多个goroutine之间传递数据
func chanDemo() {
	ch := make(chan int)
	defer close(ch)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for n := range ch {
			fmt.Println(n)
		}
	}()

	ch2 := make(chan string)
	defer close(ch2)
	go func() {
		ch2 <- "hello"
	}()
	go func() {
		str := <-ch2
		fmt.Println(str)
	}()

	//select case 通常在多个channel选择
	ch3 := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			ch3 <- "hello" + fmt.Sprintf("%d", i)
		}
	}()
	go func() {
		for {
			select {
			case str := <-ch3:
				fmt.Println("ch3", str)
			case str := <-ch2:
				fmt.Println("ch2", str)

			default:
				//fmt.Println("default")
			}
		}
	}()
	time.Sleep(time.Second * 11)
}

// 多线（协）程竞争变量安全 锁和同步器演示
func raceDemo() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	num := 1

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			num++
		}()
	}
	wg.Wait()
	fmt.Println(num)
}
