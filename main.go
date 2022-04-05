package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var movies []Movie

func main() {
	// Router
	r := mux.NewRouter()

	movies = append(movies, Movie{"1", "43227", "Spider-man", &Director{"Andrew", "I dont know"}})
	movies = append(movies, Movie{"2", "43227", "Spider-man", &Director{"Andrew", "I dont know"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	log.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
