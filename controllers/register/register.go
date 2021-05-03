package register

import (
	"encoding/json"
	"fmt"
	"go-marketplace-rest-api/helpers"
	"go-marketplace-rest-api/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func HandleNewUser(w http.ResponseWriter, r *http.Request) {
	db, err:= helpers.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	var newUser models.RegisterUser
	err = json.NewDecoder(r.Body).Decode(&newUser)

	password := []byte(newUser.Password)
	hashedPass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPass))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatal(err.Error())
		return
	}

	sqlStatement := `INSERT INTO credential (username, email, password) VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, newUser.Username, newUser.Email, hashedPass)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
