package file

import (
	"os"
	"path/filepath"
)

func Save(path string, filename string, data []byte) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	path = filepath.Join(dir, path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}

	file, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return err
	}

	return nil
}
