package utils

import (
	"os"
	"path/filepath"
)

// Creates subdirectory in the current project.
func CreateDir(directories ...string) (string, error) {
	path := filepath.Join(directories...)
	if exists := Exists(path); !exists {
		return path, os.MkdirAll(path, os.ModePerm)
	}
	return path, nil
}

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}
