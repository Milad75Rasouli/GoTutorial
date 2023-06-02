package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var item int32 = 10
	if item > 10 {
		fmt.Print("it's greater than 10.\n")
	} else if item < 10 {
		fmt.Print("it's lower than 10.\n")

	} else {
		fmt.Print("it's 10.\n")
	}
	fmt.Print("===================\n")

	hour := time.Now()
	switch {
	case hour.Hour() < 12:
		fmt.Print("Good morning fella!\n")

	case hour.Hour() < 17:
		fmt.Print("Good afternoon mate!\n")

	default:
		fmt.Print("Good evening pal!\n")
	}
	fmt.Print("===================\n")

	fmt.Print("Go runs on ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	fmt.Print("===================\n")
	now := time.Now().Weekday()
	switch now {
	case time.Saturday:
		fmt.Println("Saturday")
	case time.Sunday:
		fmt.Println("Sunday")
	case time.Monday:
		fmt.Println("Monday")
	case time.Tuesday:
		fmt.Println("Tuesday")
	case time.Wednesday:
		fmt.Println("Wednesday")
	case time.Thursday:
		fmt.Println("Thursday")
	case time.Friday:
		fmt.Println("Friday")
	}
}
