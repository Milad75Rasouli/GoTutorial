package main

import "fmt"

func a(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Printf("int(%v)\n", i)
		break
	case float32:
		a(float64(i.(float32)))
		// break
		// ++++ If you omit the break statement after ++++
		// each case in your code, it will not cause any syntax errors
		// or compilation failures - rather, it will lead to unexpected
		// behavior at runtime.
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
func main() {
	a(20.1)
	a(float32(10.8))
	t := string("text1")
	a(t)
	b := []byte(t)
	a(b)
}
