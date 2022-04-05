package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func main() {
	arguments := os.Args
	keyword, location, err := parseArgs(arguments[1:])

	if err != nil {
		printHelp("found")
		return
	}

	channels := found(keyword, location)
	for msg := range channels {
		fmt.Println(msg)
	}
}

func printHelp(arg string) {
	fmt.Printf("%s v0.1\n", arg)
	fmt.Printf("Find all matching files and directories that contains a given keyword.\n\n")
	fmt.Printf("Usage: %s keyword path\n", arg)
	fmt.Printf("Example: %s config /etc\n", arg)
}

func parseArgs(args []string) (string, string, error) {
	switch {
	case len(args) == 2:
		return args[0], args[1], nil
	default:
		return "", "", errors.New("wrong number of arguments")
	}
}

func found(keyword, location string) chan string {
	c := make(chan string)
	blue := color.New(color.FgBlue, color.Bold).SprintFunc()
	go func() {
		defer close(c)
		filepath.WalkDir(location, func(path string, d os.DirEntry, e error) (err error) {

			if e != nil {
				return e
			}
			if strings.Contains(path, keyword) {
				if d.IsDir() {
					c <- blue(path)
				} else {
					directory, file := splitOnLastIndex(path)
					c <- fmt.Sprintf("%s%s", blue(directory), file)
				}
			}
			return
		})
	}()
	return c
}

func splitOnLastIndex(s string) (string, string) {
	lastIndex := strings.LastIndex(s, "/")
	// We want to include the last / on directory string, but remove it from file
	index := lastIndex + 1
	return s[:index], s[index:]
}
