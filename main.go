package main

import (
	"github.com/mannanmcc/helloworld/config"
	"github.com/mannanmcc/helloworld/handlers"
	"github.com/mannanmcc/helloworld/models"
	"github.com/mannanmcc/helloworld/redis"
)

func main() {
	config := config.NewConfig()

	db, err := models.NewDB(config)
	if err != nil {
		panic(err)
	}

	repo := models.NewRateRepository(db)

	redisClient := redis.NewRedis()
	redisrepo := redis.NewRateRedisRepository(redisClient)

	rateHandler := handlers.NewRateHandler(repo, redisrepo)

	server := newServer(config)
	router := handlers.NewRouter(rateHandler)
	server.Run(router)
}
