package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kalougata/gomall/configs"
)

type Config struct {
	DB configs.Database
}

func New() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load env file: %s", err)
	}

	return &Config{
		DB: configs.DatabaseStore(),
	}
}
