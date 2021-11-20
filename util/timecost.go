package util

import (
	"fmt"
	"time"
)

func TimeCost() func() {
	//因为闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。
	//所以此处start变量会在第一次创建的时候变保存下来，不管后面调用几次TimeCose返回的闭包函数，都会得到第一次创建到调用闭包函数的时间差
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("程序执行消耗时间: %v\n", tc)
	}
}
