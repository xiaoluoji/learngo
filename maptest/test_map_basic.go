package maptest

import "fmt"

func TestMapBasic() {
	var colors map[string]string
	colors = make(map[string]string)
	colors["Red"] = "#da1337"
	colors["Blue"] = "#23425"
	if value, exists := colors["Blue"]; exists {
		fmt.Println("颜色 Blue 存在!", value)
	} else {
		fmt.Println("颜色 Blue不存在!", value)
	}
	colors = map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}
