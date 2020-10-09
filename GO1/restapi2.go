package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	// add our banguns route and map it to our
	// terimaData function like so
	http.HandleFunc("/index", terimaData)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

type Bangun struct {
	Panjang string `json:"p"`
	Lebar   string `json:"l"`
	Tinggi  string `json:"t"`
}

func main() {
	handleRequests()
}

func terimaData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: terimaData")
	json.NewEncoder(w).Encode(Articles)
}
