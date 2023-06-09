package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	var text string
	text = `this is my message.
that i'm writing it.
because ima do something 
with it.`
	reader := strings.NewReader(text)

	chunk := make([]byte, 8)

	for {
		n, err := reader.Read(chunk)
		fmt.Printf("%v\n", chunk)
		fmt.Printf("%q\n", chunk[:n])
		if err == io.EOF {
			break
		}
	}
}
