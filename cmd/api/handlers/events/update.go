package events

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	domain "github.com/jeissoni/EventLine/internal/domain/entities"
)

func (h EventHandler) Update(c *fiber.Ctx) error {

	id_event := c.Params("id")

	_, err := uuid.Parse(id_event)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID",
		})
	}

	// If the UUID exists, we can proceed to parse the body
	// and update the event
	_, err = h.EventService.GetByID(id_event)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Event not found",
		})
	}

	var event domain.Event

	err = c.BodyParser(&event)
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
