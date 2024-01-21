package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

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
	app.Use(cache.New())
	app.Use(cors.New())
	app.Use(csrf.New())
	app.Use(favicon.New())
	app.Use(recover.New())
	app.Use(logger.New())

	// Static files.
	app.Static("/", "./static/",
		fiber.Static{
			Compress:  true,
			ByteRange: true,
			Browse:    true,
			Index:     "index.html",
		},
	)

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
