package currencyapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const defaultBaseURL = "https://api.currencyapi.com/v3"

type ClientRepo interface {
	GetLatestRate(ctx context.Context, baseCurrency string, currencies []string) (data RateResponse, err error)
	GetHistoricalRate(ctx context.Context, baseCurrency string, currencies []string, date time.Time) (data RateResponse, err error)
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	httpClient HttpClient
	baseURL    string
	token      string
}

// ClientOption is a function that configures a Client.
type ClientOption func(*client)

// WithHTTPClient sets the HTTP client for the Currencyapi API client.
func WithHTTPClient(c HttpClient) ClientOption {
	return func(target *client) {
		target.httpClient = c
	}
}

// WithBaseURL sets the base URL for the Currencyapi API client.
func WithBaseURL(baseURL string) ClientOption {
	return func(target *client) {
		target.baseURL = strings.TrimSuffix(baseURL, "/")
	}
}

func NewClient(token string, options ...ClientOption) *client {
	c := &client{
		httpClient: http.DefaultClient,
		baseURL:    defaultBaseURL,
		token:      token,
	}

	for _, option := range options {
		option(c)
	}

	return c
}

func (c *client) GetLatestRate(ctx context.Context, baseCurrency string, currencies []string) (data RateResponse, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/latest", c.baseURL), nil)
	if err != nil {
		return data, fmt.Errorf("cannot create request: %w", err)
	}

	q := req.URL.Query()
	q.Add("base_currency", baseCurrency)
	q.Add("currencies", strings.Join(currencies, ","))
	req.URL.RawQuery = q.Encode()

	req.Header.Set("apikey", c.token)

	return c.doRequest(ctx, req)
}

func (c *client) GetHistoricalRate(ctx context.Context, baseCurrency string, currencies []string, date time.Time) (data RateResponse, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/historical", c.baseURL), nil)
	if err != nil {
		return data, fmt.Errorf("cannot create request: %w", err)
	}

	q := req.URL.Query()
	q.Add("base_currency", baseCurrency)
	q.Add("currencies", strings.Join(currencies, ","))
	q.Add("date", date.Format("2006-01-02"))
	req.URL.RawQuery = q.Encode()

	req.Header.Set("apikey", c.token)

	return c.doRequest(ctx, req)
}

func (c *client) doRequest(ctx context.Context, req *http.Request) (data RateResponse, err error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return data, fmt.Errorf("cannot send request: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return data, fmt.Errorf("failed to decode response: %w", err)
		}
	default:
		return data, fmt.Errorf("received unexpected status code: %d", resp.StatusCode)
	}

	return data, nil
}
