package rates

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL = "https://api.exchangeratesapi.io/latest?symbols="

type RateResponse struct {
	Base  string             `json:"base,omitempty"`
	Date  string             `json:"date,omitempty"`
	Rates map[string]float64 `json:"rates,omitempty"`
}

func GetRates(sourceCurrency string, destinationCurrency string) RateResponse {
	url := baseURL + destinationCurrency + "&base=" + sourceCurrency
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, dataError := ioutil.ReadAll(response.Body)
	if dataError != nil {
		log.Fatal(err)
	}

	rateResponse := RateResponse{}
	err = json.Unmarshal(data, &rateResponse)
	if err != nil {
		log.Fatal(err)
	}

	return rateResponse
}
