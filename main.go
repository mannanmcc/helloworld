package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mannanmcc/helloworld/handlers"
	"github.com/mannanmcc/helloworld/models"

	"github.com/gorilla/mux"
)

func main() {
	db, err := models.NewDB("host=postgres user=test password=password dbname=fullstack_api port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}

	env := handlers.Env{Db: db}
	router := mux.NewRouter()

	router.HandleFunc("/rate/{sourceCurrency}/{destinationCurrency}", env.GetRate).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
