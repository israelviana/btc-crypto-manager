package routers

import (
	"bitcoin-challenge/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

func MapBitcoinRoutes(bitcoinGroup fiber.Router, bitcoinHandlers ports.BitcoinHandler) {
	bitcoinGroup.Get("details/:address", bitcoinHandlers.FindDetailsPerTransactionId)
	bitcoinGroup.Get("balance/:address", bitcoinHandlers.FindBalancePerAddress)
	bitcoinGroup.Get("send", bitcoinHandlers.MountUTXO)
	bitcoinGroup.Get("tx:tx", bitcoinHandlers.FindDetailsPerTransactionId)
}