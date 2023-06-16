package Abs

import "testing"

func TestAbs(t *testing.T) {
	val := Abs(-3)
	if val.(int) < 0 {
		t.Fatalf("error in calculation")
	}
}
