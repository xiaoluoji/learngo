package maptest

import "fmt"

//noinspection GoUnusedFunction
func MapFuncTest() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}
	for k, v := range mf {
		fmt.Println("Key: ", k, " Value: ", v())
	}

}
