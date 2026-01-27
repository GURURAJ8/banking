package main

import (
	// "fmt"
	"net/http"
	"encoding/json"
	"log"
)
type Customer struct{
	Name string `json:"name"`
	Zip string `json:"zip"`
	City string `json:"city"`
}

//handler 

func main(){
	// http.HandleFunc("/", greet)
	http.HandleFunc("/customers", getAllCustomers)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func getAllCustomers(w http.ResponseWriter, r *http.Request){
		customer:= []Customer{
			{Name: "John Doe", Zip: "12345", City: "New York"},
			{Name: "Jane Smith", Zip: "67890", City: "Los Angeles"},
			{Name: "Alice Johnson", Zip: "54321", City: "Chicago"},	
		}
	json.NewEncoder(w).Encode(customer)
	w.Header().Set("Content-Type", "application/json")
	}