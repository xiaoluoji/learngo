package main

import (
	"fmt"
	"learngo/basictest"
)

//noinspection GoUnusedFunction
func testBasic() {
	fmt.Println()
	fmt.Println("switch测试")
	basictest.SwitchTest()
	fmt.Println("goto测试")
	basictest.GotoTest()
}
