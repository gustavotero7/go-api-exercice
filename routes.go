package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route _
type Route struct {
	Name    string
	Method  string
	Pattern string
	handler http.HandlerFunc
}

// Routes _
type Routes []Route

// NewRouter _
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, r := range routes {
		router.Methods(r.Method).Name(r.Name).Path(r.Pattern).HandlerFunc(r.handler)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"List Movies",
		"GET",
		"/list-movies",
		ListMovies,
	},
	Route{
		"Get Movie",
		"GET",
		"/get-movie/{id}",
		GetMovie,
	},
	Route{
		"Add Movie",
		"POST",
		"/add-movie",
		AddMovie,
	},
}
