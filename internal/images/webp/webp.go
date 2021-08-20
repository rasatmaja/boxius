package webp

import (
	"image"
	"io"

	"github.com/chai2010/webp"
)

// Decode reads a WebP image from r and returns it as an image.Image.
func Decode(r io.Reader) (image.Image, error) {
	return webp.Decode(r)
}

// Options are the encoding parameters.
// Quality ranges from 1 to 100 inclusive, higher is better.
type Options struct {
	Quality float32
}

// Encode writes the Image m to w in WebP format
func Encode(w io.Writer, m image.Image, o *Options) error {
	if o == nil {
		o = &Options{Quality: 70}
	}
	return webp.Encode(w, m, &webp.Options{Quality: o.Quality})
}
