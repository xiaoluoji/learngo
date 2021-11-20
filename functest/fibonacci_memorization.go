package functest

import (
	"fmt"
	"learngo/util"
)

const LIM = 41

var fibs [LIM]uint64

//noinspection GoUnusedFunction
func FibonacciMemorization() {
	result := uint64(0)
	execTime := util.TimeCost()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d)is : %d\n", i, result)
	}
	execTime()
}

func fibonacci(n int) (res uint64) {
	//检查fibonacci(n)是否已经被计算过
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}
