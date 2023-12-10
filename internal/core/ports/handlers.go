package ports

import (
	"github.com/gofiber/fiber/v2"
)

type BitcoinHandler interface {
	FindDetailsPerAddress(app *fiber.Ctx) error
	FindBalancePerAddress(app *fiber.Ctx) error
	MountUTXO(app *fiber.Ctx) error
	FindDetailsPerTransactionId(app *fiber.Ctx) error
}
