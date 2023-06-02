package main

import "fmt"

type Car struct {
	Name     string
	MaxSpeed int32
	Weight   float32
	Height   float32
	Price    float64
}

// it's possiable to declare here.
var car1 Car

func main() {

	car2 := Car{"BMW", 220, 1200.5, 150.1, 10000}

	car3 := Car{
		Price:    30000.2,
		Name:     "Ford",
		MaxSpeed: 250,
		Height:   120.5,
		Weight:   900.8,
	}

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	fmt.Printf("%+v\n", car1)
	fmt.Printf("%+v\n", car2)
	fmt.Printf("%v\n", car3)
}
