package strategy

import (
	"os"
	"path"
	"path/filepath"
	"stregy/internal/config"
	"stregy/internal/domain/strategy"
	"stregy/pkg/utils"
)

type storage struct{}

func NewStorage() strategy.Storage {
	return storage{}
}

func (s storage) SaveStrategy(name string, implementation *string) (string, error) {
	var err error
	defer func() {
		if err != nil {
			panic(err)
		}
	}()

	dir := path.Join(config.GetConfig().StratexecProjectPath, "local", "strategies", name)
	os.Mkdir(dir, os.ModePerm)
	utils.RemoveContents(dir)

	archivePath := filepath.Join(dir, "strategy.zip")
	f, err := os.Create(archivePath)
	if err != nil {
		return "", err
	}
	f.Write([]byte(*implementation))
	f.Close()

	utils.Unzip(archivePath, dir)
	os.Remove(archivePath)

	return dir, nil
}
