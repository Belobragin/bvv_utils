package util

import "strings"

// provide file extension:
func GetFileExtension(fileName string) string {
	temp := strings.Split(fileName, ".")
	return temp[len(temp)-1]
}
