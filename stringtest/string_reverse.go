package stringtest

import "fmt"

//noinspection GoUnusedFunction
func StringReverse() {
	var s string = "str"
	fmt.Printf("字符串: %s 的长度/2等于%d \n", s, len(s)/2)
	s1 := s[len(s)/2:] + s[:len(s)/2]
	fmt.Println(s1)

	//反正字符串
	var sByte = []byte("Google")
	//var sByte = []byte("we are 中国人")
	//使用下面方法反正unicode字符串会出问题
	ss := string(sByte)
	for i, j := 0, len(sByte)-1; i < j; i, j = i+1, j-1 {
		sByte[i], sByte[j] = sByte[j], sByte[i]
	}
	fmt.Printf("字符串：\n %s \n反转之后是：\n %s \n", ss, sByte)

	//使用runes反转unicode字符串
	chinese := "we are 中国人"
	result := reverse(chinese)
	fmt.Printf("字符串：\n %s \n反转之后是：\n %s \n", chinese, result)
}

func reverse(s string) string {
	runes := []rune(s)
	fmt.Println("Runes切片：", runes)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	/*另外一种实现方法
	n, h := len(runes), len(runes)/2
	for i := 0; i < h; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	*/
	return string(runes)
}
