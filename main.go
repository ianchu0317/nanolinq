package main

import (
	"log"
	"net/http"
	"url-shortening/server"
)

func main() {
	app := server.CreateServer()
	http.HandleFunc("/shorten", app.CreateURL)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
