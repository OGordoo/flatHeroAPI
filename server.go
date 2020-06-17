package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Superhero struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Intelligence    int    `json:"intelligence"`
	Strength        int    `json:"strength"`
	Speed           int    `json:"speed"`
	Durability      int    `json:"durability"`
	Power           int    `json:"power"`
	Combat          int    `json:"combat"`
	Gender          string `json:"gender"`
	Race            string `json:"race"`
	Height          string `json:"height"`
	Weight          string `json:"weight"`
	FullName        string `json:"fullName"`
	PlaceOfBirth    string `json:"placeOfBirth"`
	FirstAppearance string `json:"firstAppearance"`
	Publisher       string `json:"publisher"`
	Alignment       string `json:"alignment"`
	Sm              string `json:"sm"`
}

var superheroes []Superhero

func handleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// all Get handlers
func getHeroes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(superheroes)
}

func getJSON(jsonfile string) {
	jsonFile, err := os.Open(jsonfile)
	handleError(err)

	// read our opened xml/json as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)
	handleError(err)

	err = json.Unmarshal(byteValue, &superheroes)
	handleError(err)
	fmt.Println(superheroes)

	jsonFile.Close()
}

func main() {
	r := mux.NewRouter()

	getJSON("./superheroAPI.json")

	fmt.Println()

	//Route Handlers / Endpoints
	r.HandleFunc("/api/superhero", getHeroes).Methods("GET")

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
