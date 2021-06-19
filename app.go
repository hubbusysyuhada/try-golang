package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("masuk main")
	router := mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Home(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(map[string]interface{}{
		"status":  200,
		"message": "hello world",
	})
	fmt.Println("ini home")
}
