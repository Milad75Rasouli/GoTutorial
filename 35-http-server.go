package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloHandler(rw http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(rw, "Hello Dear!\nThis is my first try to make a http server.")
	if err != nil {
		fmt.Printf("ERROR:%s", err.Error())
	}
}

func main() {
	http.HandleFunc("/greeting", sayHelloHandler)

	if err := http.ListenAndServe(":5001", nil); err != nil {
		log.Fatal(err)
	}
}
