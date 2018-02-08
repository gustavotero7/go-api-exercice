package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	error := http.ListenAndServe(":2000", router)
	log.Fatal(error)
}
