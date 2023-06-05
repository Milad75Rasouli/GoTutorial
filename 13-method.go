package main

import "fmt"

type Vehicle struct {
	Name  string
	Color string
	Speed float64
}

// with *
func (vehicle *Vehicle) Drive(name string) (result string) {
	vehicle.Name = name
	result = vehicle.Name + " " + vehicle.Color + " " + fmt.Sprint(vehicle.Speed)
	return
}

// without *
func (vehicle Vehicle) Drive2(name string) (result string) {
	vehicle.Name = name
	result = vehicle.Name + " " + vehicle.Color + " " + fmt.Sprint(vehicle.Speed)
	return
}

func main() {
	var car1 Vehicle = Vehicle{Color: "Blue", Speed: 220}
	result := car1.Drive("BMW")
	fmt.Println(result)
	result = car1.Drive2("Benz")
	fmt.Println(result)
}
