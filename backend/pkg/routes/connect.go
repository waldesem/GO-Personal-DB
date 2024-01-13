package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/app/controllers"
	"backend/pkg/middlewares"
)

func ConnectRoutes(a *fiber.App) {

	a.Get("/connects/:page", middlewares.JWTProtected(), controllers.GetConnects)
	connectGroup := a.Group("/connect", middlewares.JWTProtected())
	connectGroup.Post("/", controllers.PostConnect)
	connectGroup.Patch("/:action/:id", controllers.PatchConnect)
	connectGroup.Delete("/:action/:id", controllers.DeleteConnect)
}
