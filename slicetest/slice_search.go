package slicetest

import (
	"fmt"
	"sort"
)

//noinspection GoUnusedFunction
func SliceSearchTest() {
	var s1 = []int{8, 3, 1, 2, 7, 5, 4, 6}
	fmt.Println("切片S1: ", s1, " 是否排序：", sort.IntsAreSorted(s1))
	result := sort.SearchInts(s1, 4)
	fmt.Println("在切片中搜索 4的索引结果： ", result)
	//对切片s1进行排序
	sort.Ints(s1)
	fmt.Println("切片S1: ", s1, " 是否排序：", sort.IntsAreSorted(s1))
	result = sort.SearchInts(s1, 4)
	fmt.Println("在切片中搜索 4的索引结果： ", result)
}
