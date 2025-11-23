package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"url-shortening/server"
)

func main() {
	fmt.Println("url es: ", os.Getenv("DATABASE_URL"))
	app := server.CreateServer(os.Getenv("DATABASE_URL"))
	http.HandleFunc("/shorten", app.CreateURL)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
