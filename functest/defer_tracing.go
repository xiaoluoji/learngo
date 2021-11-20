package functest

import "fmt"

//noinspection GoUnusedFunction
func DeferTracing() {
	b()
}

func trace(s string)   { fmt.Println("进入函数：", s) }
func untrace(s string) { fmt.Println("离开函数：", s) }

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("在函数A中")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("在函数B中")
	a()
}
