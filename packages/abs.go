// package main
package Abs

import (
	"fmt"
)

func Abs(number interface{}) (result interface{}) {
	result = number
	switch number.(type) {
	case int:
		if number.(int) < 0 {
			result = -number.(int)
		}
		break
	case float32:
		if number.(float32) < 0 {
			result = -number.(float32)
		}
		break
	case float64:
		if number.(float64) < 0 {
			result = -number.(float64)
		}
		break
	default:
		fmt.Errorf("the input parameter must be int or float.")
		break

	}
	return
}
