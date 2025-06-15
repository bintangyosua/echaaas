package auth

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type Module struct{}

func (m Module) RegisterRoutes(e *echo.Echo, c *dig.Container) {
	err := c.Invoke(func(h *Handler) {
		authGroup := e.Group("/api/auth")
		RegisterRoutes(authGroup, h)
	})

	if err != nil {
		panic(err)
	}
}
