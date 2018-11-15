package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func IndexView(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World!!")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", IndexView).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}