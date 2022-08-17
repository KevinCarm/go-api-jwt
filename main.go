package main

import (
	"github.com/gofiber/fiber/v2"
	"go-api-jwt/controllers"
	"go-api-jwt/database"
	"log"
)

func main() {
	database.Connect("root:12345678@tcp(localhost:3306)/gorm_db")
	database.Migrate()

	app := fiber.New()
	app.Post("/api/users", controllers.Create)
	app.Get("/api/users", controllers.GetAll)
	log.Fatal(app.Listen(":8080"))
}
