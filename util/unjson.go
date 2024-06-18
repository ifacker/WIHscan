package util

import (
	"encoding/json"
	"fmt"
)

// 解析 json
func Unjson(jsonStr string) map[string]any {
	data := make(map[string]any)
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("JSON unmarshaling failed:", err)
		return nil
	}
	return data
}
