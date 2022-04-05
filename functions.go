package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		_, _ = fmt.Fprint(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {}

func createMovie(w http.ResponseWriter, r *http.Request) {}

func updateMovie(w http.ResponseWriter, r *http.Request) {
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range movies {

		if item.Id == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}

}
