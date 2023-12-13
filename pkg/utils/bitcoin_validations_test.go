package utils

import (
	"testing"
)

func TestValidateBitcoinAddress(t *testing.T) {
	validAddresses := []string{
		"1BoatSLRHtKNngkdXEeobR76b53LETtpyT",
		"3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy",
		"bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq",
	}
	for _, addr := range validAddresses {
		if !ValidateBitcoinAddress(addr) {
			t.Errorf("Address is valid but was not recognized: %s", addr)
		}
	}

	invalidAddresses := []string{
		"",
		"1234567890",
		"zzz1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq",
		"bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdqxxxxxxxx",
	}
	for _, addr := range invalidAddresses {
		if ValidateBitcoinAddress(addr) {
			t.Errorf("Address is invalid but was recognized as valid: %s", addr)
		}
	}
}

func TestValidateTransactionID(t *testing.T) {
	validTxIDs := []string{
		"4e3b9a5f8a5e3e2a8c541c5a5e7f0f98e6e8a5853d71d3fffc9ef635f1b99f63",
		"b1fea52486d0e6b4b2d07c0336a2b06ef4e58d9c4553b827bcb9b963b07bec29",
	}
	for _, txid := range validTxIDs {
		if !ValidateTransactionID(txid) {
			t.Errorf("Transaction ID is valid but was not recognized: %s", txid)
		}
	}

	invalidTxIDs := []string{
		"",
		"123",
		"g4e3b9a5f8a5e3e2a8c541c5a5e7f0f98e6e8a5853d71d3fffc9ef635f1b99f63",
		"4e3b9a5f8a5e3e2a8c541c5a5e7f0f98",
	}
	for _, txid := range invalidTxIDs {
		if ValidateTransactionID(txid) {
			t.Errorf("Transaction ID is invalid but was recognized as valid: %s", txid)
		}
	}
}
