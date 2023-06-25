package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shinohara320/travel-agent/controller"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
}
