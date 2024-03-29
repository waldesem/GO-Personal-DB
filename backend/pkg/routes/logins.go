package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/app/controllers"
	"backend/pkg/middlewares"
)

func LoginRoutes(a *fiber.App) {

	a.Post("/login", controllers.PostLogin)
	a.Patch("/login", controllers.PatchLogin)
	a.Get("/login", middlewares.AuthRequired([]string{}, []string{}), controllers.GetLogin)
	a.Delete("/login", middlewares.AuthRequired([]string{}, []string{}), controllers.DeleteLogin)

	a.Post("/refresh", controllers.RefreshToken)
}
