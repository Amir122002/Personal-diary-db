package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/creat", Create).Methods(http.MethodPost)
	router.HandleFunc("/read", Read).Methods(http.MethodGet)
	router.HandleFunc("/update/{id}", Update).Methods(http.MethodPut)
	router.HandleFunc("/delete/{id}", Delete).Methods(http.MethodPut)

	return router
}
