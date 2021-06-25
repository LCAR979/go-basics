package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

/*
type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
*/
type Image struct {
	w, h int
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.w, im.h)
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
