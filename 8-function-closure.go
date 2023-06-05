package main

import "fmt"

func funcMaker() func(int32) int32 {
	var value int32 = 0
	fn := func(a int32) int32 {
		value += a
		return value
	}
	return fn
}

func main() {

	x, y := funcMaker(), funcMaker()

	for i := int32(10); i > 0; i-- {
		a := x(i*2 + 1)  // 21
		b := y(-i*3 + 3) // -27
		fmt.Println(a, b)
	}
	// Each closure is bound to its own "value" variable.
	// The results:
	// 21 -27
	// 40 -51
	// 57 -72
	// 72 -90
	// 85 -105
	// 96 -117
	// 105 -126
	// 112 -132
	// 117 -135
	// 120 -135

}
