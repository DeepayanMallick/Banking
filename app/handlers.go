package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customers struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
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

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from post method")
}
