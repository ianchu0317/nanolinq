package server

import "time"

// Requests body format shema

type CreateURLData struct {
	Url string `json:"url"`
}

// Response body format schemas

type ResponseCreatedURLData struct {
	ID        int       `json:"id"`
	URL       string    `json:"url"`
	ShortCode string    `json:"shortCode"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Accessed  int       `json:"accessed"`
}
