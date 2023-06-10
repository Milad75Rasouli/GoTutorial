package main

import "fmt"

func At[T comparable](m []T, i int64) (result T) {
	result = m[i]
	return
}

func main() {

	r := []string{"dfc", "D1", "yt"}
	a := At[string](r, 2)

	fmt.Println(a)
}
