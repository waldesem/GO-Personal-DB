package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"backend/pkg/configs"
	"backend/pkg/middlewares"
	"backend/pkg/routes"
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middlewares.FiberMiddleware(app)

	// Routes.
	routes.SwaggerRoute(app)
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	log.Fatal(app.Listen(":3000"))
}
