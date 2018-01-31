package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	error := http.ListenAndServe(":8080", router)
	log.Fatal(error)
}
