package cmd

import (
	"fmt"
	"os"
)

func CopyFile(sourcePath, destPath string) {
	input, err := os.ReadFile(sourcePath)
	if err != nil {
		fmt.Println("error reading file", err)
		return
	}

	err = os.WriteFile(destPath, input, 0644)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}

	fmt.Println("File", sourcePath, "copied to", destPath, "succesfully!")
}
