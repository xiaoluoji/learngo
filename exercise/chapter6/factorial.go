package chapter6

import "fmt"

//noinspection GoUnusedFunction
func Factorial() {
	for i := 0; i < 20; i++ {
		result := testFactorial(uint64(i))
		//result := testFactorial2(uint64(i))
		fmt.Printf("%d 的阶乘结果为：%d\n", i, result)
	}
}

//noinspection GoUnusedFunction
func testFactorial(n uint64) (res uint64) {
	if n > 0 {
		res = n * testFactorial(n-1)
	} else {
		res = 1
	}
	return
}

//noinspection GoUnusedFunction
func testFactorial2(n uint64) uint64 {
	if n > 0 {
		return n * testFactorial(n-1)
	}
	return 1
}
