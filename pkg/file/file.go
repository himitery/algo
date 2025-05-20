package file

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Save[T any](path string, data T) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if _, err := file.Write(jsonData); err != nil {
		return err
	}

	return nil
}
