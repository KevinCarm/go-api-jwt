package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-api-jwt/database"
	"go-api-jwt/models"
)

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Bad request",
			"error":   err.Error(),
		})
	}

	record := database.Instance.Create(&product)
	if record.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Internal error",
			"error":   record.Error.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(product)
}

func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product

	record := database.Instance.Find(&products)
	if record.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Internal error",
			"error":   record.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(products)
}

func GetProductById(c *fiber.Ctx) error {
	var id string = c.Params("id")
	var product models.Product

	record := database.Instance.Find(&product, id)
	if record.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Internal error",
			"error":   record.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	var id string = c.Params("id")
	var product models.Product

	record := database.Instance.Delete(&product, id)
	if record.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Internal error",
			"error":   record.Error.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON("User deleted successfully")
}

func UpdateProduct(c *fiber.Ctx) error {
	var id string = c.Params("id")
	var newProduct models.Product
	var findProduct models.Product

	if err := c.BodyParser(&newProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Bad request",
			"error":   err.Error(),
		})
	}

	record := database.Instance.Find(&findProduct, id)
	if record.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Internal error",
			"error":   record.Error.Error(),
		})
	}

	findProduct.Name = newProduct.Name
	findProduct.Price = newProduct.Price

	database.Instance.Save(&findProduct)
	return c.Status(fiber.StatusOK).JSON(findProduct)
}
