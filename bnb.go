package go_bnb_sdk

import (
	"encoding/json"
	"fmt"
	_ "github.com/Zhenya671/go-bnb-sdk/handlers"
	"io"
	"log"
	"net/http"
	"strconv"
)

const (
	defaultApiURL = "https://www.nbrb.by/api/exrates/rates/"
)

type Currency struct {
	CurID           int     `json:"Cur_ID"`
	Date            string  `json:"Date"`
	CurAbbreviation string  `json:"Cur_Abbreviation"`
	CurScale        int     `json:"Cur_Scale"`
	CurName         string  `json:"Cur_Name"`
	CurOfficialRate float64 `json:"Cur_OfficialRate"`
}

func GetCurrentCurrency(currentCurrencyID int) (map[string]interface{}, error) {
	id := strconv.Itoa(currentCurrencyID)
	res, err := http.Get(defaultApiURL + id)
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("connect with api cannot be closed")
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d - %s: %s", res.StatusCode, res.Status, "somthing wrong with connect")
	}

	var gotCurrency map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&gotCurrency)
	if err != nil {
		return nil, fmt.Errorf("could not decode json: %s\n", err)
	}

	return gotCurrency, nil
}
