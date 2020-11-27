package rates

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	mock_rates "github.com/mannanmcc/helloworld/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHttpClientCanRetrieveRateAndParse(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockedHttpClient := mock_rates.NewMockHttpClient(mockController)

	baseUrl := "/api/url"
	sourceCurrency := "USD"
	destinationCurrency := "GBP"
	rate := 1.3351434438
	apiEndPointFullUrl := baseUrl + destinationCurrency + "&base=" + sourceCurrency
	mockedHttpClient.EXPECT().Get(apiEndPointFullUrl).Return(getMockedResponse(rate, sourceCurrency, destinationCurrency))

	RateClient := &RateClient{Client: mockedHttpClient, baseURL: baseUrl}
	actualRate := RateClient.GetRate(sourceCurrency, destinationCurrency)

	assert.Equal(t, actualRate.Base, destinationCurrency)
	assert.Equal(t, actualRate.Rates[sourceCurrency], 1.3351434438)
}

func getMockedResponse(rate float64, currency string, baseCurrency string) (*http.Response, error) {
	rateResponse := RateResponse{
		Rates: map[string]float64{
			currency: rate,
		},
		Base: baseCurrency,
		Date: "2020-11-26",
	}
	b, _ := json.Marshal(rateResponse)

	return &http.Response{Status: "200 OK", Body: ioutil.NopCloser(bytes.NewBuffer(b))}, nil
}
