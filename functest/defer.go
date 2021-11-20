package functest

import "fmt"

//noinspection GoUnusedFunction
func DeferTest() {
	fmt.Println("defer的基本演示：")
	func1()
	fmt.Println()
	fmt.Println("使用for循环多次调用defer的演示：")
	multiDefer()
	fmt.Println("模拟数据库操作中使用defer：")
	dbDemoDefer()
}

func func1() {
	fmt.Println("现在是在函数1的最上面")
	defer func2()
	fmt.Println("现在是在函数1的最下面")
}

func func2() {
	fmt.Println("函数2: 延迟到调用函数2的函数的最下面执行（也就是函数1的最尾部）")
}

//当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
func multiDefer() {
	defer fmt.Println()
	defer fmt.Println()
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func dbDemoDefer() {
	connectToDB()
	fmt.Println("延迟数据库连接断开")
	defer disconnectFromDB()
	fmt.Println("进行数据库的增删改查操作！")
	fmt.Println("Oops! 程序异常crash或者网络发生错误！")
	fmt.Println("从函数中退出！")
	//return //结束程序
	//deferred 函数会在调用函数最后返回前执行，不管是正常的return还是非正常的程序终止
}

func connectToDB() {
	fmt.Println("ok,数据库已连接！")
}

func disconnectFromDB() {
	fmt.Println("ok,从数据库断开连接！")
}
