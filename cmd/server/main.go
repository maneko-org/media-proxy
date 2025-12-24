package main

import (
	"fmt"
	"log/slog"
	"maneko/media-proxy/internal/config"
	"maneko/media-proxy/internal/logger"
	"maneko/media-proxy/pkg/storage"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	_, err := storage.New(&cfg.Storage, log)
	if err != nil {
		panic(err)
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	addr := fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	log.Info("starting server", slog.String("addr", addr), slog.String("env", cfg.Env))

	if err := app.Listen(addr); err != nil {
		log.Error("failed to start server", slog.String("err", err.Error()))
	}
}
