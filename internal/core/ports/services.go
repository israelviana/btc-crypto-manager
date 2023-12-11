package ports

import (
	"bitcoin-challenge/internal/core/domain"
)

type BitcoinService interface {
	FindDetailsPerAddress(address string) (*domain.DetailsAddress, error)
	FindBalancePerAddress(address string) (*domain.BalanceDetail, error)
	MountUTXO(address string, amountNeeded string) (*domain.UTXOs, error)
	FindDetailsPerTransactionId(transactionId string) (*domain.Transaction, error)
}

type KleverService interface {
	RequestAddressDetailsPerAddress(address string) (*domain.AddressDetails, error)
	RequestAddressUTXODetails(address string) (*[]domain.UTXODetails, error)
	RequestTransactionDetails(transactionId string) (*domain.TransationDetails, error)
}
