package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	url := "https://echo.hoppscotch.io/text=Hi_Bunny"

	responce, err := http.Get(url)
	if err != nil {
		// fmt.Errorf("ERROR:%s\n", err.Error())
		// return
		// we can use Fatalf instead
		log.Fatalf("ERROR:%s\n", err.Error())
	}
	defer responce.Body.Close()

	io.Copy(os.Stdout, responce.Body)
}
