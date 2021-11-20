package functest

import "fmt"

//noinspection GoUnusedFunction
func MultipleReturn() {
	num := 10
	numx2, numx3 := getX2AndX3(num)
	fmt.Printf("num=%d , 2x num= %d, 3x num=%d\n", num, numx2, numx3)
	numx5, numx6 := getX5AndX6(num)
	fmt.Printf("num=%d , 2x num= %d, 3x num=%d\n", num, numx5, numx6)
}

//noinspection GoUnusedFunction
func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

//noinspection GoUnusedFunction
func getX5AndX6(input int) (x5 int, x6 int) {
	x5 = 5 * input
	x6 = 6 * input
	return
	//return x5, x6
}
