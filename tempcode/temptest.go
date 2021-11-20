package tempcode

import "fmt"

//noinspection GoUnusedFunction
func TempForTest() {
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("变量i的值: ", i)
		i++
	}
	fmt.Println("无限循环已终止")
}
