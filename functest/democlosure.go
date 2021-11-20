package functest

import "fmt"

func ClosureTest() {
	x := func(n int) int {
		result := 0
		for i := 0; i <= n; i++ {
			result = result + i
		}
		return result
	}(200)
	//x := test(100)
	fmt.Println("100的累加和：", x)
	fmt.Println("演示匿名函数的地址：")
	f()
}

func f() {
	for i := 0; i <= 5; i++ {
		g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}
