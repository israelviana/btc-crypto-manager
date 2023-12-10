package handlers

import (
	"bitcoin-challenge/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	bitcoinService ports.BitcoinService
}

func NewBitcoinHandler(bitcoinService ports.BitcoinService) *Handler {
	return &Handler{bitcoinService: bitcoinService}
}

func (srv *Handler) FindDetailsPerAddress(ctx *fiber.Ctx) error {

	return nil
}

func (srv *Handler) FindBalancePerAddress(ctx *fiber.Ctx) error {

	return nil
}

func (srv *Handler) MountUTXO(ctx *fiber.Ctx) error {

	return nil
}

func (srv *Handler) FindDetailsPerTransactionId(ctx *fiber.Ctx) error {

	return nil
}
