package pointertest

import "fmt"

//noinspection GoUnusedFunction
func Pointertest() {
	var i1 = 5
	fmt.Printf("整数变量i1: %d, 变量地址为：%p\n", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("内存地址为: %p 的值为：%d\n", intP, *intP)
}
