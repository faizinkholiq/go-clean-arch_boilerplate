package controllers

import (
	"database/sql"

	"github.com/faizinkholiq/gofiber_boilerplate/internal/entities"
	"github.com/faizinkholiq/gofiber_boilerplate/internal/repositories"
	"github.com/faizinkholiq/gofiber_boilerplate/internal/usecases"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App, db *sql.DB, redisClient *redis.Client) {
	repo := &repositories.UserRepo{DB: db}
	usecase := &usecases.UserUseCase{Repo: repo}

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.Status(201).SendString("Hei yo!!!")
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(entities.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).SendString("Bad Request")
		}

		err := usecase.RegisterUser(user)
		if err != nil {
			return c.Status(500).SendString("Could not create user")
		}

		return c.Status(201).JSON(user)
	})
}
