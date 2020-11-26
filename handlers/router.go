package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter configure all the routes with handlers responsible
func NewRouter(rateHandler *RateHandler) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/trade/book", rateHandler.BookTrade).Methods("POST")
	router.HandleFunc("/trade/rate/{sourceCurrency}/{destinationCurrency}", rateHandler.GetRate).Methods("GET")

	return router
}
