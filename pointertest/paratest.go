package pointertest

import "fmt"

//noinspection GoUnusedFunction
func ParaAtest(a *int) {
	b := a
	fmt.Printf("地址 %p 对应的值： %d \n", b, *b)
}

//noinspection GoUnusedFunction
func ParaBtest(a int) {
	b := &a
	fmt.Printf("地址 %p 对应的值： %d \n", b, *b)
}
