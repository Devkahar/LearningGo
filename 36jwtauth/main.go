package main

import (
	"jwt/auth/pkg/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connection()
	app := fiber.New(fiber.Config{
		AppName: "JWT-Auth",
	})
	app.Get("/", homeServe)
	app.Listen(":4000")
}

func homeServe(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to jwt auth.")
}
