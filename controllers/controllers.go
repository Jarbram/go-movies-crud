package controllers

import (
	"fmt"
	"golangProjects/go-movies-crud/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", services.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", services.GetMovie).Methods("GET")
	r.HandleFunc("/movies", services.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", services.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", services.DeleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
