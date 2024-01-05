package main

import (
	// orm "backend/orm"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	// db := orm.OpenDb()

	// Serve static files from the public directory
	// app.Static("*", "../public/index.html")

	app.Listen(":3000")
}
