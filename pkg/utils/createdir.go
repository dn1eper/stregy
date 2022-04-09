package utils

import (
	"os"
	"path/filepath"
)

// Creates subdirectory in the current project.
func CreateDir(directories []string) (string, error) {
	directoriesWithWd := make([]string, len(directories)+1)
	copy(directoriesWithWd[1:], directories)
	wd, _ := os.Getwd()
	directoriesWithWd[0] = wd
	newStratRepoPath := filepath.Join(directoriesWithWd...)
	return newStratRepoPath, os.MkdirAll(newStratRepoPath, os.ModePerm)
}
