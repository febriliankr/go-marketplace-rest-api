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

)
//
//CREATE TABLE public.seller (
//id integer NOT NULL,
//username character varying DEFAULT ''::character varying NOT NULL,
//email character varying DEFAULT ''::character varying,
//display_picture_url character varying DEFAULT ''::character varying,
//address character varying DEFAULT ''::character varying,
//display_name character varying
//);



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
