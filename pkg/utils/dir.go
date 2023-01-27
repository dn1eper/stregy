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

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
