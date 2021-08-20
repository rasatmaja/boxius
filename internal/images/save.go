package images

import (
	"fmt"
	"image"
	"os"

	"github.com/rasatmaja/boxius/internal/images/webp"
)

// SaveWebP a function to save image into WebP format
func SaveWebP(img image.Image, path, filename string) error {

	file, err := os.Create(fmt.Sprintf("%s/%s.webp", path, filename))
	defer file.Close()
	if err != nil {
		return err
	}

	return webp.Encode(file, img, nil)
}
