package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	//"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

var movies = Movies{Movie{"Grinch", "Some Guy", 2000, "QWERT WERTyuaids sdf"}, Movie{"Wonder Cat", "Some Guy ^2", 2011, "QWERT WERTyuaids sdf"}}
var collection = getSession().DB("MovieStore").C("movies")

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

// Index _
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from GO!, %q %s", html.EscapeString(r.URL.Path), " //meh")
}

// ListMovies _
func ListMovies(w http.ResponseWriter, r *http.Request) {

	var results []Movie
	err := collection.Find(nil).All(&results)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print("Se han obtenido los registros de la BD ", results)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)

}

// GetMovie _
func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	if id > len(movies)-1 {
		fmt.Fprintf(w, "Lo sentimos, la pelicula solicitada no ha podido ser encontrada")
	} else {
		json.NewEncoder(w).Encode(movies[id])
	}
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

	err = collection.Insert(movieData)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movieData)
	w.WriteHeader(200)
	//movies = append(movies, movieData)
}
