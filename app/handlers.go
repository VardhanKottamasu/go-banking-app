package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bonjour!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Customer1", City: "City1", Zipcode: "112233"},
		{Name: "Customer2", City: "City2", Zipcode: "112234"},
		{Name: "Customer3", City: "City3", Zipcode: "112235"},
	}
	fmt.Fprint(w, customers)
}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars:=mux.Vars(r)
	fmt.Fprint(w, vars["id"])
}
