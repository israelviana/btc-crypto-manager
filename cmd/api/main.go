package main

import (
	"bitcoin-challenge/internal/server"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	server.InitServer(app)
}
