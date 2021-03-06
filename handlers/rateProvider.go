package handlers

import (
	"log"
	"strconv"

	"github.com/mannanmcc/helloworld/rates"
	"github.com/mannanmcc/helloworld/redis"
)

type RateProvider struct {
	RateRedis  *redis.RateRedisRepository
	RateClient *rates.RateClient
}

func NewRateProvider(rateRedis *redis.RateRedisRepository, rateClient *rates.RateClient) *RateProvider {
	return &RateProvider{
		RateRedis:  rateRedis,
		RateClient: rateClient,
	}
}

func (rateProder *RateProvider) GetRate(sourceCurrency string, destinationCurrency string) float64 {
	rateInRedis := rateProder.RateRedis.GetRate(sourceCurrency, destinationCurrency)

	var currencyRate float64

	if rateInRedis != "" {
		currencyRate, _ = strconv.ParseFloat(rateInRedis, 64)
		log.Println("rate found in in redis" + rateInRedis)
	} else {
		log.Println("rate not found in in redis")
		//get rate from remote api
		rateResponse := rateProder.RateClient.GetRate(sourceCurrency, destinationCurrency)
		currencyRate = rateResponse.Rates[destinationCurrency]

		//store rate in cache
		rateProder.RateRedis.SaveRate(sourceCurrency, destinationCurrency, currencyRate)
	}

	return currencyRate
}
