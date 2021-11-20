package functest

import (
	"fmt"
	"learngo/util"
)

func Fibonacci() {
	execTime := util.TimeCost()
	for i := 0; i <= 41; i++ {
		result := fibonacciTest(i)
		fmt.Printf("fibonacci(%d) is :%d\n", i, result)
	}
	execTime()
}

func fibonacciTest(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacciTest(n-1) + fibonacciTest(n-2)
	}
	return
}
