package auth

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group, h *Handler) {
	g.POST("/register", h.RegisterUser)
}