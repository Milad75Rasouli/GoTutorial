package main

import (
	"fmt"
)

func SayHello(name string) string {
	return "Hi " + name
}

func main() {
	fmt.Println(SayHello(("Milad")))
}
