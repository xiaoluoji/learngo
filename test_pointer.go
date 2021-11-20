package main

import (
	"fmt"
	"learngo/pointertest"
)

//noinspection GoUnusedFunction
func testPointer() {
	fmt.Println()
	fmt.Println("指针测试")
	//pointertest.Pointertest()
	//pointertest.ModifyPointerValue()
	var i = 10
	pointertest.ParaAtest(&i)
	pointertest.ParaBtest(i)

}
