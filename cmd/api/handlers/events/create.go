package events

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeissoni/EventLine/internal/domain/entities"
)

func (h EventHandler) CreateEvent(c *fiber.Ctx) error {

	var event entities.Event

	err := c.BodyParser(&event)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = h.EventService.Create(event)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(event)
}
