// handler echoes the HTTP request.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"io"
	"image/gif"
	"image"
	"math/rand"
	"math"
	"image/color"
)

var mu sync.Mutex
var count int

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	count++
	mu.Unlock()
	lissajous(w)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer)  {
	const (
		cycles = 5  // number of complete x oscillator revolutions
		res = 0.001 // angular resolution
		size = 100  // image canves convers [-size..+size]
		nframes = 64    // number of animation frames
		delay = 10   // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0    // relative frequency of y oscillator
	anim := gif.GIF{LoopCount:nframes}
	phase := 0.0    // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}