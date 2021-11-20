package interface_test

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//noinspection GoUnusedFunction
func CurlDemoSimple() {
	//从web得到响应
	//url := "http://www.google.com"
	url := "http://172.20.20.50"
	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	//从body复制到stdout
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
