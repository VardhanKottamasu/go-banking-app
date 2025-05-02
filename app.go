package main

import (
	"fmt"
	"net/http"
)

func Start() {
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/getAllCustomers", getAllCustomers)
	fmt.Println("Listening on port 5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("Error starting the server")
	}
}
