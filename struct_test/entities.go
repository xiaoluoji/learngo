package struct_test

import (
	"fmt"
	"learngo/test_package"
)

//noinspection GoUnusedFunction
func Entities() {
	a := test_package.Admin{Rights: 10}
	//设置未公开的内部类型的公开字段的值
	a.Name = "Bill"
	a.Email = "bill@gmail.com"
	fmt.Printf("User: %v\n", a)
}
