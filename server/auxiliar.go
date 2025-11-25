package server

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
)

const _SHORTEN_LEN = 10

// createShortenURL takes a string and returns a shorten version of it (hash)
func createShortCode(url string) string {
	h := sha256.New()
	h.Write([]byte(url))
	hashBytes := h.Sum(nil)

	// Encode the hash to a hexadecimal string
	hexHash := hex.EncodeToString(hashBytes)

	// Return a truncated portion of the hex hash
	if len(hexHash) > _SHORTEN_LEN {
		return hexHash[:_SHORTEN_LEN]
	}
	return hexHash
}

// Helpers - Auxiliar functions

// ReturnJSON take a request and respond with JSON
func ReturnJSON(w http.ResponseWriter, r *http.Request, responseData *ResponseCreatedURLData, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		log.Fatalf("Internal server error converting to JSON, %v", err)
		http.Error(w, "Internal server error converting to JSON", http.StatusInternalServerError)
		return
	}
}

// ReturnError takes an error, message and status code and returns error to user / and log
func ReturnError(w http.ResponseWriter, err error, message string, statusCode int) {
	log.Printf("%s: %v", message, err)
	http.Error(w, message, statusCode)
}
