package functest

import (
	"fmt"
	"log"
	"runtime"
)

//noinspection GoUnusedFunction
func ClosureForDebug() {
	//当您在分析和调试复杂的程序时，无数个函数在不同的代码文件中相互调用，如果这时候能够准确地知道哪个文件中的
	//具体哪个函数正在执行，对于调试是十分有帮助的。您可以使用 runtime 或 log 包中的特殊函数来实现这样的功
	//能。包 runtime 中的函数 Caller() 提供了相应的信息，因此可以在需要的时候实现一个 where() 闭包函数来打印函
	//数执行的位置
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	fmt.Println("现在调用测试函数A")
	testDebugA()
	where()
	fmt.Println("现在调用测试函数B")
	testDebugB()
	where()
}

func testDebugA() {
	fmt.Println("执行测试函数A代码")
}

func testDebugB() {
	fmt.Println("执行测试函数B代码")
}
