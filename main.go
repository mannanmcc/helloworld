package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	rates "github.com/mannanmcc/helloworld/rates"

	"github.com/gorilla/mux"
)

type apiResponse struct {
	Source      string  `json:"source"`
	Destination string  `json:"destination"`
	Rate        float64 `json:"rate"`
}

func getCurrency(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sourceCurrency := params["sourceCurrency"]
	destinationCurrency := params["destinationCurrency"]
	rateResponse := rates.GetRates(sourceCurrency, destinationCurrency)
	res := &apiResponse{
		Source:      sourceCurrency,
		Destination: destinationCurrency,
		Rate:        rateResponse.Rates[destinationCurrency],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/rate/{sourceCurrency}/{destinationCurrency}", getCurrency).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
