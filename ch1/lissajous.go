package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

/*[]color.Color{...} is a composite literal, a compact notation for instantiating any of Go's composite types from a sequence of element values
In this case a 'slice', a dynamically sized view into the elements of an array */
var palette = []color.Color{color.Black, color.White}

const (
	blackIndex = 0 // next color in palette
	whiteIndex = 1 // first color in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	// constants values are fixed at compile time
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	// gig.GIF is another composite literal, in this case a struct, similar to javascript object
	// Package gif implements a GIF image decoder and encoder.
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), whiteIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
