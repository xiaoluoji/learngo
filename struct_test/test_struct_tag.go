package struct_test

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool   `json:"是否重要回答"`
	field2 string `json:"类型名称"`
	field3 int    `json:"有多少"`
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

//noinspection GoUnusedFunction
func TestStructTag() {
	tt := TagType{true, "floyd", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}
