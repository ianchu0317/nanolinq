package server

import "net/http"

type Server interface {
	// CloseServer closes database connection
	CloseServer()

	// CreateURL handles posts requests from POST /shorten
	CreateURL(w http.ResponseWriter, r *http.Request)
}
