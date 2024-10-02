package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/faizinkholiq/go-clean-arch_boilerplate/config"
	"github.com/faizinkholiq/go-clean-arch_boilerplate/internal/interface/http"
	"github.com/faizinkholiq/go-clean-arch_boilerplate/internal/repository"
	"github.com/faizinkholiq/go-clean-arch_boilerplate/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbUser := config.App.DBUser
	dbPass := config.App.DBPass
	dbName := config.App.DBName
	dbHost := config.App.DBHost

	connection := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", dbUser, dbPass, dbName, dbHost)
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)

	userUseCase := usecase.NewUserUseCase(userRepo)

	app := fiber.New()

	userHandler := http.NewUserHandler(userUseCase)

	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		result := make(map[string]interface{})
		result["user_agent"] = c.Get("User-Agent")
		result["ip_address"] = c.IP()
		result["message"] = "Hello, Home 👋!"

		return c.JSON(result)
	})

	api.Get("/users/:id", userHandler.GetUserByID)
	api.Get("/users", userHandler.GetUserList)
	api.Post("/users", userHandler.CreateUser)

	log.Printf("Starting server on :%s", config.App.AppPort)
	if err := app.Listen(":" + config.App.AppPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
