package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

func main() {
	ch1, ch2 := make(chan float64), make(chan float64)

	go func(ch1 chan float64, ch2 chan float64) {
		ch1 <- math.Sqrt(88)
	}(ch1, ch2)

	// by select we are able to select between 2 or more channels easily
	select {
	case sqrt_val := <-ch1:
		fmt.Printf("%f\n", sqrt_val)
	case pow_val := <-ch2:
		fmt.Printf("%f\n", pow_val)
	}

	fmt.Println("===================")

	body := make(chan string)
	err := make(chan error)
	url := "https://go.dev/tour/basics/3"
	go func(url string, body chan string, err chan error) {
		responce, err1 := http.Get(url)
		// defer responce.Body.Close()
		if err1 != nil {
			err <- fmt.Errorf("ERROR:%v", err1)
			return // missed
		}
		defer responce.Body.Close()

		body_content, err2 := ioutil.ReadAll(responce.Body)
		if err2 != nil {
			err <- fmt.Errorf("ERROR:%v", err1)
			return // missed
		}

		body <- string(body_content)

	}(url, body, err)

	select {
	case content := <-body:
		fmt.Printf("%s: %+v\n", url, content)
	case err_message := <-err:
		fmt.Printf("ERROR:%v\n", err_message)
	case <-time.After(1000 * time.Millisecond): // in fact it is a channal
		fmt.Printf("Timeout: unable to get %s\n", url)
	}

}
