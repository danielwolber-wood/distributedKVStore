package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	store := NewKVStore()
	r := mux.NewRouter()
	r.HandleFunc("POST /v1/api", store.handlerPost)
	r.HandleFunc("PUT /v1//api", store.handlerPut)
	r.HandleFunc("DELETE /v1/api/{key}", store.handlerDelete)
	r.HandleFunc("GET /v1/api/{key}", store.handlerGet)
	r.HandleFunc("/health", handlerHealthCheck)
	fmt.Println("Listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
