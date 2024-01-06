package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projects/luhn_Algorithm/LuhmValidator"
)

type ValidationRequest struct {
	CardNumber string `json:"card_number"`
}

type ValidationResult struct {
	CardNumber string `json:"card_number"`
	IsValid    bool   `json:"is_valid"`
}

func main() {
	// url = "http://localhost:5000"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusBadRequest)
			return
		}

		// Store request data in ValidationRequest struct instance
		var request ValidationRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Validate the request card number with LuhnValidator paclage's IsValid function
		isValid := LuhmValidator.IsValid(request.CardNumber)
		result := ValidationResult{
			CardNumber: request.CardNumber,
			IsValid:    isValid,
		}
		// Marshal the struct into JSON
		jsonResult, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set the content type header to JSON
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response
		w.Write(jsonResult)
	})

	fmt.Println("Server listening on port 5000...")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
