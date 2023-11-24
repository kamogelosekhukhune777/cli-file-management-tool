package cmd

import (
	"fmt"
	"os"
)

func DeleteFile(filepath string) {
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println("error deleting file:", filepath)
	}

	fmt.Println("File", filepath, "successfully delted!")
}
