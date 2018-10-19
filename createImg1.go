package main

import (
	"os"
	"fmt"
	"image"
	"image/color"
	"image/png"
)

func main() {
	const (
		dx = 300
		dy = 300
	)

	filename := "./dist/demo1.png"

	imgFile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return

	}
	defer imgFile.Close()

	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			if (x+1)%30 == 0 || x == 0 || (y+1)%30 == 0 || y == 0  {
				img.Set(x, y, color.RGBA{uint8(x % 256), uint8(y % 256), 0, 255})
			}else{
				img.Set(x, y, color.RGBA{100,100,0,255})
			}
		}
	}

	err = png.Encode(imgFile, img)
	if err != nil {
		fmt.Println(err)
	}
}
