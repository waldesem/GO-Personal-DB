package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/app/controllers"
	"backend/pkg/middlewares"
)

func LoginRoutes(a *fiber.App) {

	loginGroup := a.Group("/login")
	loginGroup.Get("/", middlewares.AuthRequired([]string{}, []string{}), controllers.GetLogin)
	loginGroup.Post("/", controllers.PostLogin)
	loginGroup.Patch("/", controllers.PatchLogin)
	loginGroup.Delete("/", middlewares.AuthRequired([]string{}, []string{}), controllers.DeleteLogin)

	a.Post("/refresh", controllers.RefreshToken)
}
