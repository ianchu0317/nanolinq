package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Server struct and creation

type shortenServer struct {
	mu sync.Mutex
}

func CreateServer() Server {
	return &shortenServer{}
}

// Server Handlers

func (s *shortenServer) CreateURL(w http.ResponseWriter, r *http.Request) {
	// Lock connection resources
	s.mu.Lock()
	defer s.mu.Unlock()

	if r.Method != "POST" {
		http.Error(w, "Only POST Method allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read body content
	defer r.Body.Close()
	var bodyData CreateURLData
	err := json.NewDecoder(r.Body).Decode(&bodyData)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Get URL from Body and get shorten url
	url := bodyData.Url
	shortenURL := createShortenURL(url)

	//fmt.Println("url is: ", shortenURL)
	fmt.Fprintf(w, "url is: %s\n", shortenURL)
	// Store new content on db

	// Return response to user
}
