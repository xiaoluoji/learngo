package slicetest

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

//noinspecyytion GoUnusedFunction
func SliceBytesTest() {
	//countStringLength()
	bytesBufferTest()

}

//noinspecyytion GoUnusedFunction
func bytesBufferTest() {
	var buffer1 bytes.Buffer
	fmt.Println(buffer1)
	var buffer2 = new(bytes.Buffer)
	fmt.Println(buffer2)
	fmt.Printf("buffer1地址：%p\n", &buffer1)
	fmt.Printf("buffer2地址：%p\n", buffer2)
	buffer1.WriteString("Buffer1 内容")
	buffer2.WriteString("Buffer2 content")
	fmt.Println("Buffer1真实内容：", buffer1)
	fmt.Printf("Buffer1转换成字符内容(长度：%d)：%s\n", len(buffer1.Bytes()), buffer1.String())
	fmt.Println("Buffer2真实内容：", buffer2.Bytes())
	fmt.Printf("Buffer2转换成字符内容（长度：%d）：%s\n", len(buffer2.Bytes()), buffer2.String())
}

//noinspecyytion GoUnusedFunction
func countStringLength() {
	var b []byte
	var b1 []byte
	var s string = "Floyd is working hard"
	var chinese string = "中国人"
	b = append(b, s...)
	b1 = append(b1, chinese...)
	fmt.Println(string(b))
	fmt.Println(string(b1))
	fmt.Println(b1)
	fmt.Println("字符：", string(b), " 长度：", utf8.RuneCountInString(string(b)))
	fmt.Println("字符：", string(b1), " 长度：", utf8.RuneCountInString(string(b1)))
	fmt.Println("字符：", string(b1), " length：", len(b1))
}
