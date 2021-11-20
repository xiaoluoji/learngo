package slicetest

import "fmt"

//noinspection GoUnusedFunction
func TestSliceBasic() {
	//创建一个整形切片
	//其长度和容量都是5
	slice := []int{10, 20, 30, 40, 50}
	//创建一个新切片
	//其长度和是2个元素，容量是4个元素（）
	//对底层数组容量是K的slice[i:j]来说，长度=j-i,容量等于k-i
	//对底层数组容量是5 的切片slice[1:3]来说, 3-1=2（长度）, 5-1=4（容量）
	newSlice := slice[1:3]
	//修改了slice索引为1的值，同时也修改了原来的slice的索引为2的值
	newSlice[1] = 35
	fmt.Println("打印旧slice")
	printSlice(slice)
	fmt.Println("打印新slice")
	printSlice(newSlice)
	//修改newSlice索引为3的元素
	//这个元素对于newSlice来说并不存在
	//panic: runtime error: index out of range [3] with length 2
	//newSlice[3] = 45
	//使用原来的容量来分配一个新元素,将新元素赋值为60
	newSlice = append(newSlice, 60)
	fmt.Println("打印append之后的newSlice")
	printSlice(newSlice)
	fmt.Println("打印append之后的旧切片")
	printSlice(slice)
	//继续为新的slice追加2个元素，由于原来的底层数组没有足够的容量append 函数会创建一个新的底层数组，将被引用的现有的值复制到新数组里，再追加新的值
	//函数append 会智能地处理底层数组的容量增长。在切片的容量小于1000 个元素时，总是
	//会成倍地增加容量。一旦元素个数超过1000，容量的增长因子会设为1.25，也就是会每次增加25%
	//的容量。随着语言的演化，这种增长算法可能会有所改变。
	newSlice = append(newSlice, 80, 90)
	fmt.Println("打印继续append之后的newslice")
	printSlice(newSlice)
	fmt.Printf("newslice的容量 %d\n", cap(newSlice)) //此时新切片的容量已经扩容到8

	fmt.Println("打印继续append之后的旧切片")
	printSlice(slice)
	fmt.Printf("旧切片的容量 %d\n", cap(slice)) //旧切片的容量还是5

	//创建切片时的三个索引
	fmt.Println()
	source := []string{"apple", "orange", "plum", "banana", "grape"}
	var source2 = make([]string, len(source))
	copy(source2, source)
	fmt.Println("打印原始切片内容：")
	printSliceString(source)
	//将第三个元素切片，并限制容量，长度是：3-2=1，容量是4-2=2
	fruitSlice := source[2:3:4]
	fmt.Printf("新水果切片的内容(容量%d)：\n", cap(fruitSlice))
	printSliceString(fruitSlice)
	fruitSlice = append(fruitSlice, "cherry")
	fmt.Printf("修改新水果切片的内容(容量%d): \n", cap(fruitSlice))
	printSliceString(fruitSlice)
	fmt.Printf("修改新切片之后原始切片的内容(容量%d )：\n", cap(source)) //测试原来的souce切片内容也会受到更改
	printSliceString(source)

	//我们之前讨论过，内置函数append 会首先使用可用容量。一旦没有可用容量，会分配一个
	//新的底层数组。这导致很容易忘记切片间正在共享同一个底层数组。一旦发生这种情况，对切片
	//进行修改，很可能会导致随机且奇怪的问题。对切片内容的修改会影响多个切片，却很难找到问题的原因。
	//如果在创建切片时设置切片的容量和长度一样，就可以强制让新切片的第一个append 操作
	//创建新的底层数组，与原有的底层数组分离。新切片与原有的底层数组分离后，可以安全地进行
	//后续修改
	fmt.Println()
	fmt.Println("创建切片使用了3索引方式")
	fruitSlice2 := source2[2:3:3]
	fmt.Printf("打印设置了容量和长度一样的新水果切片内容(容量%d) :\n", cap(fruitSlice2))
	printSliceString(fruitSlice2)
	fmt.Printf("打印source2切片内容(容量%d) :\n", cap(source2))
	printSliceString(source2)
	fruitSlice2[0] = "mango"
	fmt.Printf("修改(只是修改不append)容量和长度一样的新水果切片内容(容量%d) :\n", cap(fruitSlice2))
	printSliceString(fruitSlice2)
	fmt.Printf("打印修改后source2切片内容(容量%d) :\n", cap(source2))
	printSliceString(source2)
	fruitSlice2 = append(fruitSlice2, "peach")
	fmt.Printf("修改(使用append)容量和长度一样的新水果切片内容(容量%d) :\n", cap(fruitSlice2))
	printSliceString(fruitSlice2)
	fmt.Printf("打印新切片使用了append只有后source2切片内容(容量%d) :\n", cap(source2))
	printSliceString(source2)

	//将一个切片追加到另外一个切片
	fmt.Println("将两个切片追加到一起:")
	fmt.Printf("%v\n", append(fruitSlice2, fruitSlice...))
}

func printSlice(slice []int) {
	for k, v := range slice {
		fmt.Printf("索引：%d  值: %d\n", k, v)
	}
}
func printSliceString(slice []string) {
	for k, v := range slice {
		fmt.Printf("索引：%d  值: %s\n", k, v)
	}
}
