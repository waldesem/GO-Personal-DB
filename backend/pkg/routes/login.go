package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/app/controllers"
	"backend/pkg/middlewares"
)

func LoginRoutes(a *fiber.App) {

	a.Get("/login", middlewares.JWTProtected(), controllers.GetLogin)
	a.Post("/login", controllers.PostLogin)
	a.Patch("/login", controllers.PatchLogin)
	a.Delete("/login", middlewares.JWTProtected(), controllers.DeleteLogin)
	a.Post("/refresh", controllers.RefreshToken)
}
