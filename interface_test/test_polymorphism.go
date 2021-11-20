package interface_test

import "fmt"

//emailOperation 定义邮件操作的接口
type emailOperation interface {
	send()
	receive()
}

//定义guest用户类型
type guest struct {
	name  string
	email string
}

//定义admin用户类型
type admin struct {
	name  string
	email string
}

//使用指针类型实现emailOperation接口
func (u *guest) send() {
	fmt.Printf("发送邮件给guest用户：%s<%s>\n", u.name, u.email)
}
func (u *guest) receive() {
	fmt.Printf("guest用户收到邮件：%s<%s>\n", u.name, u.email)
}

//使用指针类型实现emailOperation接口
func (u *admin) send() {
	fmt.Printf("发送邮件给admin用户：%s<%s>\n", u.name, u.email)
}
func (u *admin) receive() {
	fmt.Printf("admin用户收到邮件：%s<%s>\n", u.name, u.email)
}

//send接手一个实现了emailOperaion接口的值
func send(u emailOperation) {
	u.send()
}

//send接手一个实现了emailOperaion接口的值
func receive(u emailOperation) {
	u.receive()
}

//noinspection GoUnusedFunction
func TestPolymorphism() {
	//创建一个guest用户
	guestUser := guest{"guest", "guest@gmail.com"}
	fmt.Println(guestUser)
	//执行send和receive操作
	send(&guestUser)
	receive(&guestUser)

	//创建一个admin用户
	adminUser := admin{"floyd", "floyd@gmail.com"}
	fmt.Println(adminUser)
	//执行send和receive操作
	send(&adminUser)
	receive(&adminUser)

}
