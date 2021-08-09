package server

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestFileSysytem(t *testing.T) {
	app := App{
		server: fiber.New(),
	}
	defer app.server.Shutdown()
	app.FileSystem()
}
