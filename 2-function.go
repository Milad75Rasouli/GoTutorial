package main

import (
	"fmt"
	"math"
)

func devid(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("a/b: b must not be 0.")
	}
	return a / b, nil

}

func my_pow(a, n float64) float64 {
	return math.Pow(a, n)
}
func print_something(text string) {
	fmt.Print(text)
}
func main() {
	print_something("Hello World\n")

	fmt.Print(add(2, 4), "\n")

	fmt.Print(my_pow(2, 4), "\n")

	result, err := devid(3, 2)
	if err == nil {
		fmt.Printf("result is %.2f", result)
	} else {
		fmt.Printf("%s", err)
	}
}

func add(a, b int32) (result int32) {
	result = a + b
	return
}
