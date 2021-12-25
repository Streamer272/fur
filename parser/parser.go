package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func FindByInDir(dir string, by string) ([]os.FileInfo, error) {
	var result []os.FileInfo

	content, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range content {
		match, err := regexp.MatchString(by, file.Name())
		if err != nil {
			return nil, err
		}
		if match {
			fmt.Printf("Found match! %v in %v\n", by, file.Name())
			result = append(result, file)
		}
	}

	return result, nil
}

func FindAllByPath(path string, root string) ([]string, error) {
	fmt.Printf("FindAllByPath called with path = \"%v\" and root = \"%v\"\n", path, root)

	var result []string

	// path = mnt/sda1/Desktop/(?)*\.py
	// root = /
	// root always ends with "/"

	foundFiles, err := FindByInDir(root, strings.Split(path, "/")[0])
	if err != nil {
		return nil, err
	}

	// foundFile = mnt
	for _, foundFile := range foundFiles {
		if foundFile.IsDir() {
			var newPath string
			var newRoot string

			newPath = strings.Join(strings.Split(path, "/")[1:], "/")
			newRoot = root + foundFile.Name() + "/"

			found, err := FindAllByPath(newPath, newRoot)
			if err != nil {
				return nil, err
			}

			result = append(result, found...)
		} else {
			result = append(result, root+foundFile.Name())
		}
	}

	fmt.Printf("FindAllByPath returning \"%v\"\n", result)

	return result, nil
}