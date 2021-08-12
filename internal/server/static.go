package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// FileSystem is a function to start file syetem to serve static content
func (a *App) FileSystem() {

	fmt.Println("[ SRVR ] Preparing filesystem...")

	a.server.Static("/", "./public", fiber.Static{ByteRange: true})
}
