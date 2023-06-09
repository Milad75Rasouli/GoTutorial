package main

import "fmt"

type NilInterface interface {
	Method1()
}

func main() {
	var nilI NilInterface
	fmt.Printf("(%v,%T)\n", nilI, nilI)
	//(<nil>,<nil>)

	// Calling a method on a nil interface is a run-time error
	// because there is no type inside the interface tuple to
	// indicate which concrete method to call.
	// nilI.Method1()

	// Empty interfaces are used by code that handles
	// values of unknown type.
	var u interface{}
	fmt.Printf("(%v,%T)\n", u, u)

	var r int32 = int32(410)
	u = r
	fmt.Printf("(%v,%T)\n", u, u)

	// b := string{"salma"} is surly wrong
	b := string("salma")
	// var b string = "Message"
	u = b
	fmt.Printf("(%v,%T)\n", u, u)

}
