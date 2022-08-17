package go_bnb_sdk

import (
	"errors"
	"net/http"
	"time"
)

const (
	host           = "https://www.nbrb.by/api/exrates/currencies"
	defaultTimeout = 5 * time.Second
)

type (
	addRequest struct {
		URL          string `json:"url"`
		CurrencyName string
	}

	AddInput struct {
		URL string
	}
)

func (i AddInput) validate() error {
	if i.URL == "" {
		return errors.New("required URL value is empty")
	}
	return nil
}

func (i AddInput) generateRequest(currencyName string) addRequest {
	return addRequest{
		URL:          i.URL,
		CurrencyName: currencyName,
	}
}

type Client struct {
	client       *http.Client
	currencyName string
}

func NewRequest(currencyName string) (*Client, error) {
	if currencyName == "" {
		return nil, errors.New("currency name doesn't specified")
	}

	return &Client{
		client: &http.Client{
			Timeout: defaultTimeout,
		},
		currencyName: currencyName,
	}, nil
}

func (c *Client) doHTTP() {

}
