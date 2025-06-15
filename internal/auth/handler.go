package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service *Service
}

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) RegisterUser(c echo.Context) error {
	var req UserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := h.Service.RegisterUser(req.Name, req.Email); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "User created"})
}