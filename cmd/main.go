package main

import (
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
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := config.InitDB(cfg)
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

	if err := app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
