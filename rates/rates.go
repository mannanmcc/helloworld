package rates

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL = "https://api.exchangeratesapi.io/latest?symbols="

/*
RateResponse give structure of api respinse
*/
type RateResponse struct {
	Base  string             `json:"base,omitempty"`
	Date  string             `json:"date,omitempty"`
	Rates map[string]float64 `json:"rates,omitempty"`
}

/**
GetRates return rates from api
*/
func GetRates(sourceCurrency string, destinationCurrency string) RateResponse {
	url := baseURL + destinationCurrency + "&base=" + sourceCurrency
	response, err := http.Get(url)

	if err != nil {
		log.Fatal("error connrecting remote api")
		log.Fatal(err)
	}

	data, dataError := ioutil.ReadAll(response.Body)
	if dataError != nil {
		log.Fatal("error reading data from remote api")
		log.Fatal(err)
	}

	rateResponse := RateResponse{}
	err = json.Unmarshal(data, &rateResponse)
	if err != nil {
		log.Fatal(err)
	}

	return rateResponse
}
