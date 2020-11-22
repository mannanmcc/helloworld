package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mannanmcc/helloworld/models"
)

// RateHandler handle the getrate and book trade request
type RateHandler struct {
	RateRepository *models.RateRepository
	RateProvider   *RateProvider
}

// NewRateHandler provide handler
func NewRateHandler(rateRepo *models.RateRepository, rateProvider *RateProvider) *RateHandler {
	return &RateHandler{
		RateRepository: rateRepo,
		RateProvider:   rateProvider,
	}
}

// GetRate handle the getrate request
func (handler *RateHandler) GetRate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sourceCurrency := params["sourceCurrency"]
	destinationCurrency := params["destinationCurrency"]

	currencyRate := handler.RateProvider.getRate(sourceCurrency, destinationCurrency)

	res := &apiResponse{
		Source:      sourceCurrency,
		Destination: destinationCurrency,
		Rate:        currencyRate,
	}

	rate := models.Rate{
		BuyCurrency:  destinationCurrency,
		SellCurrency: sourceCurrency,
		Rate:         currencyRate,
		CreatedOn:    time.Now(),
	}
	handler.RateRepository.AddRate(rate)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
