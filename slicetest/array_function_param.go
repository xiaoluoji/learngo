package slicetest

import (
	"fmt"
	"learngo/util"
	"runtime"
)

//noinspection GoUnusedFunction
func TestFuncWithArray() {
	var arrint [1e7]int
	util.ReadMemStats()
	testFuncMemory(&arrint)
	testFuncMemroy2(arrint)
	fmt.Println("===> [force gc].")
	runtime.GC() //强制调用gc回收
	fmt.Println("===>[Done].")
	util.ReadMemStats()
}

func testFuncMemory(array *[1e7]int) {
	fmt.Println("引用数组传递测试")
	util.ReadMemStats()
}

func testFuncMemroy2(array [1e7]int) {
	fmt.Println("值类型数组传递测试")
	util.ReadMemStats()
}
