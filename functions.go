package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		_, _ = fmt.Fprint(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, val := range movies {
		if val.Id == params["id"] {
			err := json.NewEncoder(w).Encode(val)
			if err != nil {
				_, _ = fmt.Fprint(w, "500 Internal Server Error", http.StatusInternalServerError)
				break
			}
		}
	}
	return
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.Id = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)

	err := json.NewEncoder(w).Encode(movie)
	if err != nil {
		_, _ = fmt.Fprint(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)

			var movie Movie

			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			err := json.NewEncoder(w).Encode(movie)
			if err != nil {
				_, _ = fmt.Fprint(w, "500 Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		}
	}
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
