package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "media-proxy",
		ServerHeader: "media-proxy",
	})

	addr := fmt.Sprintf("%s:%s", "0.0.0.0", "8080")

	if err := app.Listen(addr); err != nil {
		log.Fatalf("failed to start media-proxy server: %v", err)
	}
}
