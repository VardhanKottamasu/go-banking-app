package app

import (
	// "encoding/json"
	"fmt"
	"net/http"
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
	// dataJson, err := json.Marshal(customers)
	// if err != nil {
	// 	fmt.Println("Error marshalling data")
	// }
	fmt.Fprint(w, customers)
	// w.Header().Add("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(customers)
}
