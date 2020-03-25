package main

import (
	"log"
	"net/http"

	. "../handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server Starting...")

	r := mux.NewRouter()

	r.HandleFunc("/api/products", GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/products/{id}", GetProductHandler).Methods("GET")
	r.HandleFunc("/api/products", PostProductHandler).Methods("POST")
	r.HandleFunc("/api/products/{id}", PutProductHandler).Methods("PUT")
	r.HandleFunc("/api/products/{id}", DeleteProductHandler).Methods("Delete")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	server.ListenAndServe()

	log.Println("Server Ending...")

}
