package cmd

import (
	"fmt"
	"os"
)

func CreateFile(filepath string) {
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("error creating the file", err)
		return
	}
	defer file.Close()

	fmt.Println("File:", file, "Created succesfully!")
}
