package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(rateHandler *RateHandler) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/rate/{sourceCurrency}/{destinationCurrency}", rateHandler.GetRate).Methods("GET")

	return router
}
