package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ifat420/_goBookApi/pkg/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", handlers.AddBooks).Methods(http.MethodPost)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
