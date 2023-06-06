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

// none-struct
type MyFloat float32

func (f MyFloat) Squer() float32 {
	result := float32(f * f)
	return result
}

func main() {
	var car1 Vehicle = Vehicle{Color: "Blue", Speed: 220}
	result := car1.Drive("BMW")
	fmt.Println(result)
	result = car1.Drive2("Benz")
	fmt.Println(result)

	f := MyFloat(3.14)
	fmt.Printf("%v\n", f.Squer())

}
