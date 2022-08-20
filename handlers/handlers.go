package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func CurrentCurrencyURL(apiURL string, currentCurrencyID int) string {
	return fmt.Sprintf("%s/%s", apiURL, IntToString(currentCurrencyID))
}

func IntToString(int int) string {
	return strconv.Itoa(int)
}

func ApiResponseToBytes(res *http.Response) []byte {
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
