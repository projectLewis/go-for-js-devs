package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

// Planet is a planet from the star wars universe
type Planet struct {
	Name       string `json:"name"`
	Terrain    string `json:"terrain"`
	Population string `json:"population"`
}

// Person is an individual person
type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld    Planet
}

// AllPeople is a collection of "Persons"
type AllPeople struct {
	People []Person `json:"results"`
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(BaseURL + "people")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("failed request to Star Wars API")
	}
	fmt.Println(res)

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("failed to parse request body")
	}

	fmt.Fprintln(w, string(bytes))

	var people AllPeople
	if err := json.Unmarshal(bytes, &people); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("unmarshall error")
	}
	fmt.Println(people)
}

func main() {
	http.HandleFunc("/people", getPeople)
	fmt.Printf("Server is running on port: %d", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
