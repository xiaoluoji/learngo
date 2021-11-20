package chapter6

import (
	"fmt"
	"strings"
)

//noinspection GoUnusedFunction
func FactoryTest() {
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	fileName := addBmp("testbmp")
	fmt.Println("执行addbmp匿名函数：", fileName)
	fileName = addJpeg("testJpeg")
	fmt.Println("执行addJpeg匿名函数：", fileName)
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
