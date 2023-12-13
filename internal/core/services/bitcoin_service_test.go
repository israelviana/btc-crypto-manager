package services

import (
	"bitcoin-challenge/internal/core/services/bitcoin"
	"bitcoin-challenge/internal/core/services/klever"
	"testing"
)

func TestFindDetailsPerAddress(t *testing.T) {
	kleverService := klever.NewKleverService()
	bitcoinService := bitcoin.NewBitcoinService(kleverService)

	type TestBitcoinAddress struct {
		BitcoinAddress string
		Expected       bool
	}

	casesTestsBitcoinAddress := TestBitcoinAddress{
		BitcoinAddress: "bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r",
		Expected:       true,
	}

	detailsPerAddress, err := bitcoinService.FindDetailsPerAddress(casesTestsBitcoinAddress.BitcoinAddress)
	if err != nil {
		t.Errorf("Error to find details per address test: %s", err.Error())
	}

	result := detailsPerAddress.Address != ""

	if result != casesTestsBitcoinAddress.Expected {
		t.Errorf("Output %t not equal to expected %t", result, casesTestsBitcoinAddress.Expected)
	}

}

func TestFindBalancePerAddress(t *testing.T) {
	kleverService := klever.NewKleverService()
	bitcoinService := bitcoin.NewBitcoinService(kleverService)

	type TestBitcoinAddress struct {
		BitcoinAddress string
		Expected       bool
	}

	casesTestsBitcoinAddress := TestBitcoinAddress{
		BitcoinAddress: "bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r",
		Expected:       true,
	}

	detailsPerAddress, err := bitcoinService.FindBalancePerAddress(casesTestsBitcoinAddress.BitcoinAddress)
	if err != nil {
		t.Errorf("Error to find balance per address test: %s", err.Error())
	}

	result := detailsPerAddress.Confirmed != "" || detailsPerAddress.Unconfirmed != ""

	if result != casesTestsBitcoinAddress.Expected {
		t.Errorf("Output %t not equal to expected %t", result, casesTestsBitcoinAddress.Expected)
	}

}

func TestMountUTXO(t *testing.T) {
	kleverService := klever.NewKleverService()
	bitcoinService := bitcoin.NewBitcoinService(kleverService)

	type TestUTXOS struct {
		BitcoinAddress string
		AmountNeeded   string
		Expected       bool
	}

	casesTestsUTXOS := TestUTXOS{
		BitcoinAddress: "1P5ZEDWTKTFGxQjZphgWPQUpe554WKDfHQ",
		AmountNeeded:   "13243",
		Expected:       true,
	}

	utxos, err := bitcoinService.MountUTXO(casesTestsUTXOS.BitcoinAddress, casesTestsUTXOS.AmountNeeded)
	if err != nil {
		t.Errorf("Error to find balance per address test: %s", err.Error())
	}

	result := len(utxos.UTXOs) > 0

	if result != casesTestsUTXOS.Expected {
		t.Errorf("Output %t not equal to expected %t", result, casesTestsUTXOS.Expected)
	}

}

func TestFindDetailsPerTransactionId(t *testing.T) {
	kleverService := klever.NewKleverService()
	bitcoinService := bitcoin.NewBitcoinService(kleverService)

	type TestDetailsPerTransactionId struct {
		TransactionID string
		Expected      bool
	}

	caseTestUTXOS := TestDetailsPerTransactionId{
		TransactionID: "6b5ae35d752ceff9618d6850a453719464d63ffdfa17332a0bd61db361909394",
		Expected:      true,
	}

	details, err := bitcoinService.FindDetailsPerTransactionId(caseTestUTXOS.TransactionID)
	if err != nil {
		t.Errorf("Error to find balance per address test: %s", err.Error())
	}

	result := details.TxID != ""

	if result != caseTestUTXOS.Expected {
		t.Errorf("Output %t not equal to expected %t", result, caseTestUTXOS.Expected)
	}

}
