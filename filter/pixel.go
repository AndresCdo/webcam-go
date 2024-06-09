// Description: This file contains the implementation of the pixel filter.
// The pixel filter is a simple filter that applies a filter to an image.
// The filter can be one of the following:
// - Grey: Converts the image to greyscale.
// - Invert: Inverts the colors of the image.
// - Noise: Adds random noise to the image.
// - Red: Sets the blue and green channels to 0, making the image red.
// The filter is applied to the image by iterating over the pixels and applying the filter to each pixel.

package filter

import (
	"math/rand"
	"time"
)

// Pixels is a byte array to represent image pixels.
type Pixels []byte

// MakeGrey applies a grey filter to pixels.
func (p Pixels) MakeGrey() {
	for i := 0; i < len(p); i += 4 {
		red := p[i]
		green := p[i+1]
		blue := p[i+2]

		grey := uint8(float64(red)*0.299 + float64(green)*0.587 + float64(blue)*0.114)
		p[i] = grey
		p[i+1] = grey
		p[i+2] = grey
	}
}

// Invert applies an invert filter to pixels.
func (p Pixels) Invert() {
	for i := 0; i < len(p); i += 4 {
		p[i] = 255 - p[i]
		p[i+1] = 255 - p[i+1]
		p[i+2] = 255 - p[i+2]
	}
}

// MakeNoise applies a random number to pixels.
func (p Pixels) MakeNoise() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(p); i += 4 {
		random := rand.Float64()
		p[i] = uint8(float64(p[i]) * random)
		p[i+1] = uint8(float64(p[i+1]) * random)
		p[i+2] = uint8(float64(p[i+2]) * random)
	}
}

// MakeRed applies a red filter to pixels by setting blue and green to 0.
func (p Pixels) MakeRed() {
	for i := 0; i < len(p); i += 4 {
		p[i+1] = 0
		p[i+2] = 0
	}
}
