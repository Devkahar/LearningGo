package controller

import "github.com/gofiber/fiber/v2"

func AuthPage(ctx *fiber.Ctx) error {
	return ctx.SendString("Hey Welcome to auth page")
}
