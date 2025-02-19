package events

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	domain "github.com/jeissoni/EventLine/internal/domain/custonErrors"
)

func (h EventHandler) GetByID(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	event, err := h.EventService.GetByID(id)

	if err != nil {

		var appErr domain.DomainError

		if errors.As(err, &appErr) {

			if appErr.Code == domain.ErrCodeNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": appErr.Message,
				})
			}

		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(event)
}
