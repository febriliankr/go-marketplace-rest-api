package helpers

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"log"
	"time"
)



var (
	host     = "localhost"
	port     = 5432
	dbname   = "marketplace"
	user     = "postgres"
	password = "timerunner"
)

func InitDB() (*sqlx.DB, error) {
	// Connect SQL DB
	db, err := sqlx.Connect("postgres", "user=postgres dbname=marketplace sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
db.SetConnMaxLifetime(5*time.Minute)
	return db, nil
}

func OpenConnection() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected! ✅")
	return db
}
