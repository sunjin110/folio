package utils

import (
	"os"
	"path/filepath"
)

func ProjectRoot() string {
	currentDir, err := os.Getwd()
	if err != nil {
		return ""
	}
	for {
		_, err := os.ReadFile(filepath.Join(currentDir, "go.mod"))
		if os.IsNotExist(err) {
			if currentDir == filepath.Dir(currentDir) {
				return ""
			}
			currentDir = filepath.Dir(currentDir)
			continue
		} else if err != nil {
			return ""
		}
		break
	}
	return currentDir
}

func ToPointer[T any](t T) *T {
	return &t
}

func SliceToMap[K comparable, V any](slice []V, getKey func(v V) K) map[K]V {
	m := map[K]V{}
	for _, v := range slice {
		k := getKey(v)
		m[k] = v
	}
	return m
}
