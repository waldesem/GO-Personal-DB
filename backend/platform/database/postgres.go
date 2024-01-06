package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDb() *gorm.DB {

	host := os.Getenv("DBHOST")
	usr := os.Getenv("DBUSER")
	pwd := os.Getenv("DBPASSWORD")
	name := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")

	var dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, usr, pwd, name, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
