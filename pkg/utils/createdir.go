package utils

import (
	"os"
	"path/filepath"
)

func CreateStratRepo(userID string, strategyID string) (string, error) {
	wd, _ := os.Getwd()
	newStratRepoPath := filepath.Join(wd, "repository", "strategies", userID, strategyID)
	return newStratRepoPath, os.MkdirAll(newStratRepoPath, os.ModePerm)
}
