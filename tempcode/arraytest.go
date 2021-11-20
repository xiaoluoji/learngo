package tempcode

import "fmt"

//noinspection GoUnusedFunction
func TempArrayTest() {
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
	for k, v := range a {
		fmt.Println("Array item", k, "is", v)
	}
}
