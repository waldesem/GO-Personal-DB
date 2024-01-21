package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDb() *gorm.DB {
	dbhost, _ := os.LookupEnv("DBHOST")
	dbusr, _ := os.LookupEnv("DBUSER")
	dbpwd, _ := os.LookupEnv("DBPASSWORD")
	dbname, _ := os.LookupEnv("DBNAME")
	dbport, _ := os.LookupEnv("DBPORT")

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
