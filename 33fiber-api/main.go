package main

import (
	"encoding/json"
	"fiber-api/pkg/route"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "AddressbookWithFiber",
	})
	app.Get("/", homeServe)
	app.Route("/api", route.AuthRoute).Name("Auth")
	data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	fmt.Println(string(data))
	log.Fatal(app.Listen(":4000"))
}

func homeServe(ctx *fiber.Ctx) error {
	return ctx.SendString("Back Again with addressbook\nThis time using fiber")
}
