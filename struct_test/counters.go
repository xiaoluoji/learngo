package struct_test

import (
	"fmt"
	"learngo/test_package"
)

func Counters() {
	//使用test_package包公开的函数来创建一个未公开的变量
	//将工厂函数命名为New是Go语言的一个习惯
	counter := test_package.New(10)
	//如果视图直接创建未公开的alertCounter类型的值，则会报错
	//提示不能引用未公开的名字,因为alertCounter类型是使用小写字母声明的，所以这个标识符是未公开的
	//test:=test_package.alertCounter
	fmt.Printf("Counters: %d\n", counter)
	var test int
	test = int(test_package.New(20))
	fmt.Println("test value: ", test)

}
