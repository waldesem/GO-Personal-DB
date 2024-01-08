package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/app/controllers"
	"backend/pkg/middlewares"
)

func AdminRoutes(a *fiber.App) {

	a.Get("/users", middlewares.JWTProtected(), controllers.GetUsers)

	userGroup := a.Group("/user", middlewares.JWTProtected())
	userGroup.Patch("/", controllers.PatchUser)
	userGroup.Post("/", controllers.PostUser)
	userGroup.Delete("/:id", controllers.DeleteUser)
	userGroup.Get("/:action/:id", controllers.GetUser)

	roleGroup := a.Group("/role", middlewares.JWTProtected())
	roleGroup.Get("/:value/:user_id", controllers.GetRoles)
	roleGroup.Delete("/:value/:user_id", controllers.DelRoles)

	groupGroup := a.Group("/group", middlewares.JWTProtected())
	groupGroup.Get("/:value/:user_id", controllers.GetGroups)
	groupGroup.Delete("/:value/:user_id", controllers.DelGroups)

	tableGroup := a.Group("/table/:item", middlewares.JWTProtected())
	tableGroup.Post("/:page", controllers.PostTablesRows)
	tableGroup.Delete("/:item_id", controllers.DelTableRows)
}
