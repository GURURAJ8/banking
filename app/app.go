package app

import (
	// "fmt"
	"net/http"
	"log"
)

func Start(){
	// http.HandleFunc("/", greet)
	http.HandleFunc("/customers", getAllCustomers)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}