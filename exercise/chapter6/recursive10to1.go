package chapter6

import "fmt"

//noinspection GoUnusedFunction
func Recursive10to1() {
	recursivePrint(1)
}

func recursivePrint(n int) {
	if n < 10 {
		recursivePrint(n + 1)
	}
	fmt.Printf("%d, ", n)
}
