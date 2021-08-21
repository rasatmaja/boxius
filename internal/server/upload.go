package server

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/boxius/internal/images"
)

// BasePath ...
const BasePath = "public"

// ImagePath ...
const ImagePath = "images"

// ImageUploadReq request
type ImageUploadReq struct {
	TargetPath     string `json:"target_path" form:"target_path"`
	TargetFilename string `json:"target_filename" form:"target_filename"`
}

// ImageUploadRes response
type ImageUploadRes struct {
	URL string `json:"url"`
}

// ImageUpload ...
func (A *App) ImageUpload(c *fiber.Ctx) error {

	req := new(ImageUploadReq)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Error cant parse request, got: %v", err))
	}

	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Error getting image file, got: %v", err))
	}

	img, err := images.Decode(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error decode image, got: %v", err))
	}

	// build image path
	path := filepath.Join(BasePath, ImagePath, req.TargetPath)
	filename := req.TargetFilename

	if len(filename) == 0 {
		filename = strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename))
	}

	if path, err = images.SaveWebP(img, path, filename); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error upload image, got: %v", err))
	}

	path = strings.TrimPrefix(path, BasePath)
	return c.Status(fiber.StatusOK).JSON(&ImageUploadRes{
		URL: filepath.Join(c.Hostname(), path),
	})
}
