package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux

	mux.HandleFunc("/search", GetInfo)
	mux.HandleFunc("/event/{id}", EventInfo)
	mux.HandleFunc("/event/{id}/book/{seat}", BookTheSeat)

	fmt.Printf("Server has started")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
