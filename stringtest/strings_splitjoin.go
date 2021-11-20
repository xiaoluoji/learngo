package stringtest

import (
	"fmt"
	"strings"
)

func StringsSplitjoin(){
	str:="The quick brown fox jumps over the lazy dog"
	s1:=strings.Fields(str)
	fmt.Printf("分割后的字符串切片：%v\n",s1)
	for _,val:=range s1{
		fmt.Printf("%s - ",val)
	}
	fmt.Println()
	str2:="GO1|The ABC of Go|25"
	s12:=strings.Split(str2,"|")
	fmt.Printf("分割后的字符串切片：%v\n",s12)
	for _,val:=range s12{
		fmt.Printf("%s - ",val)
	}
	fmt.Println()
	str3:=strings.Join(s12,";")
	fmt.Printf("使用;连接分割后的字符串：%s\n",str3)
}