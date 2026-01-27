package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", greet)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func greet(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello, World!")
	}