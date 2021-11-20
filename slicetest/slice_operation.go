package slicetest

import "fmt"

//noinspection GoUnusedFunction
func SliceOperationTest() {
	var a = []int{0, 1, 2, 3, 4, 5, 6, 7}
	i := 4
	//通过临时变量s1实现删除切片a中指定位置i的元素
	fmt.Println("通过临时变量s1实现删除切片a中指定位置i的元素")
	s1 := a[:i:i]
	fmt.Println(s1)
	fmt.Println(a)
	s1 = append(s1, a[i+1:]...)
	fmt.Println(s1)
	fmt.Println(a)
	//通过copy切片a的i+1位置后面的元素到i位置实现删除i位置的元素
	fmt.Println("通过copy切片a的i+1位置后面的元素到i位置实现删除i位置的元素")
	n := copy(a[i:], a[i+1:])
	fmt.Println(n)
	fmt.Println(a)
	s2 := a[:i+n]
	fmt.Println(s2)
	a = a[:len(a)-1]
	fmt.Println(a)
	//实现pop操作
	fmt.Println("实现pop操作")
	var x int
	x, a = a[len(a)-1], a[:len(a)-1]
	fmt.Println(x, a)
}
