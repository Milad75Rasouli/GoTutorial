package main

import (
	"fmt"
	"os"

	Abs "go.mod/packages"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR:main:%s\n", err)
		}
	}()

	a := Abs.Abs(-5.1)
	print(a.(float64))
}
