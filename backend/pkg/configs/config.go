package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout:  time.Second * time.Duration(readTimeoutSecondsCount),
		ServerHeader: "Fiber",
		AppName:      "GO-Personal-DB v.0.1.0",
	}
}
