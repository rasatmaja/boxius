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
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Error getting image from form file, got: %v", err))
	}

	_, err = images.Decode(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error decode image, got: %v", err))
	}

	// Save file to root directory:
	err = c.SaveFile(file, fmt.Sprintf("./public/images/%s", file.Filename))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error upload image from form file, got: %v", err))
	}

	return c.SendStatus(fiber.StatusOK)
}
