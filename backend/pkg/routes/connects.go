package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/app/controllers"
	"backend/pkg/middlewares"
)

func ConnectRoutes(a *fiber.App) {

	a.Get(
		"/connects/:page",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
		controllers.GetConnects,
	)

	connectGroup := a.Group(
		"/connect",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
	)
	connectGroup.Post("/", controllers.PostConnect)
	connectGroup.Patch("/:action/:id", controllers.PatchConnect)
	connectGroup.Delete("/:action/:id", controllers.DeleteConnect)
}
