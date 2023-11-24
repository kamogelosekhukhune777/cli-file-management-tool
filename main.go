package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kamogelosekhukhune777/cli-file-management-tool/cmd"
)

func main() {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createFilePath := createCmd.String("file", "", "File path to Create")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteFilePath := deleteCmd.String("file", "", "File path to delete")

	copyCmd := flag.NewFlagSet("copy", flag.ExitOnError)
	copySourceDest := copyCmd.String("source", "", "source file path")
	copySource := copyCmd.String("destination", "", "destination file path")

	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchRoot := searchCmd.String("root", ".", "root directory to search")
	searchFileName := searchCmd.String("file", "", "file name to search")

	if len(os.Args) < 2 {
		fmt.Println("please provide a sub-command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:])
		if *createFilePath != "" {
			cmd.CreateFile(*createFilePath)
		} else {
			fmt.Println("please provide a file path to create")
		}
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *deleteFilePath != "" {
			cmd.DeleteFile(*deleteFilePath)
		} else {
			fmt.Println("please provide a file to delete")
		}
	case "copy":
		copyCmd.Parse(os.Args[:2])
		if *copySourceDest != "" && *copySource != "" {
			cmd.CopyFile(*copySource, *copySourceDest)
		} else {
			fmt.Println("please provide source and destination file paths")
		}

	case "search":
		searchCmd.Parse(os.Args[:2])
		if *searchFileName != "" {
			searchResults := cmd.SearchFiles(*searchRoot, *searchFileName)
			if len(searchResults) > 0 {
				fmt.Println("files found:")
				for _, file := range searchResults {
					fmt.Println(file)
				}
			} else {
				fmt.Println("no files found:s")
			}
		} else {
			fmt.Println("Please provide a file name to search")
		}
	default:
		fmt.Println("Invalid Command")
		os.Exit(1)
	}
}
