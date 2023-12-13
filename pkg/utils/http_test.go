package utils

import (
	"bitcoin-challenge/internal/core/domain"
	"bitcoin-challenge/pkg/utils/constants"
	"context"
	"net/http"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	headers := http.Header{}
	headers.Add("Authorization", "Basic "+BasicAuth(constants.Username, constants.Password))

	//example body response
	var addressDetails domain.AddressDetails

	//mount url
	bitcoinAddress := "1P5ZEDWTKTFGxQjZphgWPQUpe554WKDfHQ"
	url := constants.URLKlever + constants.SufixUrlAddressDetails + bitcoinAddress

	//example requisition to klever api
	if err := Get(ctx, url, &addressDetails); err != nil {
		t.Errorf("Error in utils.Get: %s", err.Error())
	}

	if addressDetails.Address != bitcoinAddress {
		t.Errorf("Error to get details of bitcoin, expected address: %s; received: %s", bitcoinAddress, addressDetails.Address)
	}

}

func TestBasicAuth(t *testing.T) {
	tests := []struct {
		username string
		password string
		expected string
	}{
		{"user", "pass", "dXNlcjpwYXNz"},
		{"admin", "12345", "YWRtaW46MTIzNDU="},
	}

	for _, test := range tests {
		encoded := BasicAuth(test.username, test.password)
		expected := test.expected
		if encoded != expected {
			t.Errorf("BasicAuth(%q, %q) = %q, want %q", test.username, test.password, encoded, expected)
		}
	}
}
