package main

import (
	"fmt"

	// "github.com/OneHundred86/gotest/goroutine"

	"github.com/OneHundred86/gotest/base"
	"github.com/OneHundred86/gotest/logger"
	"github.com/OneHundred86/gotest/mysql"
	"github.com/OneHundred86/gotest/pkg1"
	"github.com/OneHundred86/gotest/pkg1/pkg11"
	"github.com/OneHundred86/gotest/pkg2"
	"rsc.io/quote"
)

func main() {
	fmt.Println("hello world")

	// go语言基本特性
	base.TestMap()
	base.TestBuiltin()
	base.TestTime()
	base.TestStruct()
	// fmt.Printf("%+v %v \n", base.W, base.W.name)

	// 本地包使用(本项目目录)
	fmt.Println(pkg1.GetPkg())
	pkg1.PrintA1Var()

	pkg2.SayHi()
	fmt.Println("pkg2调用pkg1的函数:", pkg2.CallPkg1Fuc())

	pkg11.SayHi()

	// 第三方模块 go.mod
	fmt.Println("第三方module调用：", quote.Hello())

	//
	// goroutine.Test()

	// channel.TestRange()

	logger.Test()
	logger.TestLogToFile()
	return

	// mysql.Test()
	mysql.TestGorm()
}
