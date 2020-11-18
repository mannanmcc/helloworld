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

/*
GetRate retrieve the rate and return
*/
func (env *Env) GetRate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sourceCurrency := params["sourceCurrency"]
	destinationCurrency := params["destinationCurrency"]

	redisClient := redis.NewRedis()
	redisrepo := redis.RateRedisRepository{Client: redisClient}
	rateInRedis := redisrepo.GetRate(sourceCurrency, destinationCurrency)

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
		redisrepo.SaveRate(sourceCurrency, destinationCurrency, currencyRate)
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
	rateRepository := models.RateRepository{Db: env.Db}

	rateRepository.AddRate(rate)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
