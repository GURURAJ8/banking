package app

import (
	// "fmt"
	"net/http"
	"encoding/json"
	"log"
)

//Data Transfer Object
type Customer struct{
	Name string `json:"name" xml:"name"`
	Zip string `json:"zip" xml:"zip"`
	City string `json:"city" xml:"city"`
}

//handler functions
func getAllCustomers(w http.ResponseWriter, r *http.Request){
		customer:= []Customer{
			{Name: "John Doe", Zip: "12345", City: "New York"},
			{Name: "Jane Smith", Zip: "67890", City: "Los Angeles"},
			{Name: "Alice Johnson", Zip: "54321", City: "Chicago"},	
		}

	if r.Header.Get("Content-Type") == "application/xml" {
		//Marshal to XML and write to response
		xml.NewEncoder(w).Encode(customer)
		w.Header().Set("Content-Type", "application/xml")
		return
	}else{
		// Default to JSON if no Content-Type is specified
		json.NewEncoder(w).Encode(customer)
		w.Header().Set("Content-Type", "application/json")
		return
	}	
}