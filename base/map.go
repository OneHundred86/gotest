package base

import "fmt"

func TestMap(){
	var m1 map[string]int = make(map[string]int)
	m1["two"] = 2
	m1["ten"] = 10
	fmt.Println(m1)

	// 赋值，左边一个值时，取的是map的value
	v := m1["two"]
	fmt.Println(v)

	// 赋值，左边两个值时，第一个是map的value，第二个是代表map的值是否存在
	v1, ok1 := m1["one"]
	fmt.Println(v1, ok1)

	// 赋值初始化
	m2 := map[string]string{"one":"1", "two":"2"}
	m2["three"] = "3"
	fmt.Println(m2)
}