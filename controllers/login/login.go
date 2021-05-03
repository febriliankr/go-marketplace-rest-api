package login

import (
	"encoding/json"
	"fmt"
	"go-marketplace-rest-api/controllers/password"
	"go-marketplace-rest-api/helpers"
	"go-marketplace-rest-api/models"
	"log"
	"net/http"
	"strconv"
)



func HandleLoginWithEmailAndPassword(w http.ResponseWriter, r *http.Request) {

	db, err := helpers.InitDB()
	if err != nil {
		panic(err.Error())
	}

	var loggingInUser models.UsernameAndPassword
	err = json.NewDecoder(r.Body).Decode(&loggingInUser)

	sqlStatement := `SELECT id, username, password FROM credential WHERE (username = $1)`

	rows, err := db.Queryx(sqlStatement, loggingInUser.Username)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	var credentials []models.IdUsernameAndPassword

	for rows.Next() {

		var credential models.IdUsernameAndPassword
		err := rows.StructScan(&credential)
		if err != nil {
			log.Fatalln(err)
		}

		matchingPassword := password.ComparePasswords(credential.Password, []byte(loggingInUser.Password))

		if matchingPassword {

			intId, err := strconv.Atoi(credential.Id)

			if err != nil {
				fmt.Errorf(err.Error())
			}

			responseData := models.LoginResponse{
				Status:   "OK",
				UserExists: true,
				Id: intId,
				Username: credential.Username,
			}

			responseBytes, _ := json.MarshalIndent(responseData, "", "\t")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(responseBytes)
			return
		} else {
			responseData := models.LoginErrorResponse{
				Status:   "ERROR",
				UserExists: true,
				Message: "Wrong Password",
			}
			responseBytes, _ := json.MarshalIndent(responseData, "", "\t")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(responseBytes)
			return
		}

		credentials = append(credentials, credential)
	}

}
