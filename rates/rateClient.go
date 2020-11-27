package rates

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RateResponse struct {
	Base  string             `json:"base,omitempty"`
	Date  string             `json:"date,omitempty"`
	Rates map[string]float64 `json:"rates,omitempty"`
}

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type RateClient struct {
	Client  HttpClient
	baseURL string
}

func NewRateClient(url string) *RateClient {
	return &RateClient{
		Client:  &http.Client{},
		baseURL: url,
	}
}

func (rateClient *RateClient) GetRate(sourceCurrency string, destinationCurrency string) RateResponse {
	url := rateClient.baseURL + destinationCurrency + "&base=" + sourceCurrency
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
