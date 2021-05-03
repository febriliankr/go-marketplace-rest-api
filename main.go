package main

import (
	"go-marketplace-rest-api/controllers"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	http.HandleFunc("/", controllers.GetSeller)
	http.HandleFunc("/insert", controllers.InsertSeller)
	http.Handle("/client", http.FileServer(http.Dir("./static")))

	// Get the PORT from heroku env
	port := os.Getenv("PORT")

	// Verify if heroku provided the port or not
	if os.Getenv("PORT") == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
