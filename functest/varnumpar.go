package functest

import "fmt"

//noinspection GoUnusedFunction
func Varnumpar() {
	x := min(1, 3, 2, 0)
	fmt.Printf("传入参数的最小值为：%d\n", x)
	//如果参数被存储在一个 slice 类型的变量 slice 中，则可以通过 slice... 的形式来传递参数，调用变参函数。
	slice := []int{7, 9, 3, 5, 1}
	x = min(slice...)
	fmt.Printf("数组中最小值为：%d\n", x)
}
func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	x := s[0]
	for _, v := range s {
		if v < x {
			x = v
		}
	}
	return x
}
