package strategy

import (
	"os"
	"path/filepath"
	"stregy/internal/domain/strategy"
	"stregy/pkg/utils"
)

type storage struct{}

func NewStorage() strategy.Storage {
	return storage{}
}

func (s storage) SaveStrategy(implementation *string, userID, strategyID string) error {
	dirpath, _ := utils.CreateDir([]string{"repository", "strategies", userID, strategyID})
	f, err := os.Create(filepath.Join(dirpath, "strategy"))
	defer f.Close()
	if err != nil {
		return err
	}
	f.Write([]byte(*implementation))
	return nil
}
