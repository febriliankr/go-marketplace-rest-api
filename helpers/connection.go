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

	fmt.Println("Connected! ✅")


	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
db.SetConnMaxLifetime(5*time.Minute)
	return db, nil
}
