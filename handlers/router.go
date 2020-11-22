package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(rateHandler *RateHandler) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/trade/{sourceCurrency}/{destinationCurrency}", rateHandler.GetRate).Methods("GET")
	router.HandleFunc("/tade/book", rateHandler.GetRate).Methods("POST")

	return router
}
