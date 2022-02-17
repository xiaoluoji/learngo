package main

import (
	"fmt"
	"learngo/util"
)

var timeCost = util.TimeCost()

func main() {
	//testString()
	//testPointer()
	//testTime()
	//testBasic()
	testExercise()
	//testTempCode()
	//testFunc()
	//testArraySlice()
	//testMap()
	//util.GetCurrentPath()
	//testStruct()
	//testInterface()
	fmt.Print("程序总耗时：")
	timeCost()
}
