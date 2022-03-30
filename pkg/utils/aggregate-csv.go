package utils

import (
	"io"
	"os"
	"path"
	"path/filepath"
)

func aggregateCSV(source, destination, outFilename string) error {
	output_filepath := path.Join(destination, outFilename)
	os.Remove(output_filepath)
	output_file, err := os.OpenFile(output_filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer output_file.Close()

	var input_filepaths []string
	err = filepath.Walk(source, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		filename := info.Name()
		if filename != outFilename && filename[len(filename)-4:] == ".csv" {
			input_filepaths = append(input_filepaths, path.Join(source, filename))
		}
		return nil
	})
	if err != nil {
		return err
	}

	for _, input_filepath := range input_filepaths {

		input_file, err := os.Open(input_filepath)
		if err != nil {
			return err
		}
		defer input_file.Close()

		_, err = io.Copy(output_file, input_file)
		if err != nil {
			output_file.Close()
			os.Remove(output_filepath)
			return err
		}
		input_file.Close()
		os.Remove(input_filepath)
	}

	return nil
}
