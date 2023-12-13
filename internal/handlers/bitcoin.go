package handlers

import (
	"bitcoin-challenge/internal/core/domain"
	"bitcoin-challenge/internal/core/ports"
	"bitcoin-challenge/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Handler struct {
	bitcoinService ports.BitcoinService
}

func NewBitcoinHandler(bitcoinService ports.BitcoinService) *Handler {
	return &Handler{bitcoinService: bitcoinService}
}

// FindDetailsPerAddress gets the details for a specific Bitcoin address
// @Summary Find details per Bitcoin address
// @Description Retrieves detailed information for a given Bitcoin address
// @Tags Bitcoin
// @Accept  json
// @Produce  json
// @Param   address path string true "Bitcoin Address"
// @Success 200 {object} domain.DetailsAddress "Bitcoin Details"
// @Failure 400 {object} domain.HTTPErrorResponse "Bad Request"
// @Router /address/{address} [get]
func (srv *Handler) FindDetailsPerAddress(ctx *fiber.Ctx) error {
	address := ctx.Params("address")

	if !utils.ValidateBitcoinAddress(address) {
		return utils.HTTPFail(ctx, http.StatusBadRequest, nil, "invalid address bitcoin")
	}

	perAddress, err := srv.bitcoinService.FindDetailsPerAddress(address)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "error to find bitcoin details per address")
	}

	return ctx.Status(http.StatusOK).JSON(&perAddress)
}

// FindBalancePerAddress gets the balance for a specific Bitcoin address
// @Summary Find balance per Bitcoin address
// @Description Retrieves balance information for a given Bitcoin address
// @Tags Bitcoin
// @Accept  json
// @Produce  json
// @Param   address path string true "Bitcoin Address"
// @Success 200 {object} domain.BalanceDetail "Bitcoin Balance"
// @Failure 400 {object} domain.HTTPErrorResponse "Bad Request"
// @Router /balance/{address} [get]
func (srv *Handler) FindBalancePerAddress(ctx *fiber.Ctx) error {
	address := ctx.Params("address")

	if !utils.ValidateBitcoinAddress(address) {
		return utils.HTTPFail(ctx, http.StatusBadRequest, nil, "invalid address bitcoin")
	}

	perAddress, err := srv.bitcoinService.FindBalancePerAddress(address)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "error to find balance details per address")
	}

	return ctx.Status(http.StatusOK).JSON(&perAddress)
}

// MountUTXO mounts a UTXO set for a Bitcoin transaction
// @Summary Mount UTXO
// @Description Creates a UTXO set for a given Bitcoin address and amount
// @Tags Bitcoin
// @Accept  json
// @Produce  json
// @Param   request body domain.BitcoinRequest true "Bitcoin Request"
// @Success 200 {object} []domain.UTXOs "UTXO"
// @Failure 400 {object} domain.HTTPErrorResponse "Bad Request"
// @Router /send [post]
func (srv *Handler) MountUTXO(ctx *fiber.Ctx) error {
	var request domain.BitcoinRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error to read received body"})
	}

	utxo, err := srv.bitcoinService.MountUTXO(request.Address, request.Amount)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error to read received body"})
	}

	return ctx.Status(fiber.StatusOK).JSON(&utxo)
}

// FindDetailsPerTransactionId gets details for a specific Bitcoin transaction ID
// @Summary Find details per transaction ID
// @Description Retrieves detailed information for a given Bitcoin transaction ID
// @Tags Bitcoin
// @Accept  json
// @Produce  json
// @Param   tx path string true "Transaction ID"
// @Success 200 {object} domain.Transaction "Transaction Details"
// @Failure 400 {object} domain.HTTPErrorResponse "Bad Request"
// @Router /tx/{tx} [get]
func (srv *Handler) FindDetailsPerTransactionId(ctx *fiber.Ctx) error {
	transactionID := ctx.Params("tx")

	if !utils.ValidateTransactionID(transactionID) {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": "invalid transactionID"})
	}

	detailsPerTransaction, err := srv.bitcoinService.FindDetailsPerTransactionId(transactionID)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": "error to find details per transactionID"})
	}

	return ctx.Status(http.StatusOK).JSON(&detailsPerTransaction)
}
