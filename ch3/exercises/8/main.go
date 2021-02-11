// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
)

func main() {
	http.HandleFunc("/", makeHandler(mandelbrot128))
	http.HandleFunc("/mandelbrot64", makeHandler(mandelbrot64))
	http.HandleFunc("/mandelbrot128", makeHandler(mandelbrot128))

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func makeImage(f func(interface{}) color.Color) *image.RGBA {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, f(z))
		}
	}
	return img
}

func makeHandler(f func(interface{}) color.Color) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := png.Encode(w, makeImage(f))
		if err != nil {
			log.Fatalf("mandelbrot: %v", err)
		}
	}
}

func mandelbrot128(z interface{}) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	zz := z.(complex128)
	for n := uint8(0); n < iterations; n++ {
		v = v*v + zz
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrot64(z interface{}) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	zz := complex64(z.(complex128))
	for n := uint8(0); n < iterations; n++ {
		v = v*v + zz
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
