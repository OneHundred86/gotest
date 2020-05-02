package pkg2

import(
    "fmt"
    "github.com/OneHundred86/gotest/pkg1"
)

func SayHi(){
    fmt.Println("hi, this is pkg2")
}

// 调用pkg1的函数
func CallPkg1Fuc() string {
    return pkg1.GetPkg()
}