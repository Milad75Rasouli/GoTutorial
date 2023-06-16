package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	// "github.com/pkg/errors"
)

func GetDir(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, errors.Wrap(err, "this is my error message that has been wraped")
		// return nil, err
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
	files, err := GetDir("/33")
	if err != nil {
		// write the error on stderr
		fmt.Fprintf(os.Stderr, "%v", err)
	} else {
		// write the files on stdout
		fmt.Fprintf(os.Stdout, "%v", files)
	}
}
