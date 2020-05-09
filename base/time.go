package base

import (
	"fmt"
	"time"
)

func TestTime() {
	fmt.Println(time.Now().Unix())

	fmt.Println(time.Now().UnixNano())
}
