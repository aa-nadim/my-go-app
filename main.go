package main

import (
	"log"
	"net/http"
	"os"

	"my-go-app/routers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Load routes from routers package
	routers.SetupRoutes()

	log.Printf("Server running at http://localhost:%s/", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
