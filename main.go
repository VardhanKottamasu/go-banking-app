package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",greetHandler)
	fmt.Println("Listening on port 5000")
	err:=http.ListenAndServe(":5000",nil); if err!=nil{
		fmt.Println("Error starting the server")
	}
}


func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Bonjour!!")
}