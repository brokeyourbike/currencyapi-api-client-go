package currencyapi_test

import (
	"encoding/json"
	"testing"

	"github.com/brokeyourbike/currencyapi-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRateResponse(t *testing.T) {
	var data currencyapi.RateResponse

	assert.NoError(t, json.Unmarshal(latestResponse, &data))
	assert.Len(t, data.Data, 3)
	assert.Contains(t, data.Data, "EUR")
	assert.Contains(t, data.Data, "CAD")
	assert.Contains(t, data.Data, "USD")
	assert.Equal(t, 0.9104501416, data.Data["EUR"].Value)
}

func TestNewRateResponse(t *testing.T) {
	resp := currencyapi.NewRateResponse(map[string]float64{"USD": 1.23})
	require.Len(t, resp.Data, 1)

	v, ok := resp.Data["USD"]
	require.True(t, ok)
	require.Equal(t, v.Code, "USD")
	require.Equal(t, v.Value, 1.23)
}
