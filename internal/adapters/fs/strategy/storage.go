package strategy

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"stregy/internal/domain/strategy"
	"stregy/pkg/utils"
)

type storage struct{}

func NewStorage() strategy.Storage {
	return storage{}
}

func (s storage) SaveStrategy(implementation, userID, strategyID string) error {
	dec, _ := base64.StdEncoding.DecodeString(implementation)
	dirpath, _ := utils.CreateDir([]string{"repository", "strategies", userID, strategyID})
	f, err := os.Create(filepath.Join(dirpath, "strategy"))
	defer f.Close()
	if err != nil {
		return err
	}
	f.Write(dec)
	return nil
}
