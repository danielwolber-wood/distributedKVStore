package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	store := NewKVStore()
	http.HandleFunc("POST /api", store.handlerPost)
	http.HandleFunc("PUT /api", store.handlerPut)
	http.HandleFunc("DELETE /api/{key}", store.handlerDelete)
	http.HandleFunc("GET /api/{key}", store.handlerGet)
	http.HandleFunc("/health", handlerHealthCheck)
	fmt.Println("Listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
