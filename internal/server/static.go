package server

import (
	"fmt"
)

// FileSystem is a function to start file syetem to serve static content
func (a *App) FileSystem() {

	fmt.Println("[ SRVR ] Preparing filesystem...")

	a.server.Static("/", "./public")
}
