package events

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h EventHandler) Delete(c *fiber.Ctx) error {

	id_event := c.Params("id")

	_, err := uuid.Parse(id_event)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID",
		})
	}

	err = h.EventService.Delete(id_event)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(nil)

}
