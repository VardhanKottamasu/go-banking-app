package app

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func Start() {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/greet", greetHandler).Methods("GET")
	http.HandleFunc("/getAllCustomers", getAllCustomers)
	fmt.Println("Listening on port 5000")
	err := http.ListenAndServe(":5000", muxRouter)
	if err != nil {
		fmt.Println("Error starting the server")
	}
}
