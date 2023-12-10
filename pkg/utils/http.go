package utils

import (
	"bitcoin-challenge/pkg/logger"
	"bitcoin-challenge/pkg/utils/constants"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type HttpMethod string

const (
	MethodGet     HttpMethod = "GET"
	MethodHead    HttpMethod = "HEAD"
	MethodPost    HttpMethod = "POST"
	MethodPut     HttpMethod = "PUT"
	MethodPatch   HttpMethod = "PATCH" // RFC 5789
	MethodDelete  HttpMethod = "DELETE"
	MethodConnect HttpMethod = "CONNECT"
	MethodOptions HttpMethod = "OPTIONS"
	MethodTrace   HttpMethod = "TRACE"
)

type Options interface {
	GetHeadersRequest() http.Header
}

var client http.Client

func Get(ctx context.Context, url string, method HttpMethod, responseBody interface{}, options ...Options) error {
	req, err := http.NewRequestWithContext(ctx, string(method), url, nil)
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

	err = json.Unmarshal(bodyRes, responseBody)
	if err != nil {
		logger.Error(constants.ErrorToReadResponseBody, err)
		return err
	}

	return nil
}

func Post() {

}
