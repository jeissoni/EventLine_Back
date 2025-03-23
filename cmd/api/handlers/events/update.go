package events

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/jeissoni/EventLine/internal/domain/entities"
)

func (h EventHandler) Update(c *fiber.Ctx) error {

	var event domain.Event

	err := c.BodyParser(&event)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = h.EventService.Update(event)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(event)

}
