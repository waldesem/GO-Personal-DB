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
	app.Static("/", "./public/index.html")

	// Middlewares.
	middlewares.FiberMiddleware(app)

	// Routes.
	routes.LoginRoutes(app)
	// routes.AdmiRoutes(app)
	routes.PublicRoutes(app)
	routes.NotFoundRoute(app)

	log.Fatal(app.Listen(":3000"))
}
