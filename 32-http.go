package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type PostData struct {
	Name   string `json:Name`
	Action string `json:Action`
	Count  int32  `json:Count`
}

func main() {
	baseURL := "https://webhook.site/481692e5-6b56-45f7-b165-b5ca77cfbd19"

	params := url.Values{}
	params.Set("ID", "3299")
	params.Set("Name", "Josh")

	get_url := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	// fmt.Println(get_url)
	res, err := http.Get(get_url)
	if err != nil {
		log.Fatalf("Get ERROR:%s\n", err.Error())
	}

	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
	fmt.Println("======================")
	user_post_data := &PostData{
		Name:   "Tom",
		Action: "studing books",
		Count:  3,
	}

	var result_post_data bytes.Buffer
	enc := json.NewEncoder(&result_post_data)
	if err := enc.Encode(user_post_data); err != nil {
		log.Fatal("ERROR: %s\n", err.Error())
	}
	// fmt.Printf("%s\n", result_post_data)
	post_response, err := http.Post(baseURL, "application/json", &result_post_data)
	if err != nil {
		log.Fatalf("ERROR:%s\n", err)
	}
	defer post_response.Body.Close()
	io.Copy(os.Stdout, post_response.Body)

}
