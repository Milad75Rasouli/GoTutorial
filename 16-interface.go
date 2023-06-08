package main

import "fmt"

type Fruits interface {
	Energy() int64
}

func main() {
	show := func(f Fruits) {
		fmt.Printf("The fruit has %d energy.\n", f.Energy())
	}

	anApple := Apple{20, 10}
	aBenana := Banana{20, 10, 5}

	show(&anApple)
	show(&aBenana)

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
