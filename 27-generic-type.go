package main

import "fmt"

type List[T any] struct {
	list []T
	size int64
}

func (l *List[T]) Push(a interface{}) {
	l.list = append(l.list, a.(T))
	l.size++
}
func (l *List[T]) Pop() (result T, err error) {
	err = nil
	if l.size <= 0 {
		err = fmt.Errorf("Pop can be invoked when the size is 0.")
		return
	} else {
		l.size--
		result = l.list[l.size]
		// here the end element must be deleted
		l.list = l.list[:l.size]
		return
	}
}
func main() {
	e := List[string]{}
	e.Push("Hello")
	e.Push("World!")

	val1, err1 := e.Pop()
	e.Push("!!!!!!")
	val2, err2 := e.Pop()
	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
	}
	fmt.Printf("%v, %v\n", val1, val2)
}
