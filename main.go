package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lelandaure/testing-in-go/controller"
	"net/http"
)

var port = ":8080"

func Add(a, b int) int {
	return a + b
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/pokemon/{id}", controller.GetPokemon).Methods("GET")

	err := http.ListenAndServe(port, router)
	if err != nil {
		fmt.Print("Error found")
	}
	fmt.Println("Listening on port", port)
}
