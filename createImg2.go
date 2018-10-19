package main

import (
	"os"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"image/color"
)

func main() {
	const (
		dx = 640
		dy = 480
	)

	inputFileName := "./frame8"
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer, err := ioutil.ReadAll(inputFile)

	filename := "./dist/demo1.png"
	imgFile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return

	}
	defer imgFile.Close()

	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	offset := 0
	for y := 0; y < dy; y++ {
		for x := 0; x < dx*2; x += 4 {
			//Y  := buffer[x+offset]
			//Y2 := buffer[x+offset+2]
			//img.Set(x/2, y, color.RGBA{uint8(Y), uint8(Y), uint8(Y), 255})
			//img.Set(x/2+1, y, color.RGBA{uint8(Y2), uint8(Y2), uint8(Y2), 255})

			Y := buffer[x+offset]
			U := buffer[x+offset+1]
			Y2 := buffer[x+offset+2]
			V := buffer[x+offset+3]

			R, G, B := yuv2rgb(float32(Y), float32(U), float32(V))
			img.Set(x/2, y, color.RGBA{R, G, B, 255})
			R1, G1, B1 := yuv2rgb(float32(Y2), float32(U), float32(V))
			img.Set(x/2+1, y, color.RGBA{R1, G1, B1, 255})

			//B := 1.164*float32((Y - 16)) + 2.018*float32((U - 128));
			//R := 1.164*float32((Y - 16)) + 1.159*float32((V - 128));
			//G := 1.164*float32((Y - 16)) - 0.391*float32((U - 128)) - 0.813*float32((V - 128));
			//B2 := 1.164*float32((Y2 - 16)) + 2.018*float32((U - 128));
			//R2 := 1.164*float32((Y2 - 16)) + 1.159*float32((V - 128));
			//G2 := 1.164*float32((Y2 - 16)) - 0.391*float32((U - 128)) - 0.813*float32((V - 128));
			//img.Set(x/2, y, color.RGBA{getValue(R), getValue(G), getValue(B), 255})
			//img.Set(x/2+1, y, color.RGBA{getValue(R2), getValue(G2), getValue(B2), 255})

		}
		offset += dx * 2
	}

	err = png.Encode(imgFile, img)
	if err != nil {
		fmt.Println(err)
	}
}

func yuv2rgb(y, u, v float32) (r, g, b uint8) {
	r1 := y + (1.370705 * (v - 128));
	g1 := y - (0.698001 * (v - 128)) - (0.337633 * (u - 128));
	b1 := y + (1.732446 * (u - 128));
	if r1 > 255 {
		r1 = 255
	}
	if g1 > 255 {
		g1 = 255;
	}
	if b1 > 255 {
		b1 = 255;
	}
	if r1 < 0 {
		r1 = 0;
	}
	if g1 < 0 {
		g1 = 0;
	}
	if b1 < 0 {
		b1 = 0;
	}
	r = uint8(r1 * 220 / 256);
	g = uint8(g1 * 220 / 256);
	b = uint8(b1 * 220 / 256);
	return
}

func getValue(val float32) uint8 {
	if val > 255 {
		return 255
	}
	if val < 0 {
		return 0
	}
	return uint8(val)
}
