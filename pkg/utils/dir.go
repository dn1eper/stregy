package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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

func CopyFolderWin(src, dst string, skipNames []string) error {
	sn := make(map[string]bool)
	for _, name := range skipNames {
		sn[name] = true
	}

	items, _ := ioutil.ReadDir(src)
	for _, item := range items {
		if _, exists := sn[item.Name()]; !exists {
			if item.IsDir() {
				cmd := exec.Command("Xcopy", "/e", "/h", "/c", "/i", "/y", filepath.Join(src, item.Name()), filepath.Join(dst, item.Name()))
				output, err := cmd.Output()
				if err != nil {
					return fmt.Errorf("%v: %v", err, string(output))
				}
			} else {
				cmd := exec.Command("cmd", "/c", "copy", filepath.Join(src, item.Name()), dst)
				output, err := cmd.Output()
				if err != nil {
					return fmt.Errorf(fmt.Sprintf("%v: %v", err, string(output)))
				}
			}
		}
	}

	return nil
}
