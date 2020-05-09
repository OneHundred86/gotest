package base

import (
	"fmt"
	"reflect"
	"strings"
)

func testString() {
	str1 := "hello, world"
	fmt.Println(str1)

	str2 := `hello, world`
	fmt.Println(str2, reflect.TypeOf(str2))

	str3 := "你好"
	fmt.Println("字符串长度:", len(str3), ";文字个数:", len(strings.Split(str3, "")))

	str4 := str1 + str3
	fmt.Println(str4)

	var pstr *string
	pstr = &str4
	fmt.Println(*pstr)
}

func TestBuiltin() {
	testString()
	fmt.Printf("builtin.go: struct.go: %+v %v \n", W, W.name)
}
