package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mannanmcc/helloworld/handlers"
	"github.com/mannanmcc/helloworld/models"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := models.NewDB(dsn)
	if err != nil {
		panic(err)
	}

	env := handlers.Env{Db: db}
	router := handlers.NewRouter(env)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
