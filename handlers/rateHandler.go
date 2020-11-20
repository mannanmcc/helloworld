package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mannanmcc/helloworld/models"
	"github.com/mannanmcc/helloworld/rates"
	"github.com/mannanmcc/helloworld/redis"
)

type RateHandler struct {
	RateRepository  *models.RateRepository
	RedisRepository *redis.RateRedisRepository
}

func NewRateHandler(rateRepo *models.RateRepository, redisRepo *redis.RateRedisRepository) *RateHandler {
	return &RateHandler{
		RateRepository:  rateRepo,
		RedisRepository: redisRepo,
	}
}

func (handler *RateHandler) GetRate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sourceCurrency := params["sourceCurrency"]
	destinationCurrency := params["destinationCurrency"]

	rateInRedis := handler.RedisRepository.GetRate(sourceCurrency, destinationCurrency)

	var currencyRate float64

	if rateInRedis != "" {
		currencyRate, _ = strconv.ParseFloat(rateInRedis, 64)
		log.Println("rate found in in redis" + rateInRedis)
	} else {
		log.Println("rate not found in in redis")
		//get rate from remote api
		rateResponse := rates.GetRates(sourceCurrency, destinationCurrency)
		currencyRate = rateResponse.Rates[destinationCurrency]

		//store rate in cache
		handler.RedisRepository.SaveRate(sourceCurrency, destinationCurrency, currencyRate)
	}

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
