package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	// Create error response object
	errResp := ErrorResponse{
		Status:  statusCode,
		Message: message,
	}

	// Set content type and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Encode error response to JSON and write to ResponseWriter
	err := json.NewEncoder(w).Encode(errResp)
	if err != nil {
		// If encoding fails, log the error and write a simple message
		log.Printf("Failed to encode error response: %v", err)
		w.Write([]byte(`{"status":500,"message":"Internal server error"}`))
	}
}

type GetSuccessResponse struct {
	Status int    `json:"status"`
	Value  string `json:"value"`
}

func WriteGetSuccess(w http.ResponseWriter, statusCode int, value string) {
	successResponse := GetSuccessResponse{
		Status: statusCode,
		Value:  value,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(successResponse)
	if err != nil {
		log.Printf("Failed to encode error response: %v", err)
		w.Write([]byte(`{"status":500,"message":"Internal server error"}`))
	}
}
