package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

type MM struct {
	f float64
}

func (m *MM) M() {
	fmt.Println(m.f)
}

type Do struct {
	Q I
}

func (d *Do) SetQ(i I) {
	d.Q = i
}
func (d *Do) DoIt() {
	d.Q.M()
}
func main() {
	var i I
	i = T{"hello"}
	i.M()
	i = &MM{23.44}
	i.M()

	fmt.Println("===========")
	var e I = T{"salam salam"}
	d := Do{e}
	d.DoIt()
	var y I
	y = &MM{10.1}
	d.SetQ(y)
	d.DoIt()

}
