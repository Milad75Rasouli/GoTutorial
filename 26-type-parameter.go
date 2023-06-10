package main

import "fmt"

func At[T comparable](m []T, i int64) (result T, err error) {
	err = nil
	if int(i) >= len(m) {
		err = fmt.Errorf("no suitable index for the slice")
		return
	}
	result = m[i]
	return
}

func main() {

	r := []string{"dfc", "D1", "yt"}
	a, err := At[string](r, 7)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}
}
