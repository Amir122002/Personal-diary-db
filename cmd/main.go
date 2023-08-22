package main

import (
	"diary_db/internal/configs"
	"diary_db/internal/handlers"
	"log"
	"net/http"
)

func main() {

	config, err := configs.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	address := config.Host + config.Port
	router := handlers.InitRouter()

	srv := http.Server{
		Addr:    address,
		Handler: router,
	}
	log.Println("start")
	err = srv.ListenAndServe()
	if err != nil {
		log.Println("listening and service error:", err)
	}
}
