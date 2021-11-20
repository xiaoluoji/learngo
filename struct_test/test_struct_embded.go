package struct_test

import "fmt"

//emailOperation 定义邮件操作的接口
type emailOperation interface {
	send()
}

//定义guest用户类型
type guest struct {
	name  string
	email string
}

//定义admin用户类型
type admin struct {
	//嵌入一个类型只需要声明这个类型的名字就可以了
	//此时我们可以说guest是外部类型admin的内部类型
	guest
	level string
}

/*
type admin struct {
	name  string
	email string
}
*/

//使用指针类型实现guest类型的emailOperation接口方法
func (u *guest) send() {
	fmt.Printf("发送邮件给guest用户：%s<%s>\n", u.name, u.email)
}

//使用指针类型实现admin类型的emailOperation接口方法
//如果注释掉以下实现针对admin类型的接口方法
//因为admin类型是内嵌了guest类型，所以当admin类型的指针在调用接口方法send()的时候，会调用guest类型的接口实现方法
func (u *admin) send() {
	fmt.Printf("发送邮件给admin用户：%s<%s>\n", u.name, u.email)
}

//send接手一个实现了emailOperaion接口的值
func send(u emailOperation) {
	u.send()
}

//noinspection GoUnusedFunction
func TestStructEmbded() {
	//创建一个guest用户
	guestUser := guest{"guest", "guest@gmail.com"}
	fmt.Println(guestUser)
	//执行send和receive操作
	send(&guestUser)

	//创建一个admin用户
	adminUser := admin{guest: guest{"floyd", "floyd@gmail.com"}, level: "super"}
	fmt.Println(adminUser)
	//执行send和receive操作
	//如果没有实现针对admin类型的接口方法，则send函数会调用admin继承的针对guest类实现的接口方法
	send(&adminUser)
	//我们可以直接访问内部类型的方法
	adminUser.guest.send()

}
