package main

import (
	"net/http"

	"github.com/mannanmcc/helloworld/config"
	"github.com/mannanmcc/helloworld/handlers"
	"github.com/mannanmcc/helloworld/models"
	"github.com/mannanmcc/helloworld/rates"
	"github.com/mannanmcc/helloworld/redis"
	"go.uber.org/dig"
)

const baseURL string = "https://api.exchangeratesapi.io/latest?symbols="

// BuildContainer builds the whole dependancy container
func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(config.NewConfig)
	container.Provide(models.NewDB)
	container.Provide(models.NewRateRepository)
	container.Provide(redis.NewRedis)
	container.Provide(redis.NewRateRedisRepository)
	container.Provide(rates.NewRateClient(baseURL))
	container.Provide(handlers.NewRateProvider)
	container.Provide(handlers.NewRateHandler)
	container.Provide(newServer)
	container.Provide(handlers.NewRouter)

	return container
}

func main() {
	container := BuildContainer()

	err := container.Invoke(func(server *Server, rateHandler http.Handler) {
		server.Run(rateHandler)
	})

	if err != nil {
		panic(err)
	}
}
