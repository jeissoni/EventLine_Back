package user

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/jeissoni/EventLine/internal/domain/entities"
)

// CreateUser handles the creation of a new user
func (h UserHandler) CreateUser(c *fiber.Ctx) error {
	var user domain.User

	// Parse the request body into the user struct
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Call the service to create the user
	err = h.UserService.Create(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
