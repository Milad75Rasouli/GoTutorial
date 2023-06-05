package main

import "fmt"

const (
	MyFirst  = 2.2
	MySecond = 8.4
)

func add(a float32, b float32) (result float32) {
	result = a + b
	return
}

func main() {
	fmt.Println(MySecond)
	fmt.Println(MyFirst)
	fmt.Println(add(MyFirst, MySecond))
}
