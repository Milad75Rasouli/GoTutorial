package main

import (
	"fmt"
	"os"
	// "github.com/pkg/errors"
)

func GetDir(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err // errors.warp(err,"this is my error message that has been wraped")
	}
	var files_list []string
	for _, file := range files {
		file_name := file.Name()
		files_list = append(files_list, file_name)
		// if _, err := fmt.Println(file_name); err != nil {
		// 	return nil, err
		// }
	}
	return files_list, nil
}

func main() {
	files, err := GetDir("/")
	if err != nil {
		// write the error on stderr
		fmt.Fprintf(os.Stderr, "%v", err)
	} else {
		// write the files on stdout
		fmt.Fprintf(os.Stdout, "%v", files)
	}
}
