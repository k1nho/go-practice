package main

import(
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	height int
	width int
	color uint8
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0,0,im.width, im.height)
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) At(x int, y int) color.Color {
	return color.RGBA{im.color + uint8(x), im.color + uint8(y), 255, 255}
} 

func main() {
	m := Image{244, 244, 255}
	pic.ShowImage(m)
}
