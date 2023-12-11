package utils

import (
	"bitcoin-challenge/internal/core/domain"
	"bitcoin-challenge/pkg/logger"
	"bitcoin-challenge/pkg/utils/constants"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
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

func Post(ctx context.Context, url string, requestBody interface{}, responseBody interface{}, options ...Options) error {
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		logger.Error(constants.ErrorToCreateRequest, err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Error(constants.ErrorToCreateRequest, err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	for _, option := range options {
		for key, values := range option.GetHeadersRequest() {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}

	res, err := client.Do(req)
	if err != nil {
		logger.Error(constants.ErrorToMakeRequest, err)
		return err
	}
	defer res.Body.Close()

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
func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func HTTPFail(ctx *fiber.Ctx, code int, err error, message string) error {
	errJson, _ := json.Marshal(err)

	result := &domain.HTTPErrorResponse{
		Error:   errJson,
		Message: message,
	}

	if err != nil {
		result.ErrorMessage = err.Error()
	}

	return ctx.Status(code).JSON(result)
}
