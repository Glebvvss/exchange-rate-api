package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
)

type JsonRate struct {
	Ccy string `json:"ccy"`
	Buy string `json:"buy"`
}

func GetUsdRate() (float32, error) {
	resp, err := http.Get(os.Getenv("RATE_API_URL"))
	defer resp.Body.Close()

	if err != nil || resp.StatusCode != http.StatusOK {
		return 0, err
	}

	var jsonRates []JsonRate
	bodyBytes, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal([]byte(bodyBytes), &jsonRates); err != nil {
		return 0, err
	}

	var usdJsonRate JsonRate
	for _, r := range jsonRates {
		if r.Ccy == "USD" {
			usdJsonRate = r
			break
		}
	}

	result, convErr := strconv.ParseFloat(usdJsonRate.Buy, 32)
	if convErr != nil {
		return 0, err
	}

	return float32(result), nil
}
