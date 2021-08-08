package server

import (
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/boxius/internal/config"
)

func TestNew(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		New()
	})
}

func TestServer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svr := fiber.New()
		app := &App{
			server: svr,
			env:    config.LoadENV(),
		}
		go OSInterupt(t)
		app.Start()

	})

	t.Run("error", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		svr := fiber.New()
		app := &App{
			server: svr,
			env:    config.LoadENV(),
		}
		app.env.ServerHost = "9009009090"
		app.Start()

	})
}

func OSInterupt(t *testing.T) {
	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal(err)
		t.Fail()
	}
	time.Sleep(2 * time.Second)
	proc.Signal(os.Interrupt)
}
