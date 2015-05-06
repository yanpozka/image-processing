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

	image_result := image.NewGray(img.Bounds())
	var count int64 = 0

	// fmt.Println(img.At(-1, 0) : we can access with negative index
	for i := 0; i < H; i++ {
		for k := 0; k < W; k++ {
			if applyMatrixPixel(img, image_result, i, k) {
				count++
			}
		}
	}
	toimgpng, _ := os.Create("../../new.png")
	defer toimgpng.Close()

	if err := png.Encode(toimgpng, image_result); err != nil {
		fmt.Println("[-] Error trying to create the new image.", err)
	} else {
		fmt.Println("[+] Image created", toimgpng.Name())
	}
	fmt.Println(count)
}

//
// Matrix of convolution is:
//  1  1  1
//  1 -4  1
//  1  1  1
//
// TODO: Refactor a little bit to allow a custom matrix instead of that fixed one
//
func applyMatrixPixel(src image.Image, dst *image.Gray, r, c int) bool {
	pixel_value := getPixInt(src.At(r, c))

	// border pixels
	sum := getPixInt(src.At(r, c-1)) + getPixInt(src.At(r-1, c)) +
		getPixInt(src.At(r, c+1)) + getPixInt(src.At(r+1, c))

	// center pixel
	result := sum + (pixel_value * -4)

	var ok bool = true
	new_color := generateWhiteColor()
	if result < 0 {
		ok = false
	} else {
		new_color = getColorFromPixelInt(result)
	}
	dst.Set(r, c, new_color)
	return ok
}

//
func generateWhiteColor() color.Color {
	return color.RGBA{R: 255, G: 255, B: 255}
}

//
func getColorFromPixelInt(n int32) color.Color {
	var num uint32 = uint32(n)
	blue := uint8(num)
	num >>= 8
	green := uint8(num)
	num >>= 8
	red := uint8(num)

	return color.RGBA{R: red, G: green, B: blue}
}

// integer = red + green + blue (in that order)
func getPixInt(pixel color.Color) int32 {
	var r, g, b, _ uint32 = pixel.RGBA()
	red, green, blue := uint32(r>>8), uint32(g>>8), uint32(b>>8)
	var numpixel uint32 = red
	numpixel <<= 8
	numpixel |= green
	numpixel <<= 8
	numpixel |= blue

	return int32(numpixel)
}
