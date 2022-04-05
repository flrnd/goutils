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

func parseArgs(args []string) (string, string) {
	switch {
	case len(args) == 2:
		return args[0], args[1]
	default:
		fmt.Println("help")
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
				c <- path
			}
			return
		})
		defer close(c)
	}()
	return c
}
