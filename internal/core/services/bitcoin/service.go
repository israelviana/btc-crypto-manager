package bitcoin

import (
	"bitcoin-challenge/internal/core/domain"
	"bitcoin-challenge/internal/core/ports"
	"errors"
)

type Service struct {
	kleverService ports.KleverService
}

func NewBitcoinService(kleverService ports.KleverService) *Service {
	return &Service{kleverService: kleverService}
}

func (srv *Service) FindDetailsPerAddress(address string) (*domain.AddressDetails, error) {
	addressDetails, err := srv.kleverService.RequestAddressDetailsPerAddress(address)
	if err != nil {
		return nil, errors.New("error to get address details in klever api")
	}

	return addressDetails, nil
}

func (srv *Service) FindBalancePerAddress(address string) (*domain.AddressDetails, error) {
	var addressDetails domain.AddressDetails

	return &addressDetails, nil
}

func (srv *Service) MountUTXO() (*domain.Send, error) {
	var send domain.Send

	return &send, nil
}

func (srv *Service) FindDetailsPerTransactionId(transactionId string) (*domain.Transaction, error) {
	var transaction domain.Transaction

	return &transaction, nil
}
