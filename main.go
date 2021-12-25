package main

import (
	"fmt"
	"github.com/Streamer272/fur/parser"
	"github.com/akamensky/argparse"
	"os"
	"strings"
)

const (
	Version = "1.0.1"
)

func main() {
	argParser := argparse.NewParser("fur", "Find using RegEx")
	argParser.ExitOnHelp(true)

	separator := argParser.String("s", "sep", &argparse.Options{Required: false, Help: "Output file separator", Default: " "})
	filesToFind := argParser.StringList("p", "path", &argparse.Options{Required: false, Help: "Path to find", Default: []string{}})
	version := argParser.Flag("v", "version", &argparse.Options{Required: false, Help: "Display version", Default: false})

	err := argParser.Parse(os.Args)
	if err != nil {
		fmt.Print(argParser.Usage(err))
	}
	if *version {
		fmt.Printf("fur version %v\n", Version)
		os.Exit(0)
	}

	var result []string

	for _, fileToFind := range *filesToFind {
		var root string
		if strings.HasPrefix(fileToFind, "/") {
			root = "/"
			fileToFind = fileToFind[1:]
		} else {
			root, err = os.Getwd()
			if err != nil {
				panic(err)
			}

			root = root + "/"

			if strings.HasPrefix(fileToFind, "./") {
				fileToFind = fileToFind[2:]
			}
		}

		found, err := parser.FindAllByPath(fileToFind, root)
		if err != nil {
			panic(err)
		}

		result = append(result, found...)
	}

	fmt.Printf("%v\n", strings.Join(result, *separator))
}
