package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/app/controllers"
	"backend/pkg/middlewares"
)

func AdminRoutes(a *fiber.App) {

	a.Get(
		"/users",
		middlewares.AuthRequired([]string{"admin"}, []string{"admins"}),
		controllers.GetUsers,
	)

	userGroup := a.Group(
		"/user",
		middlewares.AuthRequired([]string{"admin"}, []string{"admins"}),
	)
	userGroup.Patch("/", controllers.PatchUser)
	userGroup.Post("/", controllers.PostUser)
	userGroup.Delete("/:id", controllers.DeleteUser)
	userGroup.Get("/:action/:id", controllers.GetUser)

	roleGroup := a.Group(
		"/role/:value/:user_id",
		middlewares.AuthRequired([]string{"admin"}, []string{"admins"}),
	)
	roleGroup.Get("/", controllers.GetRoles)
	roleGroup.Delete("/", controllers.DelRoles)

	groupGroup := a.Group(
		"/group:value/:user_id",
		middlewares.AuthRequired([]string{"admin"}, []string{"admins"}),
	)
	groupGroup.Get("/", controllers.GetGroups)
	groupGroup.Delete("/", controllers.DelGroups)

	tableGroup := a.Group(
		"/table/:item",
		middlewares.AuthRequired([]string{"admin"}, []string{"admins"}),
	)
	tableGroup.Post("/:page", controllers.PostTablesRows)
	tableGroup.Delete("/:item_id", controllers.DelTableRows)
}
