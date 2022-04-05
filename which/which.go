package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func getCommandPath(file string) (string, error) {
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, directory := range pathSplit {
		fullpath := filepath.Join(directory, file)
		fileInfo, err := os.Stat(fullpath)

		if err == nil {
			mode := fileInfo.Mode()
			if mode.IsRegular() {
				if mode&0111 != 0 {
					return fullpath, nil
				}
			}
		}
	}

	errorMessage := fmt.Sprintf("%s not found", file)
	return "", errors.New(errorMessage)
}

func findCommands(arguments []string) {
	for _, arg := range arguments {
		found, err := getCommandPath(arg)
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Println(found)
	}
}

func main() {
	arguments := os.Args

	switch {
	case len(arguments) == 1:
		fmt.Println("Need more arguments")
	case len(arguments) > 1:
		findCommands(arguments[1:])
	default:
		fmt.Printf("This should not happen: %s!", arguments)
	}
}
