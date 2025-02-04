package events

import (
	"github.com/gofiber/fiber/v2"
)

func (h EventHandler) GetAll(c *fiber.Ctx) error {

	events, err := h.EventService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting all events",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"events": events,
	})

}
