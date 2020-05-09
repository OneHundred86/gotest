package base

import "fmt"

type Human struct {
	name string
	sex  int
	age  int
}

type Worker struct {
	career string
	Human  // 这里的Human也是一种类型，此时它相当于一种匿名字段
}

var (
	W Worker
)

func init() {
	W.name = "fromInit"
	W.sex = 1
}

func TestStruct() {
	fmt.Printf("struct.go: %+v %v \n", W, W.name) //小写字段，本包可以通过W.name使用，其他包不可以

	w1 := Worker{career: "worker", Human: Human{name: "worker1", sex: 1, age: 21}}
	fmt.Println(w1.age, w1.Human.age) // 两种写法都可以

	w2 := Worker{}
	w2.name = "worker2"
	w2.career = "worker"
	w2.age = 22
	fmt.Println(w2) // {worker {worker2 0 22}}

	var w3 Worker = w2 // 值复制
	fmt.Println(w3)    // {worker {worker2 0 22}}

	w3.age = 23

	fmt.Println(w2, w3)               // {worker {worker2 0 22}} {worker {worker2 0 23}}
	fmt.Printf("%p, %p \n", &w2, &w3) // 0xc000090030, 0xc000090090  # 所以是值复制

	var pw4 *Worker = &w3
	pw4.age = 24
	fmt.Println(w3, pw4)              // {worker {worker2 0 24}} &{worker {worker2 0 24}}
	fmt.Printf("%p, %p \n", &w3, pw4) // 0xc000090090, 0xc000090090
}
