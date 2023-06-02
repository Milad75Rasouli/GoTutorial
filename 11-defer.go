// The results are the following:
// Hello
// Cleaning up in the end
// Cleaning up at the beginning

package main

import "fmt"

func Cleanup(name string) {
	fmt.Printf("Cleaning up %v\n", name)
}

func DoSomething() {
	defer Cleanup("at the beginning")
	fmt.Println("Hello")
	defer Cleanup("in the end")
}

func main() {

	DoSomething()
}
