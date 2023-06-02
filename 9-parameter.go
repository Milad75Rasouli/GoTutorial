package main

import (
	"fmt"
)

func DoubleIt(number *int32) {
	*number = *number + *number
}

func DoubleSlice(numbers []int32) {
	// this wont change the slice
	//	for number := range numbers {
	//		number = number * 2
	//	}

	// but this modify the slice
	size := len(numbers)
	for index := 0; index < size; index++ {
		numbers[index] *= 2
	}
}

// pass by value
func ReverseString(input string) string {
	var output string
	input_size := len(input) - 1
	for index := input_size; index >= 0; index-- {
		output = output + string(input[index])
		//fmt.Printf("%c", input[index])
	}
	return output
}

func main() {
	number := int32(2)
	numbers := []int32{3, 5, 8}
	sample_string := "Hello"

	DoubleIt(&number)

	DoubleSlice(numbers)

	result_string := ReverseString(sample_string)

	fmt.Printf("pass by pointer:%v\n", number)
	fmt.Printf("pass by slice:%v\n", numbers)
	fmt.Printf("pass by value:%v\n", result_string)
}
