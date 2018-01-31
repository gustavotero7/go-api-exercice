package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var movies Movies = Movies{Movie{"Grinch", "Some Guy", 2000, "QWERT WERTyuaids sdf"}, Movie{"Wonder Cat", "Some Guy ^2", 2011, "QWERT WERTyuaids sdf"}}

// Index _
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from GO!, %q %s", html.EscapeString(r.URL.Path), " //meh")
}

// ListMovies _
func ListMovies(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(movies)
}

// GetMovie _
func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprintf(w, "Movie #%s", params["id"])
}

// AddMovie _
func AddMovie(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movieData Movie
	err := decoder.Decode(&movieData)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	log.Println(movieData)
	json.NewEncoder(w).Encode(movieData)
	movies = append(movies, movieData)
}
