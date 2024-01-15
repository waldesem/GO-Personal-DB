package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDb() *gorm.DB {

	host, _ := os.LookupEnv("DBHOST")
	usr, _ := os.LookupEnv("DBUSER")
	pwd, _ := os.LookupEnv("DBPASSWORD")
	name, _ := os.LookupEnv("DBNAME")
	port, _ := os.LookupEnv("DBPORT")

	var dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, usr, pwd, name, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	return db
}
