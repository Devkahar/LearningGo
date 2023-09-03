package route

import (
	"fiber-api/pkg/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(api fiber.Router) {
	api.Get("/authPage", controller.AuthPage)
}
