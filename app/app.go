package app

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func Start(){
	// http.HandleFunc("/", greet)
	// http.HandleFunc("/customers", getAllCustomers)
	router:=mux.NewRouter()	

	router.HandleFunc("/", greet).Methods("GET")
	router.HandleFunc("/customers", getAllCustomers).Methods("GET")
	//url segment example
	router.HandleFunc("/customer/{customer_id:[0-9]+}", getCustomers).Methods("GET")
	router.HandleFunc("/customer", createCustomer).Methods("POST")

	//start server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func createCustomer(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Create Customer Endpoint Hit")

}

func getCustomers(w http.ResponseWriter, r *http.Request){
	params:=mux.Vars(r)
	fmt.Fprintf(w, params["customer_id"])
}