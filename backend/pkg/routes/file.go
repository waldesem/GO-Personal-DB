package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/app/controllers"
	"backend/pkg/middlewares"
)

func FileRoutes(a *fiber.App) {

	managerGroup := a.Group("/manager", middlewares.JWTProtected())
	managerGroup.Get("/", controllers.GetFiles)
	managerGroup.Post("/:action", controllers.PostFiles)
}
