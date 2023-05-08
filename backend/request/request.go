package request

import (
	"context"
	"encoding/json"
	"io"
	"jorge/env"
	"net/http"
)

type DolarBrlApiResponse struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func GetLastQuote() string {
	ctx, cancel := context.WithTimeout(context.Background(), env.EXTERNAL_API_TIMEOUT)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	respData, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var apiResponse DolarBrlApiResponse

	err = json.Unmarshal(respData, &apiResponse)

	if err != nil {
		panic(err)
	}

	return apiResponse.Usdbrl.Bid
}
