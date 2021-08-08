package server

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestInitializeRoute(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := &App{
			server: fiber.New(),
		}

		app.Route()
	})
}

func TestResourceNotFound(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := &App{
			server: fiber.New(),
		}
		app.server.Use(app.ResourceNotfound)
		resp, err := app.server.Test(httptest.NewRequest("GET", "/unknown", nil))

		// begin assert response from http test
		if err != nil {
			t.Error("error should be nil")
			t.Fail()
		}

		if resp.StatusCode != fiber.StatusNotFound {
			t.Errorf("status code should be %d, but got %d", fiber.StatusNotFound, resp.StatusCode)
			t.Fail()
		}
	})
}
