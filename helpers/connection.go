package helpers

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var (
	host     = "ec2-23-23-128-222.compute-1.amazonaws.com"
	port     = 5432
	dbname   = "ddjpiu188bbrp1"
	user     = "pawihmetgbdaiw"
	password = "2652632a0823686f891bdf84bbb5c22b2c6cdcbffcd69fc4e8e0cdf5a9f72af1"
)




func InitDB() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected! ✅")


	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
db.SetConnMaxLifetime(5*time.Minute)
	return db, nil
}
