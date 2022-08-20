package go_bnb_sdk

import (
	"github.com/Zhenya671/go-bnb-sdk/handlers"
	_ "github.com/Zhenya671/go-bnb-sdk/handlers"
	"io"
	"log"
	"net/http"
)

const (
	defaultApiURL = "https://www.nbrb.by/api/exrates/rates/"
)

func GetCurrentCurrency(currentCurrencyID int) ([]byte, error) {
	res, err := doHttpGet(currentCurrencyID)
	if err != nil {
		log.Fatal("can't do request")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("connect with api cannot be closed")
		}
	}(res.Body)

	return gotData(res), nil
}

func gotData(res *http.Response) []byte {
	return handlers.ApiResponseToBytes(res)
}

func doHttpGet(currentCurrencyID int) (*http.Response, error) {
	return http.Get(handlers.CurrentCurrencyURL(defaultApiURL, currentCurrencyID))
}
