package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func SearchFiles(root, fileName string) []string {
	var foundFiles []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == fileName {
			foundFiles = append(foundFiles, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error searching files:", err)
	}

	return foundFiles
}
