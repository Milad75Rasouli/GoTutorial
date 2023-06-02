package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var mmm = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

type Book struct {
	Name  string
	Price float32
}

func main() {
	// will result in a runtime error because
	// the a map hasn't been initialized before
	// assigning values to it. In Go, when you
	// declare a map using the var keyword,
	// it's always initialized with a zero
	// value of nil. This means that you cannot
	// assign values directly to its keys without
	// first initializing it. To fix this issue,
	// you need to initialize the map before adding
	// key-value pairs into it. Here's an example:
	//	var a map[int32]string

	var mm map[Book]string = map[Book]string{
		Book{"Name1", 12.3}: "description1",
		Book{"Name2", 2.29}: "description2",
	}

	fmt.Println(mm)

	fmt.Println(mmm)

	m := make(map[int32]string)

	m[1] = "Hello"
	//	a[1] = "World"
	// https://go.dev/tour/moretypes/22
	m[33] = "em"
	m[21] = "qwa"
	fmt.Printf("%v\n", m)
	//	fmt.Printf("%s\n", a[1])
	_, exists := m[33]
	if exists == true {
		delete(m, 33)
	}
	fmt.Printf("%v\n", m)
}
