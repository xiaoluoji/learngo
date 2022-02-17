package chapter7

import "fmt"

//noinspection GoUnusedFunction
func TestInsertSlice() {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	in := []string{"A", "B", "C"}
	res := insertStringSlice(s, in, 0) // at the front
	fmt.Println(res)                   // [A B C M N O P Q R]
	res = insertStringSlice(s, in, 3)  // [M N O A B C P Q R]
	fmt.Println(res)
}

func insertStringSlice(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result

}
