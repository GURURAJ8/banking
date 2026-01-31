package app

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"github.com/gorilla/mux"
	"github.com/GURURAJ8/banking/domain"
	"github.com/GURURAJ8/banking/Service"
)

func sanityCheck(){
	if os.Getenv("SERVER_ADDRESS") =="" || os.Getenv("SERVER_PORT")==""{
		log.Fatal("SERVER_ADDRESS or SERVER_PORT environment variable not defined")
	}
	if os.Getenv("DB_USER") =="" || os.Getenv("DB_PASS")=="" || os.Getenv("DB_ADDR")=="" || os.Getenv("DB_PORT")=="" || os.Getenv("DB_NAME")==""{
		log.Fatal("Database environment variables are not defined")
	}
}

func Start(){
	sanityCheck()
	router:=mux.NewRouter()	

	//wire up dependencies
	ch:=CustomerHandlers{Service.NewCustomerService(domain.NewCustomerRepositoryDb())}	
	//Define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomerById).Methods("GET")

	//start server
	address:=os.Getenv("SERVER_ADDRESS")
	port:=os.Getenv("SERVER_PORT")
	fmt.Println("Starting server on address:", address, "port:", port)
	log.Fatal(http.ListenAndServe(address+":"+port, router))
}

