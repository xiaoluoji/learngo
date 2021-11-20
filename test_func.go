package main

import (
	"fmt"
	"learngo/functest"
)

//noinspection GoUnusedFunction
func testFunc() {
	//fmt.Println("测试多返回值")
	//functest.MultipleReturn()
	//fmt.Println("测试传递变长参数")
	//functest.Varnumpar()
	//fmt.Println("测试defer")
	//functest.DeferTest()
	//functest.DeferTracing()
	//fmt.Println("求斐波那契数")
	//functest.Fibonacci()
	//fmt.Println("闭包测试")
	//functest.ClosureTest()
	//fmt.Println("利用闭包调试代码")
	//functest.ClosureForDebug()
	fmt.Println("求斐波那契数列（内存优化）")
	functest.FibonacciMemorization()

}
