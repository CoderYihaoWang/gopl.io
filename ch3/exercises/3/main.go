// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _, aok := corner(i+1, j)
			bx, by, c, bok := corner(i, j)
			cx, cy, _, cok := corner(i, j+1)
			dx, dy, _, dok := corner(i+1, j+1)
			if aok && bok && cok && dok {
				fmt.Printf("<polygon fill='rgb(%d,%d,%d)' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					c.R, c.G, c.B,
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (sx float64, sy float64, c color.RGBA, ok bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)
	if !ok {
		return
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	cscale := (z+1)/2
	var r, g, b uint8
	r = uint8(255*cscale)
	if cscale <= 0.5 {
		g = uint8(512*cscale)
	} else {
		g = uint8(512-512*cscale)
	}
	b = uint8(255-255*cscale)
	c = color.RGBA{r, g, b, 255}
	return
}

func f(x, y float64) (z float64, ok bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	z = math.Sin(r) / r
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0.0, false
	}
	return z, true
}

//!-
