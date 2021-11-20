package pointertest

import "fmt"

//noinspection GoUnusedFunction
func ModifyPointerValue() {
	s := "good bye"
	var p *string = &s
	*p = "hello"
	fmt.Printf("指针P的地址：%p\n", p)
	fmt.Printf("修改后的指针对应的值：%s\n", *p)
	fmt.Printf("初始变量S对应的值: %s\n", s)
}
