package seller

import (
	"encoding/json"
	"fmt"
	"go-marketplace-rest-api/helpers"
	"go-marketplace-rest-api/models"
	"log"
	"net/http"
)

func GetSeller(w http.ResponseWriter, r *http.Request) {
	db, err := helpers.InitDB()
	if err != nil {
		message := fmt.Errorf("something Happened with the database, %s", err.Error())
		fmt.Println(message)
	}

	rows, err := db.Queryx("SELECT * FROM seller")
	if err != nil {
		panic(err)
	}

	var sellers []models.Seller

	for rows.Next() {
		var seller models.Seller
		err := rows.StructScan(&seller)
		if err != nil {
			log.Fatalln(err)
		}
		sellers = append(sellers, seller)
	}

	responseBytes, _ := json.MarshalIndent(sellers, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBytes)
}

func InsertSeller(w http.ResponseWriter, r *http.Request) {
	db, err := helpers.InitDB()

	if err != nil {
		panic(err.Error())

	}

	var newSeller models.Seller
	err = json.NewDecoder(r.Body).Decode(&newSeller)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatal(err.Error())
		return
	}

	sqlStatement := `INSERT INTO seller (username, display_name, email, display_picture_url, address) VALUES ($1, $2, $3, $4, $5)`
	_, err = db.Queryx(sqlStatement, newSeller.Username, newSeller.Email, newSeller.DisplayPictureUrl, newSeller.DisplayName, newSeller.Address)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
