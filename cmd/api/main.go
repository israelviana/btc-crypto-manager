package main

import (
	"bitcoin-challenge/internal/server"
	"github.com/gofiber/fiber/v2"
)

// @title BTC-CRYPTO-MANAGER
// @description API for bitcoin operations.
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()
	server.InitServer(app)
}
