package utils

import (
	"encoding/hex"
	"regexp"
)

func ValidateBitcoinAddress(address string) bool {
	if len(address) < 26 || len(address) > 42 {
		return false
	}

	validPrefix := regexp.MustCompile("^(1|3|bc1)")
	if !validPrefix.MatchString(address) {
		return false
	}

	validChars := regexp.MustCompile("^[A-Za-z0-9]+$")
	if !validChars.MatchString(address) {
		return false
	}

	return true
}

func ValidateTransactionID(txidHex string) bool {
	if len(txidHex) != 64 {
		return false
	}

	_, err := hex.DecodeString(txidHex)
	if err != nil {
		return false
	}

	return true
}
