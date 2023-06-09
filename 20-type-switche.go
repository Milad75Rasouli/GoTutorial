package main

import "fmt"

func main() {
	a := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Printf("int(%v)\n", i)
			break
		case float32:
		case float64:
			fmt.Printf("float(%v)\n", i)
			break
		case string:
			fmt.Printf("string(%v)\n", i)
			break
		default:
			fmt.Printf("unkowun(%v)\n", i)
			fmt.Printf("%T\n", i)
			break
		}
	}
	// a(20.1)
	a(float32(10))
	// a("text1")
}
