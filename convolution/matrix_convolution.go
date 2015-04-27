package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	file_img, errfi := os.Open("../Lenna.png")
	if errfi != nil {
		panic("[-] Error fatal. os.Open")
	}
	defer file_img.Close()

	img, err := png.Decode(bufio.NewReader(file_img))
	if err != nil {
		panic("[-] Error trying to create a jpeg.Decode.")
	}
	s := img.Bounds().Size()
	var W, H int = s.X, s.Y
	fmt.Println(W, H)

	image_result := image.NewRGBA(image.Rect(0, 0, W, H))

	// fmt.Println(img.At(-1, 0) : we can access with negative index
	for i := 0; i < H; i++ {
		for k := 0; k < W; k++ {
			applyMatrixPixel(img, image_result, i, k)
		}
	}
	toimgpng, _ := os.Create("new.png")
	defer toimgpng.Close()

	if err := png.Encode(toimgpng, image_result); err != nil {
		fmt.Println("[-] Error trying to create the new image.", err)
	} else {
		fmt.Println("[+] Image created", toimgpng.Name())
	}
}

//
// By the moment only displace the image 5 pixels up left
//
func applyMatrixPixel(src image.Image, dst *image.RGBA, r, c int) {
	// pixel_value := getPixInt(src.At(r, c))
	// fmt.Println(pixel_value)

	// TODO: transform pixel after calculation

	dst.Set(r, c, src.At(r+5, c+5))
}

//
func getPixInt(pixel color.Color) uint32 {
	var r, g, b, _ uint32 = pixel.RGBA()
	red, green, blue := uint32(r>>8), uint32(g>>8), uint32(b>>8)
	var numpixel uint32 = red
	numpixel <<= 8
	numpixel |= green
	numpixel <<= 8
	numpixel |= blue

	return numpixel
}
