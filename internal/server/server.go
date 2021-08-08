package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// App is struct that define server, repo, and all component app needs
type App struct {
	server *fiber.App
}

// New is a function to initialize sever and its component
func New() *App {

	// setup server
	svr := fiber.New(
		fiber.Config{
			UnescapePath: true,
		},
	)

	return &App{
		server: svr,
	}
}

// Start is a function to start server
func (a *App) Start() {
	fmt.Println("[ SRVR ] Starting server...")
	a.FileSystem()
	a.Route()
	a.InitializeShutdownSequence()
	defer fmt.Println("[ SRVR ] Server Shutdown...")

	a.ServerListen()
}

// ServerListen is a function to initialize server listen
func (a *App) ServerListen() {

	var err error
	host := fmt.Sprintf("%s:%d", "localhost", 9000)

	err = a.server.Listen(host)

	if err != nil {
		panic(err)
	}
}
