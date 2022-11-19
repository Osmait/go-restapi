package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnetion() {
	DBS := "host=localhost user=osmait password=123456 dbname=testGO port=5432"
	var err error
	DB, err = gorm.Open(postgres.Open(DBS), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DataBase Connetion successful")
}
