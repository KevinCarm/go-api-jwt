package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-api-jwt/database"
	"go-api-jwt/models"
)

func Create(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Bad request",
			"error":   err.Error(),
		})
	}

	err := user.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Internal error",
			"error":   err.Error(),
		})
	}

	record := database.Instance.Create(&user)
	if record.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Internal error",
			"error":   record.Error.Error(),
		})
	}

	return c.JSON(user)
}
