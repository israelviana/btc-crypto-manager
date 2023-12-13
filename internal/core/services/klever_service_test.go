package services

import (
	"bitcoin-challenge/internal/core/services/klever"
	"testing"
)

func TestRequestAddressDetailsPerAddress(t *testing.T) {
	kleverService := klever.NewKleverService()

	type TestBitcoinAddress struct {
		BitcoinAddress string
		Expected       bool
	}

	caseTestBitcoinAddress := TestBitcoinAddress{
		BitcoinAddress: "bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r",
		Expected:       true,
	}

	address, err := kleverService.RequestAddressDetailsPerAddress(caseTestBitcoinAddress.BitcoinAddress)
	if err != nil {
		t.Errorf("Error to find details per address test: %s", err.Error())
	}

	result := address.Address != ""

	if result != caseTestBitcoinAddress.Expected {
		t.Errorf("Output %t not equal to expected %t", result, caseTestBitcoinAddress.Expected)
	}
}

func TestRequestAddressUTXODetails(t *testing.T) {
	kleverService := klever.NewKleverService()

	type TestBitcoinAddress struct {
		BitcoinAddress string
		Expected       bool
	}

	caseTestBitcoinAddress := TestBitcoinAddress{
		BitcoinAddress: "bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r",
		Expected:       false,
	}

	utxoDetails, err := kleverService.RequestAddressUTXODetails(caseTestBitcoinAddress.BitcoinAddress)
	if err != nil {
		t.Errorf("Error to find utxo details per address test: %s", err.Error())
	}

	result := len(*utxoDetails) > 0

	if result != caseTestBitcoinAddress.Expected {
		t.Errorf("Output %t not equal to expected %t", result, caseTestBitcoinAddress.Expected)
	}
}

func TestRequestTransactionDetails(t *testing.T) {
	kleverService := klever.NewKleverService()

	type TestTransactionID struct {
		TransactionID string
		Expected      bool
	}

	caseTestTransactionID := TestTransactionID{
		TransactionID: "3654d26660dcc05d4cfb25a1641a1e61f06dfeb38ee2279bdb049d018f1830ab",
		Expected:      true,
	}

	transactionDetails, err := kleverService.RequestTransactionDetails(caseTestTransactionID.TransactionID)
	if err != nil {
		t.Errorf("Error to find details per address test: %s", err.Error())
	}

	result := transactionDetails.TxId != ""

	if result != caseTestTransactionID.Expected {
		t.Errorf("Output %t not equal to expected %t", result, caseTestTransactionID.Expected)
	}
}
