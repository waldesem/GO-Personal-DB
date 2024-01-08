package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/app/controllers"
	"backend/pkg/middlewares"
)

func PublicRoutes(a *fiber.App) {

	a.Post("/index/:item/:page", middlewares.JWTProtected(), controllers.PostIndex)
	a.Get("/classes", controllers.GetClasses)

	// Create routes group.
	// route := a.Group("/api/v1")

	// Routes for POST method:
	// route.Post("/book", middlewares.JWTProtected(), controllers.CreateBook)           // create a new book
	// route.Post("/user/sign/out", middlewares.JWTProtected(), controllers.UserSignOut) // de-authorization user
	// route.Post("/token/renew", middlewares.JWTProtected(), controllers.RenewTokens)   // renew Access & Refresh tokens
	// route.Delete("/book", middlewares.JWTProtected(), controllers.DeleteBook) // delete one book by ID
}
