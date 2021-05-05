package routes

import (
	"go-marketplace-rest-api/controllers/login"
	"go-marketplace-rest-api/controllers/register"
	"go-marketplace-rest-api/controllers/seller"
	"net/http"
)

func InitializeRoutes(){:
	http.HandleFunc("/seller", seller.GetSeller)
	http.HandleFunc("/seller/insert", seller.InsertSeller)
	http.HandleFunc("/register", register.HandleNewUser)
	http.HandleFunc("/login", login.HandleLoginWithEmailAndPassword)
	http.Handle("/client", http.FileServer(http.Dir("./static")))
}