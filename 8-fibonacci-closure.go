package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var previous int = 0
	var next int = -1
	fn := func() (result int) {
		if next < 1 {
			next++
			return next
		}
		result = next + previous
		previous = next
		next = result
		//previous++
		return
	}
	return fn
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
