package chapter7

import "fmt"

//noinspection GoUnusedFunction
func FibonacciArray() {
	var fibArr [50]int64
	fibArr[0], fibArr[1] = 1, 1
	for i := 2; i < 50; i++ {
		fibArr[i] = fibArr[i-1] + fibArr[i-2]
	}
	for k, v := range fibArr {
		fmt.Printf("fibonacci[%d] is: %v \n", k, v)
	}
}
