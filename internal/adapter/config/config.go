package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App *App
		HTTP *HTTP
		DB *DB
		Redis *Redis
		JWT *JWT
	}

	App struct {
		Name string
		Env string
	}

	HTTP struct {
		Host string
		Port string
		AllowedOrigins string
	}

	DB struct {
		User string
		Password string
		Host string
		Port string
		Name string
	}

	Redis struct {
		Host string
		Port string
		Password string
	}

	JWT struct {
		Secret string
		Duration string
		Env string
		Host string
	}
)

func InitConfig() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("failed to load .env file: %v", err.Error())
		}
	}

	App := &App{
		Name: os.Getenv("APP_NAME"),
		Env: os.Getenv("APP_ENV"),
	}

	HTTP := &HTTP{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("HTTP_PORT"),
		AllowedOrigins: os.Getenv("ALLOWED_ORIGINS"),
	}

	DB := &DB{
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
	}

	Redis := &Redis{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASS"),
	}

	JWT := &JWT{
		Secret: os.Getenv("JWT_SECRET"),
		Duration: os.Getenv("TOKEN_DURATION"),
		Env: os.Getenv("APP_ENV"),
		Host: os.Getenv("HTTP_HOST"),
	}

	return &Container{
		App: App,
		HTTP: HTTP,
		DB: DB,
		Redis: Redis,
		JWT: JWT,
	}, nil
}
