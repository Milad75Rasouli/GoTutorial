package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {

	pic := image.NewRGBA(image.Rect(0, 0, 200, 300))
	myColor := color.RGBA{200, 100, 200, 255} // just need to set your fav. color here
	for x := 0; x < pic.Bounds().Size().X; x++ {
		for y := 0; y < pic.Bounds().Size().Y; y++ {
			pic.Set(x, y, myColor)
		}
	}

	file, err := os.Create("pic.png")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err2 := png.Encode(file, pic)
	if err2 != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("your picture is avaiable now.")
}
