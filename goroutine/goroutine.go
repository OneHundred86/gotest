package goroutine

import (
	"fmt"
	"sync"
)

//
var wg sync.WaitGroup

func printOddNum(ch chan int) {
	for {
		n, ok := <-ch

		// chan关闭，ok为false
		if !ok {
			break
		}
		fmt.Println("奇数:", n)
	}
	// 计数器减1
	wg.Done()
}

func printEvenNum(ch chan int) {
	// chan关闭，循环会停止
	for n := range ch {
		fmt.Println("偶数:", n)
	}
	// 计数器减1
	wg.Done()
}

func Test() {
	oddCh := make(chan int)
	evenCh := make(chan int)
	go printOddNum(oddCh)
	go printEvenNum(evenCh)
	// 计数器为2
	wg.Add(2)

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
	close(oddCh)
	close(evenCh)

	// 阻塞直到计数器为0
	wg.Wait()
}
