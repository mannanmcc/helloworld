package redis

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

type RateRedisRepository struct {
	Client redis.Conn
}

func NewRateRedisRepository(redisClient redis.Conn) *RateRedisRepository {
	return &RateRedisRepository{
		Client: redisClient,
	}
}

func (repo RateRedisRepository) SaveRate(sourceCurrency string, destinationCurrency string, rate float64) error {
	key := sourceCurrency + destinationCurrency
	_, err := repo.Client.Do("SET", key, rate)

	if err != nil {
		log.Println("Oops, could not store rate in the cache")
		return err
	}

	return nil
}

func (repo RateRedisRepository) GetRate(sourceCurrency string, destinationCurrency string) string {
	key := sourceCurrency + destinationCurrency
	result, err := redis.String(repo.Client.Do("GET", key))

	if err != nil {
		log.Println("nothing found in the cache")
		return ""
	}

	return result
}
