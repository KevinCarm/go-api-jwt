package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-api-jwt/database"
	"go-api-jwt/jwtFiles"
	"go-api-jwt/models"
)

func login(c *fiber.Ctx) error {
	var loginRequest models.LoginRequest
	var user models.User

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	//Check if email and password is correct
	record := database.Instance.Where("email", loginRequest.Email).First(&user)
	if record.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "Email or password incorrect",
		})
	}

	credentialError := user.CheckPassword(loginRequest.Password)
	if credentialError != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "Invalid credentials",
		})
	}
	// Generate TOKEN
	tokenString, err := jwtFiles.GenerateJWT(user.Name, user.Email, user.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": tokenString,
	})
}

func CreateUser(c *fiber.Ctx) error {
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

	return c.Status(fiber.StatusCreated).JSON(user)
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User

	record := database.Instance.Find(&users)
	if record.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Internal error",
			"error":   record.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func DeleteUserById(c *fiber.Ctx) error {
	var id string = c.Params("id")
	var user models.User

	record := database.Instance.Delete(&user, id)
	if record.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Internal error",
			"error":   record.Error.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON("User deleted successfully")
}

func GetUserById(c *fiber.Ctx) error {
	var id string = c.Params("id")
	var user models.User

	record := database.Instance.Find(&user, id)
	if record.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Internal error",
			"error":   record.Error.Error(),
		})
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	var id string = c.Params("id")
	var newUser models.User
	var findUser models.User

	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Internal error",
			"error":   err.Error(),
		})
	}

	record := database.Instance.Find(&findUser, id)
	if record.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "User not found",
			"error":   record.Error.Error(),
		})
	}

	findUser.Name = newUser.Name
	findUser.Email = newUser.Email
	findUser.Password = newUser.Password
	findUser.Role = newUser.Role

	err := findUser.HashPassword(findUser.Password)
	if err != nil {
		return err
	}

	database.Instance.Save(&findUser)

	return c.Status(fiber.StatusOK).JSON(findUser)
}
