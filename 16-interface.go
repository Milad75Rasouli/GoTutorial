package main

import "fmt"

type Fruits interface {
	Energy() int64
}

func main() {
	show := func(f []Fruits) {
		sum := int64(0)
		for _, energy := range f {
			sum += energy.Energy()
		}
		fmt.Printf("The fruit has %d energy.\n", sum)
	}

	myApple := []Apple{{20, 10}, {90, 10}}
	myBenana := []Banana{{20, 10, 5}, {99, 18, 21}}

	fmt.Println("Apples energy:")
	show([]Fruits{&myApple[0], &myApple[1]})
	fmt.Println("Benanas energy:")
	show([]Fruits{&myBenana[0], &myBenana[1]})

}

// Fruits
// Note:
// if you write the name of elements in lower case,
// they will be inaccessible when they're initialized.
type Apple struct {
	Weight int64
	Sugar  int64
}

type Banana struct {
	Weight int64
	Length int64
	Sugar  int64
}

func (a *Apple) Energy() int64 {
	return a.Sugar * a.Weight
}

func (b *Banana) Energy() int64 {
	return b.Weight * b.Length * b.Sugar
}
