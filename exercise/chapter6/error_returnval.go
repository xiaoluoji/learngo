package chapter6

import (
	"errors"
	"fmt"
	"math"
)

//noinspection GoUnusedFunction
func MySqrtTest(x float64) {
	fmt.Printf("使用函数mysqrt求 %f 的平方根：", x)
	result, err1 := mySqrt(x)
	if err1 != nil {
		fmt.Println("输入错误：返回值为：", result, err1)
	} else {
		fmt.Println("OK,mysqrt结果为：", result, err1)
	}
	fmt.Printf("使用函数mysqrt2求 %f 的平方根：", x)
	if ret, err2 := mySqrt2(x); err2 != nil {
		fmt.Println("输入错误：返回值为：", ret, err2)
	} else {
		fmt.Println("OK,mysqrt2结果为：", ret, err2)
	}
}

//noinspection GoUnusedFunction
func mySqrt(x float64) (result float64, err error) {
	if x < 0 {
		result = float64(math.NaN())
		err = errors.New("输入值为负数，无法计算平方根!")
		return
	}
	return math.Sqrt(x), nil
}

//noinspection GoUnusedFunction
func mySqrt2(x float64) (float64, error) {
	if x < 0 {
		return float64(math.NaN()), errors.New("无法计算输入值为负数的平方根！")
	}
	return math.Sqrt(x), nil
}
