package main

import (
	"log"
	"myapp/internal/controllers"
	"myapp/internal/infrastructure"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, err := infrastructure.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	redisClient := infrastructure.InitRedis()

	controllers.RegisterUserRoutes(app, db, redisClient)

	log.Fatal(app.Listen(":8080"))
}
