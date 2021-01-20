// Modify Lissajous server to accept input from url
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type config struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0xbb, 0xff, 0xff, 0xff},
	color.RGBA{0xdd, 0xdd, 0xff, 0xff},
	color.RGBA{0xff, 0xbb, 0xff, 0xff},
	color.RGBA{0xff, 0xdd, 0xdd, 0xff},
	color.RGBA{0xff, 0xff, 0xbb, 0xff},
	color.RGBA{0xdd, 0xff, 0xdd, 0xff},
}

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	handler := func(w http.ResponseWriter, r *http.Request) {
		config := config{5, 0.001, 100, 64, 8}

		if err := r.ParseForm(); err == nil {
			for query, value := range r.Form {
				switch query {
				case "cycles":
					if v, err := strconv.Atoi(value[0]); err == nil {
						if v > 0 {
							config.cycles = v
						}
					}
				case "res":
					if v, err := strconv.ParseFloat(value[0], 64); err == nil {
						config.res = v
					}
				case "size":
					if v, err := strconv.Atoi(value[0]); err == nil {
						if v > 0 {
							config.size = v
						}
					}
				case "nframes":
					if v, err := strconv.Atoi(value[0]); err == nil {
						if v > 0 {
							config.nframes = v
						}
					}
				case "delay":
					if v, err := strconv.Atoi(value[0]); err == nil {
						if v > 0 {
							config.delay = v
						}
					}
				}
			}
		}

		lissajous(w, config)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, config config) {
	cycles := config.cycles   // number of complete x oscillator revolutions
	res := config.res         // angular resolution
	size := config.size       // image canvas covers [-size..+size]
	nframes := config.nframes // number of animation frames
	delay := config.delay     // delay between frames in 10ms units

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			xIndex := size + int(x*float64(size)+0.5)
			yIndex := size + int(y*float64(size)+0.5)
			img.SetColorIndex(xIndex, yIndex,
				uint8(yIndex/32%(len(palette)-1))+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
