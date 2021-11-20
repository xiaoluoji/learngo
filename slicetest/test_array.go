package slicetest

import (
	"fmt"
	"learngo/util"
)

//noinspection GoUnusedFunction
func ArrayTest() {
	//声明数组
	arr1 := [5]int{1: 10, 2: 20}
	for pos, val := range arr1 {
		fmt.Printf("变量arr1 索引：%d 的值为：%d\n", pos, val)
	}
	fmt.Println()

	//指针数组
	arrPoint := [5]*int{0: new(int), 1: new(int)}
	*arrPoint[0], *arrPoint[1] = 100, 200
	for k, v := range arrPoint {
		if v != nil {
			fmt.Printf("指针数组arrPoint索引：%d 的内存地址为：%v 值为：%d \n", k, v, *v)
		} else {
			fmt.Printf("指针数组arrPoint索引：%d 的内存地址为：%v \n", k, v)
		}
	}
	//同类型数组互相赋值
	fmt.Println()
	var arr2 [5]string
	arr3 := [5]string{"red", "blue", "green", "yellow", "pink"}
	arr2 = arr3
	for k, v := range arr2 {
		fmt.Printf("字符串数组arr3 索引: %d 的值为： %s\n", k, v)
	}

	//指针数组互相赋值
	fmt.Println()
	var arr4, arr5 [3]*string
	arr4 = [3]*string{new(string), new(string), new(string)}
	*arr4[0] = "red"
	*arr4[1] = "blue"
	*arr4[2] = "green"
	arr5 = arr4
	fmt.Println("打印两个相等的指针数组")
	fmt.Println(arr4)
	fmt.Println(arr5)
	for k, v := range arr4 {
		fmt.Printf("指针数组arr4 索引: %d 的地址为：%p  值为: %s\n", k, v, *v)
	}
	for k, v := range arr5 {
		fmt.Printf("指针数组arr5 索引: %d 的地址为：%p  值为: %s\n", k, v, *v)
	}

	util.ReadMemStats()
}
