package main

import (
	"fmt"
	"os"
)

func safeValue(vals []float32, index int32) float32 {
	// it's an ananimous function
	defer func() {
		if err := recover(); err != nil {
			// here is catch the error
			fmt.Fprintf(os.Stderr, "ERROR:%s\n", err)
		}
	}()

	return vals[index]
}
func main() {
	// #1
	// i := make([]int, 3)
	// i = []int{3, 5, 1}
	// // panic
	// fmt.Println(i[5])

	// #2
	// file, err := os.Open("file.file")
	// if err != nil {
	// 	// useing panic is discourage
	// 	panic(err) // = throw
	// }
	// defer file.Close()
	// fmt.Println(file.Name())

	v := make([]float32, 2)
	v = []float32{2.2, 1.1}
	result := safeValue(v, 4)
	fmt.Println(result)
}
