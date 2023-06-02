package main

import "fmt"

// Another way to make a loop.
func MyFactorial(number int32) int32 {
	if number <= 0 {
		return 1
	}
	return number * MyFactorial(number-1)
}

func main() {

	for index := 0; index < 5; index++ {
		fmt.Printf("%b ", index)
	}
	fmt.Println("\n===============")

	var simple_array []float32 = []float32{
		3.5,
		1.4,
		9.91,
	}
	for index := range simple_array {
		item := simple_array[index]
		if item > 2 {
			fmt.Printf("%2.1f is greater that 2.\n", item)
		} else if item == 2 {
			fmt.Printf("%2.1f is 2.\n", item)
		} else {
			fmt.Printf("%2.1f is lower than 2.\n", item)
		}
	}
	fmt.Println("===============")
	counter := 10.5
	for counter > 0.0 {
		fmt.Printf("%f, ", counter)
		counter -= 1.1112
	}
	fmt.Println("\n===============")
	number := int32(4)
	fmt.Println(MyFactorial(number))
}
