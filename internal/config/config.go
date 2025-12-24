package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env     string `env:"ENV" env-default:"production"`
	HTTP    HTTPServer
	Storage Storage
}

type HTTPServer struct {
	Host string `env:"HOST" env-default:"0.0.0.0"`
	Port string `env:"PORT" env-default:"8080"`
}

type Storage struct {
	Bucket    string `env:"STORAGE_BUCKET" env-required:"true"`
	Endpoint  string `env:"STORAGE_ENDPOINT" env-required:"true"`
	Region    string `env:"STORAGE_REGION"`
	AccessKey string `env:"STORAGE_ACCESS_KEY"`
	SecretKey string `env:"STORAGE_SECRET_KEY"`
}

func MustLoad() *Config {
	_ = godotenv.Load()

	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(fmt.Errorf("failed to parse environment variables: %w", err))
	}

	return &cfg
}
