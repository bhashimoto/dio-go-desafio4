package main

import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)


type Customer struct {
	ID string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"` 
	Lastname string `json:"lastname,omitempty"`
	Email string `json:"email,omitempty"`
	Address *Address
}

type Address struct {
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var customers []Customer

func GetCustomers(wr http.ResponseWriter, req *http.Request) {
	log.Println("GetCustomers called")
	json.NewEncoder(wr).Encode(customers)
}

func GetCustomer(wr http.ResponseWriter, req *http.Request) {
	log.Println("GetCustomer called")
	params := mux.Vars(req)

	for _, item := range customers {
		if item.ID == params["id"] {
			json.NewEncoder(wr).Encode(item)
			return
		}
	}
	json.NewEncoder(wr).Encode(&Customer{})
}

func CreateCustomer(wr http.ResponseWriter, req *http.Request) {
	log.Println("CreateCustomer called")
	params := mux.Vars(req)

	var customer Customer

	_ = json.NewDecoder(req.Body).Decode(&customer)

	customer.ID = params["id"]
	customers = append(customers, customer)

	json.NewEncoder(wr).Encode(customers)
}

func DeleteCustomer(wr http.ResponseWriter, req *http.Request) {
	log.Println("DeleteCustomer called")
	params := mux.Vars(req)

	for index, item := range customers {
		if item.ID == params["id"] {
			customers = append(customers[:index], customers[index+1:]...)
			break
		}
	}
	json.NewEncoder(wr).Encode(customers)
}

func HomeHandler(wr http.ResponseWriter, req *http.Request) {
	json.NewEncoder(wr).Encode(customers)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/",HomeHandler).Methods("GET")
	r.HandleFunc("/customers", GetCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", GetCustomer).Methods("GET")
	r.HandleFunc("/customers/{id}", CreateCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", DeleteCustomer).Methods("DELETE")

	http.ListenAndServe(":3000",r)
}
