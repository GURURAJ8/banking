package app

import (
	// "fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func Start(){
	// http.HandleFunc("/", greet)
	// http.HandleFunc("/customers", getAllCustomers)
	router:=mux.NewRouter()
	router.HandleFunc("/customers", getAllCustomers).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}