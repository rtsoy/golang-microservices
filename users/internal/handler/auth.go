package handler

import (
	"users/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) register(c *fiber.Ctx) error {
	var input domain.RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	user, err := h.svc.CreateUser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
