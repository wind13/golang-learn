package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) At(x, y int) color.Color {
	v := calc(x, y)
	return color.RGBA{v, v, 255, 255}
}

func calc(x, y int) uint8 {
	return uint8((x*x + y*y))
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}
