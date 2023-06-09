package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "text.txt"
	var text []byte = []byte("a simle\ntext\nthis is just\nfor test")
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		//open text.txt: The system cannot find the file specified.
		os.Exit(1)
	}
	file.Write(text)

}
