package app

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/GURURAJ8/banking/domain"
	"github.com/GURURAJ8/banking/Service"
)

func Start(){
	router:=mux.NewRouter()	

	//wire up handlers
	// ch:=CustomerHandlers{Service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch:=CustomerHandlers{Service.NewCustomerService(domain.NewCustomerRepositoryDb())}	
	//Define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")

	//start server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

