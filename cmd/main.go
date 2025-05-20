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
	r.HandleFunc("POST /api", store.handlerPost)
	r.HandleFunc("PUT /api", store.handlerPut)
	r.HandleFunc("DELETE /api/{key}", store.handlerDelete)
	r.HandleFunc("GET /api/{key}", store.handlerGet)
	r.HandleFunc("/health", handlerHealthCheck)
	fmt.Println("Listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
