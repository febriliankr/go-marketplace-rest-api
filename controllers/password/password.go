package password

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func ComparePasswords(passwordFromDB string, inputtedPassword []byte) bool {
	bytePasswordFromDB := []byte(passwordFromDB)
	err := bcrypt.CompareHashAndPassword(bytePasswordFromDB, inputtedPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}