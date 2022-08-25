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
	go app.Post("/api/login", controllers.Login)
	go app.Post("/api/users", controllers.CreateUser)
	go app.Get("/api/users", controllers.GetAllUsers)
	go app.Get("/api/users/:id", controllers.GetUserById)
	go app.Delete("/api/users/:id", controllers.DeleteUserById)
	go app.Put("/api/users/:id", controllers.UpdateUser)
	//Product routes
	go app.Post("/api/products", controllers.CreateProduct)
	go app.Get("/api/products", controllers.GetAllProducts)
	go app.Get("/api/products/:id", controllers.GetProductById)
	go app.Delete("/api/products/:id", controllers.DeleteProduct)
	go app.Put("/api/products/:id", controllers.UpdateProduct)
	log.Println("Running server on port 8080")
	log.Fatal(app.Listen(":8080"))
}
