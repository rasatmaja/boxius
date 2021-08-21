package images

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/rasatmaja/boxius/internal/images/webp"
)

// SaveWebP a function to save image into WebP format
func SaveWebP(img image.Image, path, filename string) (string, error) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}

	filename = fmt.Sprintf("%s.webp", filename)
	path = filepath.Join(path, filename)
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return "", err
	}

	return path, webp.Encode(file, img, nil)
}
