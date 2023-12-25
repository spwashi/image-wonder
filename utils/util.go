package utils

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"math"
)

func SimpleImage() *image.RGBA {
	// generate an image
	img := image.NewRGBA(image.Rect(
		0,
		0,
		1000,
		1000,
	))

	// list of color options
	colors := []color.RGBA{
		{255, 0, 0, 255},
		{0, 255, 0, 255},
		{0, 0, 255, 255},
		{255, 255, 0, 255},
		{0, 255, 255, 255},
	}

	// fill the image with a color
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if (x+y)%13 == 0 {
				img.Set(x, y, colors[(x+y)%len(colors)])
			} else {
				pixelColor := color.RGBA{
					uint8(math.Sqrt(float64(y))),
					uint8(math.Sqrt(float64(x))),
					255,
					uint8(math.Sqrt(float64(x*x + y*y))),
				}
				img.Set(x, y, pixelColor)
			}
		}
	}

	return img
}

func NewOpenCvWindow() {
	webcam, err := gocv.OpenVideoCapture(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	if err != nil {
		fmt.Sprintf("error opening webcam: %s\n", err)
		return
	}
	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
