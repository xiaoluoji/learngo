package stringtest

import "fmt"

func ForString() {
	str := "Go is a beautiful language!"
	fmt.Printf("字符串 'Go is a beautiful language' 的长度是：%d\n", len(str))
	//for ix := 0; ix < len(str); ix++ {
	//	fmt.Printf("字符串位置 %d 的字符是：%c \n", ix, str[ix])
	//}
	for ix, char := range str {
		fmt.Printf("字符串位置 %d 的字符是：%c \n", ix, char)
	}
	str2 := "Chinese：中国话"
	fmt.Printf("字符串 'Chinese：中国话' 的长度是：%d\n", len(str2))
	//for ix := 0; ix < len(str2); ix++ {
	//	fmt.Printf("字符串位置 %d 的字符是: %c \n", ix, str2[ix])
	//}
	for ix, char := range str2 {
		fmt.Printf("字符串位置 %d 的字符是: %c \n", ix, char)
	}
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Println("index int(rune) rune    char     bytes")
	for index, rune := range str2 {
		fmt.Printf("%-4d    %-5d    %U    '%c'    % X \n", index, rune, rune, rune, []byte(string(rune)))
	}
}
