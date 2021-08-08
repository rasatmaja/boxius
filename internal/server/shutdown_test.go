package server

import (
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
)

func TestInitializeShutdownSequence(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := App{
			server: &fiber.App{},
		}
		app.InitializeShutdownSequence()

		time.Sleep(2 * time.Second)
		proc, err := os.FindProcess(os.Getpid())
		if err != nil {
			t.Fatal(err)
			t.Fail()
		}
		proc.Signal(os.Interrupt)
	})
}
