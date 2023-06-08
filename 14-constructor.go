package main

import "fmt"

type Sum struct {
	Number int64
	Index  int64
}

func (sum *Sum) ImportNumber(number int64) {
	sum.Number += number
	sum.Index++
}
func (sum Sum) Get() int64 {
	return sum.Number
}
func (sum *Sum) Avrage() (result float64) {
	result = float64(sum.Number / sum.Index)
	return result
}

func NewSum(initializerNumber int64) (*Sum, error) {
	if initializerNumber < 0 {
		return nil, fmt.Errorf("the number must be greather than 0.")
	}
	sum := &Sum{Number: initializerNumber}
	return sum, nil
}
func main() {
	sum, err := NewSum(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for index := int64(10); index > 0; index-- {
		sum.ImportNumber(index*3 + 2)
	}

	// func show() {
	// 	fmt.Printf("Number is %v\n", sum.Get())
	// 	fmt.Printf("Avrage is %v\n", sum.Avrage())
	// }
	// show()

	show := func() {
		fmt.Printf("Number is %v\n", sum.Get())
		fmt.Printf("Avrage is %v\n", sum.Avrage())
	}
	show()
}
