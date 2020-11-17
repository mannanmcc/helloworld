package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mannanmcc/helloworld/models"
	"github.com/mannanmcc/helloworld/rates"
)

/*
GetRate retrieve the rate and return
*/
func (env *Env) GetRate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sourceCurrency := params["sourceCurrency"]
	destinationCurrency := params["destinationCurrency"]
	rateResponse := rates.GetRates(sourceCurrency, destinationCurrency)
	res := &apiResponse{
		Source:      sourceCurrency,
		Destination: destinationCurrency,
		Rate:        rateResponse.Rates[destinationCurrency],
	}

	rate := models.Rate{
		BuyCurrency:  destinationCurrency,
		SellCurrency: sourceCurrency,
		Rate:         rateResponse.Rates[destinationCurrency],
		CreatedOn:    time.Now(),
	}
	rateRepository := models.RateRepository{Db: env.Db}

	rateRepository.AddRate(rate)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
