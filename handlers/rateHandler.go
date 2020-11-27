package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mannanmcc/helloworld/models"
)

type RateHandler struct {
	RateRepository *models.RateRepository
	RateProvider   *RateProvider
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type PostData struct {
	SourceCurrency      string
	DestinationCurrency string
	Amount              float64
	TradeDate           time.Time
}

func JSONResponse(status int, msg string, w http.ResponseWriter) {
	response := Response{
		Message: msg,
	}
	resp, _ := json.Marshal(response)
	w.WriteHeader(status)
	w.Write(resp)
}

func NewRateHandler(rateRepo *models.RateRepository, rateProvider *RateProvider) *RateHandler {
	return &RateHandler{
		RateRepository: rateRepo,
		RateProvider:   rateProvider,
	}
}

func (handler *RateHandler) BookTrade(w http.ResponseWriter, r *http.Request) {
	var err error
	decoder := json.NewDecoder(r.Body)

	//todo validate post if it has all required data
	var postedData PostData
	err = decoder.Decode(&postedData)
	if err != nil {
		JSONResponse(http.StatusBadRequest, "Oops invalid data provided", w)
		panic(err)
		return
	}

	//send post data to be saved
	//check the date cant be more than 3 days in further, amount cant be more than 1000000
	w.Header().Set("Content-Type", "pkglication/json")
	if hasValidData(postedData) == false {
		log.Printf("%+v\n", postedData)
		JSONResponse(http.StatusBadRequest, "Oops invalid data provided", w)
		return
	}

	JSONResponse(http.StatusOK, "Your trade has been successfully booked", w)
	return
}

func hasValidData(data PostData) bool {
	if data.SourceCurrency == "" || data.DestinationCurrency == "" || data.Amount <= 0 {
		log.Println("valid failed 1")
		return false
	}

	//check if the trade is within next 3 days
	res := checkIfTradeIsWithinNextThreeDays(data.TradeDate)
	log.Println("returning final true")

	return res
}

func checkIfTradeIsWithinNextThreeDays(tradeDate time.Time) bool {
	today := time.Now()

	threeDaysInFuture := today.AddDate(0, 0, +3)

	if tradeDate.After(threeDaysInFuture) || tradeDate.Before(today) {
		log.Println("valid failed 2")
		return false
	}

	log.Println("returning true")
	return true
}

func (handler *RateHandler) GetRate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sourceCurrency := params["sourceCurrency"]
	destinationCurrency := params["destinationCurrency"]

	currencyRate := handler.RateProvider.GetRate(sourceCurrency, destinationCurrency)

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
