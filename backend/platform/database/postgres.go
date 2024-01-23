package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDb() *gorm.DB {
	dbhost := os.Getenv("DBHOST")
	dbusr := os.Getenv("DBUSER")
	dbpwd := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")
	dbport := os.Getenv("DBPORT")

	var dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbhost, dbusr, dbpwd, dbname, dbport,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	return db
}
