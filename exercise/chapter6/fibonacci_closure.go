package chapter6

//noinspection GoUnusedFunction
func FibonacciClosure() {
	f := fib()
	for i := 0; i <= 41; i++ {
		println(i+2, f())
	}
}

func fib() func() int {
	//闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}
