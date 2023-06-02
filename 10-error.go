package main

import (
	"fmt"
	"math"
)

func Pow(a float64, b float64) (float64, error) {
	if a < 0 || b < 0 {
		return 0.0, fmt.Errorf("Numbers must be greater than 0.")
	}
	result := math.Pow(a, b)
	return result, nil
}
func main() {
	number1, number2 := float64(-1), float64(3)
	result, error := Pow(number1, number2)
	if error == nil {
		fmt.Printf("%v^%v=%v\n", number1, number2, result)
	} else {
		fmt.Println(error)
	}
}
