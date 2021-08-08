package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Route is a function to server roting
func (a *App) Route() {
	fmt.Println("[ SRVR ] Starting router...")
	defer a.server.Use(a.ResourceNotfound)
}

// ResourceNotfound is a handler to handle undefined route
func (a *App) ResourceNotfound(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotFound)
}
