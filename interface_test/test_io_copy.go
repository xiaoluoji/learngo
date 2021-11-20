package interface_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func TestIOCopy() {
	var b bytes.Buffer
	//将字符串写入 bytes.buffer类型变量
	b.Write([]byte("Hello"))
	//使用fprint将字符串拼接到buffer
	fmt.Fprint(&b, "World!")
	//将buffer内容输出到stdout
	io.Copy(os.Stdout, &b)

}
