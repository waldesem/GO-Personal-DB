package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"backend/pkg/configs"
	"backend/pkg/middlewares"
	"backend/pkg/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Static files.
	app.Static("/", "./static/index.html")

	// Middlewares.
	middlewares.FiberMiddleware(app)

	// Routes.
	routes.LoginRoutes(app)
	routes.MessageRoutes(app)
	routes.AdminRoutes(app)
	routes.PublicRoutes(app)
	routes.FileRoutes(app)
	routes.ConnectRoutes(app)
	routes.NotFoundRoute(app)

	log.Fatal(app.Listen(":3000"))
}
