package utils

import (
	"os"
	"strconv"
)

func IsFileOrDirExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func GenerateInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
