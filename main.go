package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiResponse struct {
	Base  string             `json:"base,omitempty"`
	Date  string             `json:"date,omitempty"`
	Rates map[string]float64 `json:"rates,omitempty"`
}

type responseToReturn struct {
	Source      string  `json:"source"`
	Destination string  `json:"destination"`
	Rate        float64 `json:"rate"`
}

func getCurrency(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sourceCurrency := params["sourceCurrency"]
	destinationCurrency := params["destinationCurrency"]
	url := "https://api.exchangeratesapi.io/latest?symbols=" + destinationCurrency + "&base=" + sourceCurrency
	response, err := http.Get(url)

	fmt.Println("source currency" + url)

	if err != nil {
		log.Fatal("error connrecting remote api")
		log.Fatal(err)
	}

	data, dataError := ioutil.ReadAll(response.Body)
	if dataError != nil {
		log.Fatal("error reading data from remote api")
		log.Fatal(err)
	}

	apiResponse := apiResponse{}
	err = json.Unmarshal(data, &apiResponse)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	res := &responseToReturn{
		Source:      apiResponse.Base,
		Destination: destinationCurrency,
		Rate:        apiResponse.Rates[destinationCurrency],
	}

	json.NewEncoder(w).Encode(res)
}

func main() {
	fmt.Println("Hello World!")
	//.HandleFunc("/company/{id:[0-9]+}", env.GetCompany).Methods("GET")
	router := mux.NewRouter()
	router.HandleFunc("/rate/{sourceCurrency}/{destinationCurrency}", getCurrency).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
