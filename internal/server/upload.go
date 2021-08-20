package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/boxius/internal/images"
)

// ImageUpload ...
func (A *App) ImageUpload(c *fiber.Ctx) error {
	// Get first file from form field "image":
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Error getting image file, got: %v", err))
	}

	img, err := images.Decode(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error decode image, got: %v", err))
	}

	if err = images.SaveWebP(img, "./public/images", "test"); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error upload image, got: %v", err))
	}

	return c.SendStatus(fiber.StatusOK)
}
