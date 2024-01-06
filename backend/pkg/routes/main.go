package routes

import (
	"log"

	routes "backend/internal"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Static("/", "../public/index.html")

	app.Use("/login", loginHandler)

	app.Post("/index/:item/:page", indexHandler)

	log.Fatal(app.Listen(":3000"))
}

func loginHandler(c *fiber.Ctx) error {
	switch c.Method() {
	case "GET":
		return c.SendString("login")
	case "POST":
		return c.SendString("login")
	case "PUT":
		return c.SendString("login")
	case "DELETE":
		return c.SendString("login")
	}
	return nil
}

// indexHandler handles the request to the index route.
func indexHandler(c *fiber.Ctx) error {
	payload := struct {
		Search string `json:"search"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		log.Println(err)
	}
	result, hasNext, hasPrev := routes.HandleIndex(c.Params("item"), c.Params("page"), payload.Search)
	return c.JSON(fiber.Map{"result": result, "hasNext": hasNext, "hasPrev": hasPrev})
}
