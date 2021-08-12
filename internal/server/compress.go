package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2/middleware/compress"
)

// Compress is a function to compress response
func (a *App) Compress() {

	fmt.Println("[ SRVR ] Preparing Compression middleware...")

	a.server.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
}
