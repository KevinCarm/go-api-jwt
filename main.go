package main

import "go-api-jwt/database"

func main() {
	database.Connect("root:12345678@tcp(localhost:3306)/gorm_db")
	database.Migrate()
}
