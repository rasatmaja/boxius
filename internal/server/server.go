package server

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/boxius/internal/config"
)

// App is struct that define server, repo, and all component app needs
type App struct {
	server *fiber.App
	env    *config.ENV
}

// New is a function to initialize sever and its component
func New() *App {

	// load env
	env := config.LoadENV()

	// setup server
	svr := fiber.New(
		fiber.Config{
			ReadTimeout:  time.Duration(env.ServerReadTO) * time.Second,
			WriteTimeout: time.Duration(env.ServerWriteTO) * time.Second,
			IdleTimeout:  time.Duration(env.ServerIdleTO) * time.Second,
			UnescapePath: true,
		},
	)

	return &App{
		server: svr,
		env:    env,
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
	host := fmt.Sprintf("%s:%d", a.env.ServerHost, a.env.ServerPort)

	err = a.server.Listen(host)

	if err != nil {
		panic(err)
	}
}
