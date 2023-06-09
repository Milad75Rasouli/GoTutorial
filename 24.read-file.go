package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "text.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		//open text.txt: The system cannot find the file specified.
		os.Exit(1)
	}
	fmt.Printf("%s", file)
}
