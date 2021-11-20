package chapter5

import "fmt"

//noinspection GoUnusedFunction
func BitwiseComplement() {
	for i := -10; i <= 10; i++ {
		fmt.Printf("数字 %d 的补码是：%b \n", i, ^i)
	}
}
