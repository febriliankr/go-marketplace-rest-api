package login

import (
	"encoding/json"
	"fmt"
	"go-marketplace-rest-api/helpers"
	"go-marketplace-rest-api/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func HandleLoginWithEmailAndPassword(w http.ResponseWriter, r *http.Request) {

	db, err := helpers.InitDB()
	if err != nil {
		panic(err.Error())
	}

	var loggingInUser models.UsernameAndPassword
	err = json.NewDecoder(r.Body).Decode(&loggingInUser)

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(loggingInUser.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatal(err.Error())
		return
	}
	sqlStatement := `SELECT username FROM credential WHERE (username = $1) and (password = $2)`

	rows, err := db.Queryx(sqlStatement, loggingInUser.Username, hashedPass)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	var usernames []models.Username

	for rows.Next() {

		var username models.Username
		err := rows.StructScan(&username)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("#{username}\n")

		usernames = append(usernames, username)
	}

	responseBytes, _ := json.MarshalIndent(rows, "", "\t")

	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)

}
