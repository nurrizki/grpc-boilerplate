package util

import "encoding/json"

// Stringify :
func Stringify(data interface{}) string {
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}

func CoalesceString(value, defaultValue string) string {
	if value != "" {
		return value
	}
	return defaultValue
}
