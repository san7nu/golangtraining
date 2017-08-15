package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 1024, 1024))

	fmt.Printf("%T \n", m)
	m.Set(10, 10, color.White)
	m.Set(20, 20, color.White)
	png.Encode(os.Stdout, m)
}
