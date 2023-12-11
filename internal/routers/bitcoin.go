package routers

import (
	"bitcoin-challenge/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

func MapBitcoinRoutes(bitcoinGroup fiber.Router, bitcoinHandlers ports.BitcoinHandler) {
	bitcoinGroup.Get("details/:address", bitcoinHandlers.FindDetailsPerAddress)
	bitcoinGroup.Get("balance/:address", bitcoinHandlers.FindBalancePerAddress)
	bitcoinGroup.Post("send", bitcoinHandlers.MountUTXO)
	bitcoinGroup.Get("tx/:tx", bitcoinHandlers.FindDetailsPerTransactionId)
}
