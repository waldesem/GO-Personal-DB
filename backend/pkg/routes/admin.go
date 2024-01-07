package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/app/controllers"
	"backend/pkg/middlewares"
)

func AdminRoutes(a *fiber.App) {

	a.Get("/users", middlewares.JWTProtected(), controllers.GetLogin)

	userGroup := a.Group("/user", middlewares.JWTProtected())
	userGroup.Patch("/")
	userGroup.Post("/")
	userGroup.Delete("/:id")
	userGroup.Get("/:action/:id")

	roleGroup := a.Group("/role", middlewares.JWTProtected())
	roleGroup.Get("/:value/:user_id")
	roleGroup.Delete("/:value/:user_id")

	groupGroup := a.Group("/group", middlewares.JWTProtected())
	groupGroup.Get("/:value/:user_id")
	groupGroup.Delete("/:value/:user_id")

	tableGroup := a.Group("/table/:item", middlewares.JWTProtected())
	tableGroup.Get("/:page")
	tableGroup.Post("/:page")
	tableGroup.Delete("/:item_id")
	tableGroup.Patch("/:item_id")
}
