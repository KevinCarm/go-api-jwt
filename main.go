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
	//User routes
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users", controllers.GetAllUsers)
	app.Get("/api/users/:id", controllers.GetUserById)
	app.Delete("/api/users/:id", controllers.DeleteUserById)
	app.Put("/api/users/:id", controllers.UpdateUser)
	//Product routes
	app.Post("/api/products", controllers.CreateProduct)
	app.Get("/api/products", controllers.GetAllProducts)
	app.Get("/api/products/:id", controllers.GetProductById)
	app.Delete("/api/products/:id", controllers.DeleteProduct)
	app.Put("/api/products/:id", controllers.UpdateProduct)
	log.Println("Running server on port 8080")
	log.Fatal(app.Listen(":8080"))
}
