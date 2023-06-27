// https://medium.com/@MTrax/golang-the-ultimate-guide-to-dependency-injection-4556b97f9cbd#4a6d
package main

import "fmt"

// BarInterface is a sample interface representing some external dependency
type BarInterface interface {
	DoSomething() error
}

// Bar is a sample implementation of the BarInterface interface
type Bar struct{}

func (b *Bar) DoSomething() error {
	fmt.Println("Do something with bar")
	return nil
}

// Foo is our main object that depends on the `Bar` type.
type Foo struct {
	bar BarInterface // private variable for storing the dependent instance.
}

// NewFoo creates a new instance of `Foo` and injects its dependencies through constructor injection.
func NewFoo(barInstance BarInterface) *Foo {
	return &Foo{bar: barInstance}
}

// SetBar sets the `Bar` dependency using setter injection.
func (f *Foo) SetBar(barInstance BarInterface) {
	f.bar = barInstance
}

// DoSomethingWithInjection uses method injection to pass our desired implementation of `Bar`.
func (f *Foo) DoSomethingWithInjection(barInstance BarInterface) error {
	// use provided instance of 'bar' here...
	if err := barInstance.DoSomething(); err != nil {
		return err
	}

	return nil
}

func main() {

	// Create an instance of `Bar`
	bar := &Bar{}

	/*
		Create an instance of `Foo` and inject its dependencies using constructor injection by passing our created 'bar'
		instance as argument to it.
	*/
	f1 := NewFoo(bar)
	fmt.Printf("%+v\n", f1)

	/*
	   Use setter injection to set/change the dependent value in existing Foo object when needed e.g. during runtime or when you want different implementations at different times.
	   In this case we are changing the instance of `bar` with a new implementation.
	*/
	bar2 := &Bar{}
	f1.SetBar(bar2)
	fmt.Printf("%+v\n", f1)

	/*
	   Use method injection to pass dependencies only when needed for specific functions/methods.
	   In this case we are passing our desired implementation of `Bar` through the function argument itself
	*/
	if err := f1.DoSomethingWithInjection(&Bar{}); err != nil {
		fmt.Println(err)
	}
}
