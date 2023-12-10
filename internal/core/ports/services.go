package ports

import (
	"bitcoin-challenge/internal/core/domain"
)

type BitcoinService interface {
	FindDetailsPerAddress(address string) (*domain.AddressDetails, error)
	FindBalancePerAddress(address string) (*domain.AddressDetails, error)
	MountUTXO() (*domain.Send, error)
	FindDetailsPerTransactionId(transactionId string) (*domain.Transaction, error)
}

type KleverService interface {
	RequestAddressDetailsPerAddress(address string) (*domain.AddressDetails, error)
	RequestAddressUTXODetails(address string) (*domain.AddressDetails, error)
	RequestTransactionDetails(transactionId string) (*domain.TransationDetails, error)
}
