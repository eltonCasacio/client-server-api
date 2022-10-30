package pkg

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type USDBRL struct {
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
}

type AutoGenerated struct {
	Usdbrl USDBRL `json:"USDBRL"`
}

func BuscaCotacao() (*USDBRL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	autoGenerated := AutoGenerated{}
	err = json.Unmarshal(response, &autoGenerated)
	if err != nil {
		return nil, err
	}
	var usdbrl USDBRL = autoGenerated.Usdbrl
	return &usdbrl, nil
}
