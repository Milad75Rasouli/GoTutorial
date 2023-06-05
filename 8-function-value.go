package main

import (
	"fmt"
	"math"
)

func do(math func(float64, float64) float64) (result float64) {

	fmt.Println("funtion is running!")
	result = math(3, 5)
	return
}

func main() {
	result := do(math.Max)
	fmt.Printf("%v\n", result)

	add := func(a float64, b float64) (result float64) {
		result = a + b
		return
	}

	result2 := do(add)
	fmt.Printf("%v\n", result2)

}
