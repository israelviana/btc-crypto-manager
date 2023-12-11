package bitcoin

import (
	"bitcoin-challenge/internal/core/domain"
	"bitcoin-challenge/internal/core/ports"
	"errors"
	"strconv"
)

type Service struct {
	kleverService ports.KleverService
}

func NewBitcoinService(kleverService ports.KleverService) *Service {
	return &Service{kleverService: kleverService}
}

func (srv *Service) FindDetailsPerAddress(address string) (*domain.DetailsAddress, error) {
	addressDetails, err := srv.kleverService.RequestAddressDetailsPerAddress(address)
	if err != nil {
		return nil, errors.New("error to get address details in klever api")
	}

	var t domain.DetailsAddress

	t.Address = addressDetails.Address
	t.Balance = addressDetails.Balance
	t.TotalTx = addressDetails.Txs
	t.Total.Received = addressDetails.TotalReceived
	t.Total.Sent = addressDetails.TotalSent
	t.BalanceDetail.Unconfirmed = addressDetails.UnconfirmedBalance

	return &t, nil
}

func (srv *Service) FindBalancePerAddress(address string) (*domain.BalanceDetail, error) {
	var balanceDetails domain.BalanceDetail

	details, err := srv.kleverService.RequestAddressUTXODetails(address)
	if err != nil {
		return nil, errors.New("error to get UTXO details in klever api")
	}

	var valueConfirmed int
	var valueUnconfirmed int
	for _, detail := range *details {
		valueInteger, err := strconv.Atoi(detail.Value)
		if err != nil {
			return nil, errors.New("error to convert value of UTXO details of klever api")
		}

		if detail.Confirmations < 2 {
			valueUnconfirmed += valueInteger
		} else {
			valueConfirmed += valueInteger
		}
	}

	balanceDetails.Unconfirmed = strconv.Itoa(valueUnconfirmed)
	balanceDetails.Confirmed = strconv.Itoa(valueConfirmed)

	return &balanceDetails, nil
}

func (srv *Service) MountUTXO(address string, amountNeeded string) (*domain.UTXOs, error) {
	var selectedUTXOs []domain.BitcoinReturnSend
	var utxosItem domain.UTXOs

	var totalAmount int

	utxoDetails, err := srv.kleverService.RequestAddressUTXODetails(address)
	if err != nil {
		return nil, errors.New("error to get UTXO details in klever api")
	}

	amountValueNeeded, _ := strconv.Atoi(amountNeeded)
	for _, utxo := range *utxoDetails {
		utxoValue, _ := strconv.Atoi(utxo.Value)

		totalAmount += utxoValue
		selectedUTXOs = append(selectedUTXOs, domain.BitcoinReturnSend{
			TxId:   utxo.TxID,
			Amount: utxo.Value,
		})

		if totalAmount >= amountValueNeeded {
			break
		}
	}

	utxosItem.UTXOs = selectedUTXOs

	if totalAmount < amountValueNeeded {
		return nil, errors.New("selected UTXOs do not cover the required amount")
	}

	return &utxosItem, nil
}

func (srv *Service) FindDetailsPerTransactionId(transactionId string) (*domain.Transaction, error) {
	var transaction domain.Transaction

	transactionDetails, id := srv.kleverService.RequestTransactionDetails(transactionId)
	if id != nil {
		return nil, id
	}

	transaction.TxID = transactionDetails.TxId
	transaction.Block = transactionDetails.BlockHeight

	for _, out := range transactionDetails.VOut {
		var address domain.Address
		address.Value = out.Value
		for _, addressItem := range out.Addresses {
			address.Address = addressItem
			transaction.Addresses = append(transaction.Addresses, address)
		}
	}

	return &transaction, nil
}
