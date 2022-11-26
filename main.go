package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customers struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func main() {
	// define route
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customers{
		{
			Name:    "Deepayan",
			City:    "Khulna",
			ZipCode: "9100",
		},
		{
			Name:    "Rob",
			City:    "Khulna",
			ZipCode: "9100",
		},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
