package server

import (
	"bitcoin-challenge/internal/core/services/bitcoin"
	"bitcoin-challenge/internal/core/services/klever"
	"bitcoin-challenge/internal/handlers"
	"bitcoin-challenge/internal/routers"
	"bitcoin-challenge/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func InitServer(app *fiber.App) {
	logger.InitZapLogger()

	//Init Services
	kleverService := klever.NewKleverService()
	bitcoinService := bitcoin.NewBitcoinService(kleverService)

	//InitHandlers
	bitcoinHandler := handlers.NewBitcoinHandler(bitcoinService)

	//cors
	crs := cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, DELETE, PUT, OPTIONS",
		AllowHeaders:     "*",
		AllowCredentials: true,
	})
	app.Use(crs)

	health := app.Group("/health")
	bitcoinGroup := app.Group("/")

	routers.MapBitcoinRoutes(bitcoinGroup, bitcoinHandler)

	health.Get("", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(map[string]string{"status": "OK"})
	})

	if err := godotenv.Load(".env"); err != nil {
		logger.Fatal("error to load .env", err)
	}

	if err := app.Listen(os.Getenv("port_application")); err != nil {
		logger.Fatal("error to up api", err)
	}
}
