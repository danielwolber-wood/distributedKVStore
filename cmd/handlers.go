package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (s *KVStore) handlerPost(w http.ResponseWriter, r *http.Request) {

}

func (s *KVStore) handlerGet(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	value, err := s.Get(key)
	if err != nil {
		fmt.Printf("Error getting value: %v\n", err)
		WriteError(w, http.StatusBadRequest, "value not in store")
	}
	WriteGetSuccess(w, http.StatusOK, value)
}

func (s *KVStore) handlerPut(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling put request\n")
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		fmt.Printf("request body must be json\n")
		WriteError(w, http.StatusBadRequest, "request body must be json")
		return
	}
	fmt.Printf("Passed contenttype\n")
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Printf("error parsing r.body: %v\n", err)
		WriteError(w, http.StatusBadRequest, "error parsing body")
		return
	}
	fmt.Printf("Passed Read Body\n")
	var req bodyPut
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Printf("error while unmarshalling: %v\n", err)
		WriteError(w, http.StatusBadRequest, "error unmarshalling json")
		return
	}
	fmt.Printf("Passed Unmarshal\n")
	fmt.Printf("key: %v\n", req.Key)
	fmt.Printf("value: %v\n", req.Value)
	err = s.Set(req.Key, req.Value)
	if err != nil {
		fmt.Printf("error while setting key: %v\n", err)
		WriteError(w, http.StatusInternalServerError, "error setting key/value pair")
		return
	}
	fmt.Printf("Passed set\n")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successfully added"))
}

func (s *KVStore) handlerDelete(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	_ = s.Delete(key)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("sucessfully deleted"))
}

func handlerHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("server is alive"))
}
