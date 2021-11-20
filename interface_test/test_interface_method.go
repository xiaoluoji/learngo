package interface_test

import (
	"fmt"
)

//notifier 定义个通知类行为的接口
type notifier interface {
	notify()
	receive()
}

//user 定义了一个用户类型
type user struct {
	name  string
	email string
}

//notify是使用指针接收者实现的方法
func (u *user) notify() {
	fmt.Printf("发送邮件给用户：%s<%s>\n", u.name, u.email)
}

//receive是使用值接收者实现的方法
func (u user) receive() {
	fmt.Printf("收到邮件：%s<%s>\n", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}
func receiveEmail(n notifier) {
	n.receive()
}

//noinspection GoUnusedFunction
func TestInterfaceMethod() {
	//这里floyd变量是user的值类型，按照接口规范里的定义
	//T类型的值（即某类型的值变量）的方法集只包含值接收者声明的方法
	//而本示例中func (u *user) notify()只实现了指针接收的方法
	floyd := user{"floyd", "floyd@gmail.com"}
	//不能将u（类型是user）作为sendNotification 的参数类型notifier：
	//user 类型并没有实现notifier  notify 方法使用指针接收者声明）
	//如果按下面的方式调用sendNotification会出现上面注释的报错
	//sendNotification(floyd)
	//正确的方法是按照下面的，传递指针给到调用接口值的函数
	sendNotification(&floyd)
	//为什么会有这种限制？因为编译器不是总能自动获取到一个值的地址
	receiveEmail(&floyd)

	//这里bill变量是user的指针类型，按照接口规范里的定义
	//T类型的指针（某类型的指针变量）的方法集既包含值接收者声明的方法，也包含指针接收者声明的方法
	bill := &user{"bill", "bill@gmail.com"}
	sendNotification(bill)
	receiveEmail(bill)

}
