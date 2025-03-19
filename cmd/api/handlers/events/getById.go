package events

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	domain "github.com/jeissoni/EventLine/internal/domain/custonErrors"
)

func (h EventHandler) GetByID(c *fiber.Ctx) error {

	id_event := c.Params("id")

	_, err := uuid.Parse(id_event)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID",
		})
	}

	event, err := h.EventService.GetByID(id_event)

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
