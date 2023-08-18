package handler

import (
	"users/internal/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Post("/register", h.register)
	api.Post("/login", h.login)
}
