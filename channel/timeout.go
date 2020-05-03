package channel

import (
	"fmt"
	"time"
)

func sendNum(from int, to int) (ch chan int) {
	ch = make(chan int)

	go func() {
		for i := from; i <= to; i++ {
			ch <- i
			if i%2 == 0 {
				fmt.Println("sleeping 2 seconds...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	return
}

func Test() {
	ch := sendNum(10, 20)

	for {
		select {
		case v := <-ch:
			fmt.Println("Received:", v)
		case <-time.After(time.Second):
			fmt.Println("timeout!")
		}
	}
}
