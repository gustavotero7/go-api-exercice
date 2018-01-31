package main

// Movie _
type Movie struct {
	Name        string `json:"name"`
	Director    string `json:"director"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}

// Movies _
type Movies []Movie
