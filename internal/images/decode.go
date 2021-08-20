package images

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
)

const (
	// JPEGMime jpeg/jpg content type
	JPEGMime = "image/jpeg"
	// PNGMime png content type
	PNGMime = "image/png"
)

// Decode images file
func Decode(file *multipart.FileHeader) (image.Image, error) {

	rawImg, err := file.Open()
	defer rawImg.Close()
	if err != nil {
		return nil, err
	}

	fileType := file.Header["Content-Type"][0]
	switch fileType {
	case JPEGMime:
		return jpeg.Decode(rawImg)
	case PNGMime:
		return png.Decode(rawImg)
	}

	return nil, fmt.Errorf("image type:%s not supported", fileType)
}
