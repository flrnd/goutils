package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	arguments := os.Args
	keyword, location := parseArgs(arguments[1:])
	channels := found(keyword, location)
	for msg := range channels {
		fmt.Println(msg)
	}
}

func printHelp(arg string) {
	fmt.Printf("found v0.1\n")
	fmt.Printf("Find all matching files and directories that contains a given keyword.\n\n")
	fmt.Printf("Usage: %s keyword path\n", arg)
	fmt.Printf("Example: %s config /etc\n", arg)
}

func parseArgs(args []string) (string, string) {
	switch {
	case len(args) == 2:
		return args[0], args[1]
	default:
		printHelp("found")
	}
	return "", ""
}

func found(keyword, location string) chan string {
	c := make(chan string)
	go func() {
		filepath.WalkDir(location, func(path string, d os.DirEntry, e error) (err error) {
			if e != nil {
				return e
			}
			if strings.Contains(path, keyword) {
				if d.IsDir() {
					c <- fmt.Sprintf("%s/", path)
				} else {
					c <- path
				}
			}
			return
		})
		defer close(c)
	}()
	return c
}
