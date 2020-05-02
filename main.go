package main

import(
    "fmt"
    "github.com/OneHundred86/gotest/pkg1"
    "github.com/OneHundred86/gotest/pkg2"
    "github.com/OneHundred86/gotest/pkg1/pkg11"
    "rsc.io/quote"
)

func main() {
    fmt.Println("hello world")

    // 本地包使用(本项目目录)
    fmt.Println(pkg1.GetPkg())
    pkg1.PrintA1Var()

    pkg2.SayHi()
    fmt.Println("pkg2调用pkg1的函数:", pkg2.CallPkg1Fuc())

    pkg11.SayHi()

    // 第三方模块 go.mod
    fmt.Println("第三方module调用：", quote.Hello())
}