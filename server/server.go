package server

import "net/http"

type Server interface {
	// CreateURL handles posts requests from POST /shorten
	CreateURL(w http.ResponseWriter, r *http.Request)
}
