package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/faizinkholiq/go-clean-arch_boilerplate/config"
	api "github.com/faizinkholiq/go-clean-arch_boilerplate/internal/interface/http"
	"github.com/faizinkholiq/go-clean-arch_boilerplate/internal/repository"
	"github.com/faizinkholiq/go-clean-arch_boilerplate/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Println(err)
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

	userHandler := api.NewUserHandler(userUseCase)

	api := app.Group("/api")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Yes")
	})

	api.Get("/users/:id", userHandler.GetUserByID)
	api.Get("/users", userHandler.GetUserList)
	api.Post("/users", userHandler.CreateUser)

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
