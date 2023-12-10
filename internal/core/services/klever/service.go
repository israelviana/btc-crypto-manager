package klever

import (
	"bitcoin-challenge/internal/core/domain"
	"bitcoin-challenge/pkg/logger"
	"bitcoin-challenge/pkg/utils"
	"bitcoin-challenge/pkg/utils/constants"
	"context"
	"strings"
	"time"
)

type Service struct{}

func NewKleverService() *Service {
	return &Service{}
}

func (srv *Service) RequestAddressDetailsPerAddress(address string) (*domain.AddressDetails, error) {
	var addressDetails domain.AddressDetails
	var urlAddressDetailsPerAddress strings.Builder

	urlAddressDetailsPerAddress.WriteString(constants.URLKlever + constants.SufixUrlAddressDetails + address)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := utils.Get(ctx, urlAddressDetailsPerAddress.String(), addressDetails)

	if err != nil {
		logger.Error(constants.ErrorToMakeRequest, err)
		return nil, err
	}

	return &addressDetails, nil
}

func (srv *Service) RequestAddressUTXODetails(address string) (*domain.AddressDetails, error) {
	var addressUtxoDetails domain.AddressDetails
	var urlUtxoDetailsPerAddress strings.Builder

	urlUtxoDetailsPerAddress.WriteString(constants.URLKlever + constants.SufixUrlUTXODetails + address)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := utils.Get(ctx, urlUtxoDetailsPerAddress.String(), addressUtxoDetails)

	if err != nil {
		logger.Error(constants.ErrorToMakeRequest, err)
		return nil, err
	}

	return &addressUtxoDetails, nil
}

func (srv *Service) RequestTransactionDetails(transactionId string) (*domain.TransationDetails, error) {
	var transactionDetails domain.TransationDetails
	var urlTransactionDetails strings.Builder

	urlTransactionDetails.WriteString(constants.URLKlever + constants.SufixUrlTransactionDetails + transactionId)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := utils.Get(ctx, urlTransactionDetails.String(), transactionDetails)

	if err != nil {
		logger.Error(constants.ErrorToMakeRequest, err)
		return nil, err
	}

	return &transactionDetails, nil
}
