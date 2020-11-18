package handlers

import (
	"github.com/gorilla/mux"
)

/*
NewRouter define all the routes
*/
func NewRouter(env Env) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/rate/{sourceCurrency}/{destinationCurrency}", env.GetRate).Methods("GET")

	return router
}
