package utils

import (
	"bitcoin-challenge/pkg/logger"
	"bitcoin-challenge/pkg/utils/constants"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
)

type Options interface {
	GetHeadersRequest() http.Header
}

type CustomHeaders struct {
	headers http.Header
}

func (c CustomHeaders) GetHeadersRequest() http.Header {
	return c.headers
}

func NewCustomHeaders(headers http.Header) *CustomHeaders {
	return &CustomHeaders{headers: headers}
}

var client http.Client

func Get(ctx context.Context, url string, responseBody interface{}, options ...Options) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		logger.Error(constants.ErrorToCreateRequest, err)
		return err
	}

	for _, option := range options {
		for key, values := range option.GetHeadersRequest() {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}
	req.SetBasicAuth(constants.Username, constants.Password)

	res, err := client.Do(req)
	defer res.Body.Close()

	if err != nil {
		logger.Error(constants.ErrorToMakeRequest, err)
		return err
	}

	bodyRes, err := io.ReadAll(res.Body)

	if err != nil {
		logger.Error(constants.ErrorToReadResponseBody, err)
		return err
	}

	err = json.Unmarshal(bodyRes, &responseBody)
	if err != nil {
		logger.Error(constants.ErrorToReadResponseBody, err)
		return err
	}

	return nil
}

func Post() {

}

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
