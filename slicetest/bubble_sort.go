package slicetest

import (
	"fmt"
	"learngo/util"
	"math/rand"
	"sort"
	"time"
)

//noinspection GoUnusedFunction
func BubbleSort() {
	//arr := []int{8, 11, 3, 1, 9, 4, 2, 10, 7, 6, 5}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var arr []int
	for i := 0; i < 5000; i++ {
		arr = append(arr, r.Intn(5000))
	}
	unsortedArr := make([]int, len(arr))
	util.ReadMemStats()
	copy(unsortedArr, arr)
	util.ReadMemStats()

	//fmt.Println(unsortedArr)
	execTime := util.TimeCost()
	fmt.Println("数组长度：", len(arr))
	for i := len(arr) - 1; i > 1; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("冒泡排序耗时：")
	execTime()
	//fmt.Println(arr)
	fmt.Println("切片是否已排序：", sort.IntsAreSorted(arr))
	sort.Ints(unsortedArr)
	fmt.Println("使用sorts库排序耗时：")
	execTime()
	//fmt.Println(unsortedArr)
	fmt.Println("切片是否已排序：", sort.IntsAreSorted(unsortedArr))
}
