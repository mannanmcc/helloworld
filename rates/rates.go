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

type RateClient struct {
	Client *http.Client
}

// NewRateClient create a new
func NewRateClient() *RateClient {
	return &RateClient{
		Client: &http.Client{},
	}
}

func (rateClient *RateClient) GetRates(sourceCurrency string, destinationCurrency string) RateResponse {
	url := baseURL + destinationCurrency + "&base=" + sourceCurrency
	response, err := rateClient.Client.Get(url)

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
