package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)

	b := []int32{6, 3, 1, 8, 10, 31, 2e3}
	fmt.Println(b)

	// SCLICE
	// Slices are like references to arrays:
	// A slice does not store any data,
	// it just describes a section of an
	// underlying array.
	fmt.Println("================")
	s := b[:len(b)-1]
	s2 := a[0:1]
	fmt.Println(s)
	fmt.Println(s2)
	s2[0] = "xxx"
	fmt.Println(a[0])
	aa := make([]float32, 3)
	fmt.Printf("len=%d cap=%d\n",
		len(aa), cap(aa))
	fmt.Println(aa)

	fmt.Println("================")
	xx := [][]string{
		[]string{"x", "x", "x"},
		[]string{"x", "o", "x"},
		[]string{"x", "o", "x"},
	}
	xx = append(xx, []string{"x", "x", "x"})
	fmt.Println(xx)
}
