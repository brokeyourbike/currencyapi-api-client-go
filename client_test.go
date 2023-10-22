package currencyapi_test

import (
	"bytes"
	"context"
	_ "embed"
	"errors"
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/brokeyourbike/currencyapi-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/latest.json
var latestResponse []byte

//go:embed testdata/historical.json
var historicalResponse []byte

//go:embed testdata/invalid.json
var invalidResponse []byte

func TestGetLatestRate(t *testing.T) {
	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(latestResponse))}

	mockHttpClient := currencyapi.NewMockHttpClient(t)
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	client := currencyapi.NewClient("token", currencyapi.WithHTTPClient(mockHttpClient))

	got, err := client.GetLatestRate(context.TODO(), "GBP", []string{"USD"})
	require.NoError(t, err)
	assert.Len(t, got.Data, 3)
}

func TestTestGetLatestRate_InvalidResponse(t *testing.T) {
	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(invalidResponse))}

	mockHttpClient := currencyapi.NewMockHttpClient(t)
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	client := currencyapi.NewClient("token", currencyapi.WithHTTPClient(mockHttpClient))

	_, err := client.GetLatestRate(context.TODO(), "GBP", []string{"USD"})
	require.Error(t, err)
	require.Contains(t, err.Error(), "decode response")
}

func TestTestGetLatestRate_FailDuringRequest(t *testing.T) {
	mockHttpClient := currencyapi.NewMockHttpClient(t)
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{}, errors.New("error")).Once()

	client := currencyapi.NewClient("token", currencyapi.WithHTTPClient(mockHttpClient))

	_, err := client.GetLatestRate(context.TODO(), "GBP", []string{"USD"})
	require.Error(t, err)
	require.Contains(t, err.Error(), "cannot send request")
}

func TestGetHistoricalRate(t *testing.T) {
	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(historicalResponse))}

	mockHttpClient := currencyapi.NewMockHttpClient(t)
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	client := currencyapi.NewClient("token", currencyapi.WithHTTPClient(mockHttpClient))

	got, err := client.GetHistoricalRate(context.TODO(), "GBP", []string{"USD"}, time.Time{})
	require.NoError(t, err)
	assert.Len(t, got.Data, 3)
}

func TestUnexpectedStatusCode(t *testing.T) {
	resp := &http.Response{StatusCode: http.StatusBadRequest, Body: io.NopCloser(nil)}

	mockHttpClient := currencyapi.NewMockHttpClient(t)
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	client := currencyapi.NewClient("token", currencyapi.WithHTTPClient(mockHttpClient))

	_, err := client.GetLatestRate(context.TODO(), "GBP", []string{"USD"})
	require.Error(t, err)
	require.Contains(t, err.Error(), strconv.Itoa(http.StatusBadRequest))
}
