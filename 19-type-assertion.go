package main

import "fmt"

func main() {
	var r interface{}

	q, ok := r.(float32)
	fmt.Println(q, ok)

	r = string("Text1")

	w, ok := r.(string)
	fmt.Println(w, ok)

	//fmt.Println(r.(int32)) // panic

}
