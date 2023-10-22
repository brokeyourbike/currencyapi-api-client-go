package currencyapi

type RateResponseDataItem struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}

type RateResponse struct {
	Data map[string]RateResponseDataItem `json:"data"`
}

func NewRateResponse(data map[string]float64) RateResponse {
	resp := RateResponse{Data: make(map[string]RateResponseDataItem)}
	for k, v := range data {
		resp.Data[k] = RateResponseDataItem{Code: k, Value: v}
	}
	return resp
}
