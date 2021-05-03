package main

import (
	"go-marketplace-rest-api/controllers/login"
	"go-marketplace-rest-api/controllers/register"
	"go-marketplace-rest-api/controllers/seller"
	"go-marketplace-rest-api/helpers"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)



func main() {

	http.HandleFunc("/", seller.GetSeller)
	http.HandleFunc("/seller/insert", seller.InsertSeller)
	http.HandleFunc("/register", register.HandleNewUser)
	http.HandleFunc("/login", login.HandleLoginWithEmailAndPassword)
	http.Handle("/client", http.FileServer(http.Dir("./static")))

	// Get the PORT from heroku env
	port := os.Getenv("PORT")

	// Verify if heroku provided the port or not
	if os.Getenv("PORT") == "" {
		port = "8080"
	}

	helpers.InitDB()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
