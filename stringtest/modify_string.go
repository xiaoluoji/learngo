package stringtest

import "fmt"

func ModifyString() {
	s := "hello"
	//Go 语言中的字符串是不可变的，也就是说 str[index] 这样的表达式是不可以被放在等号左侧的。如
	//果尝试运行 str[i] = 'D' 会得到错误： cannot assign to str[i]
	//s[0] = "c"
	c := []byte(s)
	c[0] = 'c'
	fmt.Println(string(c))
}
