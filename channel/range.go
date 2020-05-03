package channel

import (
	"fmt"
)

func TestRange() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	close(ch)
	// 读取完管道里面的值后，第二个值会收到false
	v, ok := <-ch
	fmt.Println(v, ok) // 10 true
	v1, ok1 := <-ch
	fmt.Println(v1, ok1) // 20 true
	v2, ok2 := <-ch
	fmt.Println(v2, ok2) // 0 false
	v3, ok3 := <-ch
	fmt.Println(v3, ok3) // 0 false

	ch1 := make(chan string, 5)
	ch1 <- "hello"
	ch1 <- "world"
	close(ch1)
	// 读取完管道里面的值后，会结束循环
	for v := range ch1 {
		fmt.Println("received:", v)
	}
	fmt.Println("end")
}

/*
10 true
20 true
0 false
0 false
received: hello
received: world
end
*/
