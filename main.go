package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}

func main() {
	// size of image x and y
	nx := 256
	ny := 256

	const ccolor = 255.999

	f, err := os.Create("out.ppm")

	defer f.Close()

	check(err, "Error opening file: %v\n")

	// http://netpbm.sourceforge.net/doc/ppm.html
	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", nx, ny)

	check(err, "Error writting to file: %v\n")

	pngFile, err := os.Create("out.png")
	check(err, "Error creating PNG file: %v\n")
	defer pngFile.Close()

	img := image.NewRGBA(image.Rect(0, 0, nx, ny))

	// writes each pixel with r/g/b values
	// from top left to bottom right
	for j := 0; j < ny; j++ {
		fmt.Printf("Lines remaining: %v\n", (ny - j))
		for i := 0; i < nx; i++ {
			// red and green values range from
			// 0.0 to 1.0
			r := float64(i) / float64(nx)
			g := float64(j) / float64(ny)
			b := 0.0

			// get intensity of colors
			ir := uint8(ccolor * r)
			ig := uint8(ccolor * g)
			ib := uint8(ccolor * b)

			img.Set(i, j, color.RGBA{ir, ig, ib, 255})

			_, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)

			check(err, "Error writting to file: %v\n")
		}
	}
	err = png.Encode(pngFile, img)
	check(err, "Error encoding PNG: %v\n")

}
