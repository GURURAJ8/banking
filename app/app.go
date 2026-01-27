package app

import (
	// "fmt"
	"net/http"
	"encoding/json"
	"log"
	"encoding/xml"
	"github.com/GURURAJ8/banking/app"
)

func start(){
	// http.HandleFunc("/", greet)
	http.HandleFunc("/customers", handlers.getAllCustomers)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}