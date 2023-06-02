package main

import "fmt"

type Dimention struct {
	X int32
	Y int32
}

func main() {
	const number1 = 3.5
	var p *int32
	var number2 int32 = 40
	var array_numbers []int32 = []int32{
		3,
		55,
		10,
	}
	// it's incorrect
	//p = &number1

	p = &number2
	fmt.Printf("%v\n", *p)
	// it's illigal to do
	// p = array_numbers
	p = &array_numbers[0]
	fmt.Printf("%v\n", *(p))

	pp := &Dimention{
		34,
		11,
	}
	pp.X = 90
	fmt.Printf("%+v\n", *pp)

}
