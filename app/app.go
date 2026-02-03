package app

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"github.com/gorilla/mux"
	"github.com/GURURAJ8/banking/domain"
	"github.com/GURURAJ8/banking/Service"
	"time"
	"github.com/jmoiron/sqlx"
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
	dbClient:=getDbClient()

	customerRepositoryDb:=domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb:=domain.NewAccountRepositoryDb(dbClient)

	//wire up dependencies
	ch:=CustomerHandlers{Service.NewCustomerService(customerRepositoryDb)}	
	ah:=AccountHandler{Service.NewAccountService(accountRepositoryDb)}

	//Define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomerById).Methods("GET")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/accounts", ah.NewAccount).Methods("POST")

	//start server
	address:=os.Getenv("SERVER_ADDRESS")
	port:=os.Getenv("SERVER_PORT")
	fmt.Println("Starting server on address:", address, "port:", port)
	log.Fatal(http.ListenAndServe(address+":"+port, router))
}

func getDbClient() *sqlx.DB {
	dbUser :=os.Getenv("DB_USER")
	dbPass :=os.Getenv("DB_PASS")
	dbAddr :=os.Getenv("DB_ADDR")
	dbPort :=os.Getenv("DB_PORT")
	dbName :=os.Getenv("DB_NAME")
	client, err := sqlx.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbAddr+":"+dbPort+")/"+dbName)
    if err != nil {
	    panic(err)
    }
// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)	
return client
}
