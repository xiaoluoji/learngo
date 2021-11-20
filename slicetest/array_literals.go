package slicetest

import "fmt"

func ArrayLiterals() {
	var arrAge = [5]int{18, 25, 12, 28, 32}
	fmt.Println(arrAge)
	var arrLazy = [...]int{5, 6, 7, 8, 10}
	fmt.Println(arrLazy)
	var arrLazy2 = []int{5, 6, 7, 8, 10}
	fmt.Println(arrLazy2)
	var arrKeyValue = [5]string{3: "chris", 4: "ron"}
	fmt.Println(arrKeyValue)
	var arrKeyValue2 = []string{2: "john", 3: "michael"}
	fmt.Println(arrKeyValue2)

}
