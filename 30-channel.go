package main

import (
	"fmt"
	"time"
)

// there are 2 kinds of channel. Buffred and Unbuffred

func main() {

	// ch := make(chan string)
	// the below line causes deadlock! fatal error
	// <-ch

	ch := make(chan string)
	go func(inner_ch chan string) {
		inner_ch <- "a text from the first function."

		// to send multiple values we cant use the following form.
		// inner_ch <- "a text from the first function."
		// inner_ch <- "a text from the first function."
		// inner_ch <- "a text from the first function."
	}(ch)

	returned_value := <-ch

	fmt.Printf("Result is \"%s\"\n", returned_value)

	// send multiple values through a channel
	ch2 := make(chan int64)

	// send some data to the ch2
	go func(in_ch chan int64) {
		for i := int64(1); i < 3; i++ {
			fmt.Printf("send %d\n", i)
			in_ch <- i
			// it makes the thread wait for 1 second
			time.Sleep(1 * time.Second)
		}
	}(ch2)

	var val int64
	for i := 1; i < 3; i++ {
		val = <-ch2
		fmt.Printf("get %d\n", val)
	}

	// so far we came to the conclution that channel are Unbuffred
	// and to send multipal values through them we need to use for loop not(a time to sleep).

	// close to signal it's done
	// send some data to the ch2
	go func(in_ch chan int64) {
		for i := int64(1); i < 10; i++ {
			fmt.Printf("2-send %d\n", i)
			in_ch <- i
			// it makes the thread wait for 1 second
			// time.Sleep(1 * time.Second)
		}
		// in_ch.Close()
		close(ch2)
	}(ch2)

	// here we dont need to know how much data do we have
	// because we closes ch2 after we are done
	for v := range ch2 {
		fmt.Printf("2-get %d\n", v)
	}
}
