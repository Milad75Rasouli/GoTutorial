package texttoimage

import (
	"fmt"
	"log"
	"os"
	"testing"
)

// $ go test
// $ go test -v
func TestMakeImageWithText(t *testing.T) {
	filename := "output.jpg"
	size := imageSize{100, 100}
	err := MakeImageWithText(size, "it's a test", filename)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(filename)
	if err != nil {
		log.Fatalf("Can't remove - %s", err)
	}
}

// Benchmarking
// $ go test -v -bench .
// $ go test -v -bench . -run A

// Profiling
// $ go test -v -bench . -run AA -cpuprofile=prof
// $ go tool pprof .\prof
// list MakeImageWithText
func BenchmarkMakeImageWithText(b *testing.B) {
	size := imageSize{100, 100}
	for i := 0; i < b.N; i++ {
		filename := fmt.Sprintf("o%d.jpg", i)
		err := MakeImageWithText(size, "it's a test", filename)
		if err != nil {
			log.Fatal(err)
		}
		err = os.Remove(filename)
		if err != nil {
			log.Fatalf("Can't remove - %s", err)
		}
	}
}
