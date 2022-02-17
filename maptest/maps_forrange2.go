package maptest

import "fmt"

//noinspection GoUnusedFunction
func MapsForange2() {
	//Version A:
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	//Version B:
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1) //此处item只是item32中的元素的拷贝
		item[1] = 2                 //这里的item会在下一次迭代的时候丢失，因为只是一个拷贝
	}
	fmt.Printf("Version B: Value of itmes: %v\n", items2)
}
