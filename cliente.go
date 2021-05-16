package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "sqlserver://delgado:xxx@localhost?database=CadastroDB"

type Cliente struct {
	Nome string `json:"nome"`
	Cidade  string `json:"cidade"`
	Email     string `json:"email"`
	Sexo     string `json:"sexo"`
}

func InitialMigration() {
	DB, err = gorm.Open(sqlserver.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	//DB.AutoMigrate(&Cliente{})
}

func GetClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Clientes []Cliente
	DB.Find(&Clientes)
	json.NewEncoder(w).Encode(Clientes)
}

func GetCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var Cliente Cliente
	DB.First(&Cliente, params["id"])
	json.NewEncoder(w).Encode(Cliente)
}

func CreateCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Cliente Cliente
	json.NewDecoder(r.Body).Decode(&Cliente)
	DB.Create(&Cliente)
	json.NewEncoder(w).Encode(Cliente)
}

func UpdateCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var Cliente Cliente
	DB.First(&Cliente, params["id"])
	json.NewDecoder(r.Body).Decode(&Cliente)
	DB.Save(&Cliente)
	json.NewEncoder(w).Encode(Cliente)
}

func DeleteCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var Cliente Cliente
	DB.Delete(&Cliente, params["id"])
	json.NewEncoder(w).Encode("The Cliente is Deleted Successfully!")
}