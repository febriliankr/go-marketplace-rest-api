package main

import (
	"go-marketplace-rest-api/helpers"
	"go-marketplace-rest-api/routes"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	routes.InitializeRoutes()



	// Get the PORT from heroku env
	port := os.Getenv("PORT")

	// Verify if heroku provided the port or not
	if os.Getenv("PORT") == "" {
		port = "8080"
	}

	helpers.InitDB()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
