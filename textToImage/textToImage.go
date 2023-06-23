// package main
package texttoimage

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type imageSize struct {
	X int32
	Y int32
}

func AddLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func MakeImageWithText(size imageSize, text string, fileName string) error {
	img := image.NewRGBA(image.Rect(0, 0, int(size.X), int(size.Y)))

	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	AddLabel(img, 5, 10, text)
	fOutPutIMG, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("ERROR:%s", err.Error())
	}
	defer fOutPutIMG.Close()

	jpeg.Encode(fOutPutIMG, img, nil)

	return nil
}

// func main() {

// 	size := imageSize{100, 100}
// 	text := "Hello World!"
// 	outputFileName := "output.jpg"

// 	err := MakeImageWithText(size, text, outputFileName)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	fmt.Println("Image saved successfully!")
// }
