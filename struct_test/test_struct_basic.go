package struct_test

import (
	"fmt"
)

type User struct {
	name       string
	email      string
	age        int
	privileged bool
}

// notify 使用值接收者实现了一个方法
func (u User) notify() {
	fmt.Printf("发送邮件给用户：%s<%s>\n", u.name, u.email)
}

// changeEmail 使用指针接收者实现了一个方法
func (u *User) changeEmail(email string) {
	u.email = email
}

//noinspection GoUnusedFunction
func TestStructBasic() {
	//使用多行形式创建自定义结构类型变量,内部变量不需要注意顺序
	bill := User{
		name:       "bill",
		email:      "bill@gmail.com",
		age:        17,
		privileged: false,
	}
	fmt.Println(bill)
	// user 类型的值可以用来调用使用值接收者声明的方法
	//方法notify接收到的是bill值的一个副本
	bill.notify()

	//user类型的值可以用来调用使用指针接收者声明的方法
	// 因为changeMail方法需要接收指针变量，go会对值类型的调用者做调整，编译器会引用bill得到一个指针，再调用方法
	bill.changeEmail("bill@163.com") //可以认为go语言实际执行了：&bill.changeEmail("bill@163.com")
	bill.notify()

	//使用单行形式创建自定义结构类型指针变量，内部变量需要注意顺序
	lisa := &User{"Bill", "lisa@gmail.com", 20, false}
	fmt.Println(lisa)

	// 指向user 类型的指针也可以用来调用使用值接收者声明的方法
	// 因为notify只是操作的一个值的副本，为了支持notify方法调用，编译器会将指针接引用为值
	lisa.notify() // 可以认为go语言实际执行了：  (*lisa).notify()
	// 指向user 类型的指针可以用来调用使用指针接收者声明的方法
	lisa.changeEmail("lisa@163.com")
	lisa.notify()
}
