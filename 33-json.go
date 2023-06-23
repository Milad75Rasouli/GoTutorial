package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// Note: The json message must end without ','
var input_data = `
{
	"Model":"Ford",
	"Hight":[1.32, 1.5],
	"Weight":1000.4,
	"Owner":"Milad"
}
`

type input_json_type struct {
	Owner  string    `json:"Owner"`
	Model  string    `json:"Model"`
	Hight  []float64 `json:"Hight"`
	Weight float64   `json:"Weight"`
}

func main() {
	// var ii = 34
	// var str = "string"
	// var ff = 44.1
	// fmt.Println(ii, str, ff)
	reader := bytes.NewBufferString(input_data)
	decoded_input_data := json.NewDecoder(reader)
	// use struct to gain the message
	parsesd_input_data := &input_json_type{}

	if err := decoded_input_data.Decode(parsesd_input_data); err != nil {
		log.Fatalf("ERROR:%v\n", err)
	}

	// Note: you cant assign any parameter to them
	// parsesd_input_data = 90.2

	// fmt.Println(parsesd_input_data.Weight)
	fmt.Printf("Parsed data is:%+v\n", parsesd_input_data)

	// creating a json message
	fmt.Println("=================")

	// use map for creating a json message
	json_data := make(map[string]interface{})
	json_data["Model"] = "Tesla"
	json_data["Owner"] = []string{"Joe", "Tom"}
	json_data["Hight"] = 1.2
	json_data["Weight"] = 1221.4

	// json_encoder := json.NewEncoder(os.Stdout)
	var json_buffer bytes.Buffer
	json_encoder := json.NewEncoder(&json_buffer)
	if err := json_encoder.Encode(json_data); err != nil {
		log.Fatalf("Error: unable to encode %s", err.Error())
	}
	fmt.Printf("%s\n", json_buffer.String())

}
