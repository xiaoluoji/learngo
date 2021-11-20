package basictest

import "fmt"

//noinspection GoUnusedFunction
func GotoTest() {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("变量i当前的值：%d 变量j的值: %d\n", i, j)
		}
	}
}
