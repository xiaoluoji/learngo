package slicetest

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

//noinspection GoUnusedFunction
func FindDigits() {
	fileName := "D:/programming/go/learngo/tmp/findDigitsTest.txt"
	var results []byte
	results = findDigitsByFile(fileName)
	fmt.Println("打印第一次匹配")
	fmt.Println(string(results))
	results = findAllDigitsByFile(fileName)
	fmt.Println("打印所有匹配")
	fmt.Println(string(results))
}

func findDigitsByFile(filename string) []byte {
	var digitRegexp = regexp.MustCompile("[0-9]+")
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}
func findAllDigitsByFile(filename string) []byte {
	var digitRegexp = regexp.MustCompile("[0-9]+")
	fileBytes, _ := ioutil.ReadFile(filename)
	c := make([]byte, 0)
	b := digitRegexp.FindAll(fileBytes, len(fileBytes))
	for _, bytes := range b {
		c = append(c, bytes...)
	}
	return c
}
