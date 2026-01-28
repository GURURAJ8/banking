package app

import (
	// "fmt"
	"net/http"
	"encoding/json"
	"encoding/xml"
	"strconv"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/GURURAJ8/banking/Service"
)

//Data Transfer Object
type Customer struct{
	customer_id int `json:"customer_id" xml:"customer_id"`
	Name string `json:"name" xml:"name"`
	Zip string `json:"zip" xml:"zip"`
	City string `json:"city" xml:"city"`
}

type CustomerHandlers struct{
	service Service.CustomerService
}
//handler functions
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request){
		// customer:= []Customer{
		// 	{customer_id: 1, Name: "John Doe", Zip: "12345", City: "New York"},
		// 	{customer_id: 2, Name: "Jane Smith", Zip: "67890", City: "Los Angeles"},
		// 	{customer_id: 3, Name: "Alice Johnson", Zip: "54321", City: "Chicago"},	
		// }
	customer, err:= ch.service.GetAllCustomers()
	if err!=nil{
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
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

func (ch *CustomerHandlers) getCustomerById(w http.ResponseWriter, r *http.Request){
	// Implementation for getting a customer by ID would go here
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}
	customer, err := ch.service.GetCustomerById(id)
	if err != nil {
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		w.WriteHeader(err.Code)
		fmt.Fprintf(w, err.Message)
	}else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
		w.WriteHeader(http.StatusOK)
	}	
}