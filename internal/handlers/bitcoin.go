package handlers

import (
	"bitcoin-challenge/internal/core/ports"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Handler struct {
	bitcoinService ports.BitcoinService
}

func NewBitcoinHandler(bitcoinService ports.BitcoinService) *Handler {
	return &Handler{bitcoinService: bitcoinService}
}

func (srv *Handler) FindDetailsPerAddress(ctx *fiber.Ctx) error {
	address := ctx.Params("address")

	if address == "" {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": "address bitcoin dont be could empty"})
	}

	perAddress, err := srv.bitcoinService.FindDetailsPerAddress(address)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": "error to find bitcoin details per address"})
	}

	return ctx.Status(http.StatusOK).JSON(&perAddress)
}

func (srv *Handler) FindBalancePerAddress(ctx *fiber.Ctx) error {
	address := ctx.Params("address")

	if address == "" {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": "address bitcoin dont be could empty"})
	}

	perAddress, err := srv.bitcoinService.FindBalancePerAddress(address)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": "error to find balance details per address"})
	}

	return ctx.Status(http.StatusOK).JSON(&perAddress)
}

func (srv *Handler) MountUTXO(ctx *fiber.Ctx) error {

	return nil
}

func (srv *Handler) FindDetailsPerTransactionId(ctx *fiber.Ctx) error {
	transactionID := ctx.Params("tx")

	if transactionID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": "transactionID dont be could empty"})
	}

	detailsPerTransaction, err := srv.bitcoinService.FindDetailsPerTransactionId(transactionID)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": "error to find details per transactionID"})
	}

	return ctx.Status(http.StatusOK).JSON(&detailsPerTransaction)
}
