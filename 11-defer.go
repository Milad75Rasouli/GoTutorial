// The results are the following:
// Hello
// Cleaning up in the end
// Cleaning up at the beginning

/*
	INFORMATION ABOUT "defer"

We use the defer keyword in Go when
we want to ensure that a certain function is
called after the completion of another
function, regardless of whether an error
occurs or not. This is useful for tasks
such as clean-up operations like closing
files, network connections, or releasing
resources.
*/
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
