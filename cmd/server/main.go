package main

import (
	"fmt"
	"log"
	"maneko/media-proxy/internal/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.MustLoad()

	app := fiber.New(fiber.Config{
		AppName:      "media-proxy",
		ServerHeader: "media-proxy",
	})

	addr := fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)

	if err := app.Listen(addr); err != nil {
		log.Fatalf("failed to start media-proxy server: %v", err)
	}
}
