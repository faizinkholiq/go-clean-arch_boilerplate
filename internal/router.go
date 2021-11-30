package api

import (
	home "github.com/faizinkholiq/gofiber_boilerplate/internal/app/home"
	fiber "github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()

	app.Get("/", home.Index)
	app.Listen(":3000")
}