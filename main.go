package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/cliente", GetClientes).Methods("GET")
	r.HandleFunc("/cliente/{id}", GetCliente).Methods("GET")
	r.HandleFunc("/cliente", CreateCliente).Methods("POST")
	r.HandleFunc("/cliente/{id}", UpdateCliente).Methods("PUT")
	r.HandleFunc("/cliente/{id}", DeleteCliente).Methods("DELETE")

	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui/")))
	r.PathPrefix("/swaggerui/").Handler(sh)

	err := http.ListenAndServe(":8000", r);
    if err != nil {
        log.Fatal(err)
    }
}


func main() {
	InitialMigration()
	initializeRouter()
}
