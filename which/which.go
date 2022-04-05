package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
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

func handleArguments(argumentList []string) {
	var waitGroup sync.WaitGroup

	for _, argument := range argumentList {
		waitGroup.Add(1)
		go func(command string) {
			defer waitGroup.Done()
			found, err := getCommandPath(command)
			if err != nil {
				fmt.Printf("%s", err)
			}
			fmt.Printf("%s\n", found)
		}(argument)
	}
	waitGroup.Wait()
}

func main() {
	arguments := os.Args

	switch {
	case len(arguments) == 1:
		fmt.Println("Need more arguments")
	case len(arguments) > 1:
		handleArguments(arguments[1:])
	default:
		fmt.Printf("This should not happen: %s!", arguments)
	}
}
