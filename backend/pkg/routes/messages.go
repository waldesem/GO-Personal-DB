package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/app/controllers"
	"backend/pkg/middlewares"
)

func MessageRoutes(a *fiber.App) {

	messageGroup := a.Group(
		"/messages",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
	)
	messageGroup.Delete("/:action/:id", controllers.DeleteMessage)
	messageGroup.Get("/:action", controllers.GetMessages)
}
