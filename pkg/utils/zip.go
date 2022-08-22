package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Zip(src, dst string) error {
	// TODO
	return nil
}

func Unzip(src, dst string) error {
	log.Info(fmt.Sprintf("Unzipping '%v' to '%v'", src, dst))

	// 1. Open the zip file
	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	// 2. Get the absolute destination path
	dst, err = filepath.Abs(dst)
	if err != nil {
		return err
	}

	// 3. Iterate over zip files inside the archive and unzip each of them
	for _, f := range reader.File {
		err := unzipFile(f, dst)
		if err != nil {
			return err
		}
	}

	return nil
}

func unzipFile(f *zip.File, dst string) error {
	// 4. Check if file paths are not vulnerable to Zip Slip
	filePath := filepath.Join(dst, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	// 5. Create directory tree
	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// 6. Create a destination file for unzipped content
	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// 7. Unzip the content of a file and copy it to the destination file
	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}
	return nil
}
