package main

import (
	"fmt"
	"net/http"
	"sync"
)

func getHearder(url string) {
	response, _err := http.Get(url)
	if _err != nil {
		fmt.Printf("Error:%v\n", _err)
		return
	}
	defer response.Body.Close()
	// defer response.Close()
	ctype := response.Header.Get("Content-Type")
	fmt.Printf("%s\n", ctype)

}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			getHearder(url)
			wg.Done()
		}(url)
	}
	wg.Wait()

}
