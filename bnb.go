package go_bnb_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

const (
	apiURL = "https://www.nbrb.by/api/exrates/currencies/"
)

var ErrFailedAPICall = errors.New("bad response from BookBeta API")

type Currency struct {
	CurID           int    `json:"Cur_ID"`
	CurParentID     int    `json:"Cur_ParentID"`
	CurCode         string `json:"Cur_Code"`
	CurAbbreviation string `json:"Cur_Abbreviation"`
	CurName         string `json:"Cur_Name"`
	CurNameBel      string `json:"Cur_Name_Bel"`
	CurNameEng      string `json:"Cur_Name_Eng"`
	CurQuotName     string `json:"Cur_QuotName"`
	CurQuotNameBel  string `json:"Cur_QuotName_Bel"`
	CurQuotNameEng  string `json:"Cur_QuotName_Eng"`
	CurNameMulti    string `json:"Cur_NameMulti"`
	CurNameBelMulti string `json:"Cur_Name_BelMulti"`
	CurNameEngMulti string `json:"Cur_Name_EngMulti"`
	CurScale        int    `json:"Cur_Scale"`
	CurPeriodicity  int    `json:"Cur_Periodicity"`
	CurDateStart    string `json:"Cur_DateStart"`
	CurDateEnd      string `json:"Cur_DateEnd"`
}

func GetCurrentCurrency(currentCurrencyID int) (map[string]interface{}, error) {
	id := strconv.Itoa(currentCurrencyID)
	res, err := http.Get(apiURL + id)
	if err != nil {
		log.Fatal(ErrFailedAPICall)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("connect with api cannot be closed")
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d - %s: %w", res.StatusCode, res.Status, ErrFailedAPICall)
	}

	var gotCurrency map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&gotCurrency)
	if err != nil {
		return nil, fmt.Errorf("could not decode json: %s\n", err)
	}

	return gotCurrency, nil
}
