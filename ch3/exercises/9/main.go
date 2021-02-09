// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", makeHandler(mandelbrot))
	http.HandleFunc("/mandelbrot", makeHandler(mandelbrot))
	http.HandleFunc("/acos", makeHandler(acos))
	http.HandleFunc("/sqrt", makeHandler(sqrt))
	http.HandleFunc("/newton", makeHandler(newton))

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func makeImage(x, y int, zoom float64, f func(complex128) color.Color) *image.RGBA {
	xmin, ymin, xmax, ymax := -2.0, -2.0, +2.0, +2.0
	width, height          := float64(x), float64(y)
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for py := 0; py < int(height); py++ {
		yy := float64(py)/height*(ymax-ymin)/zoom + ymin/zoom
		for px := 0; px < int(width); px++ {
			xx := float64(px)/width*(xmax-xmin)/zoom + xmin/zoom
			z := complex(xx, yy)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, f(z))
		}
	}
	return img
}

func makeHandler(f func(complex128) color.Color) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		var x, y int
		var zoom float64

		xs := q.Get("x")
		x, err := strconv.Atoi(xs)
		if err != nil || xs == "" || x < 0 {
			x = 1024
		}

		ys := q.Get("y")
		y, err = strconv.Atoi(ys)
		if err != nil || ys == "" || y < 0 {
			y = 1024
		}

		zs := q.Get("zoom")
		zoom, err = strconv.ParseFloat(zs, 64)
		if err != nil || zs == "" {
			zoom = 1.0
		}

		err = png.Encode(w, makeImage(x, y, zoom, f))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
