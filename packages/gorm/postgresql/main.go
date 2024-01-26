package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=gorm port=5433 sslmode=disable TimeZone=Asia/Tehran"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db = conn
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})
	seedDatabase()

	getCodes := func() ([]string, error) {
		var products []Product
		res := []string{}

		dbres := db.Model(&Product{}).Find(&products)

		log.Println("Record count:", len(products))
		for _, v := range products {
			res = append(res, v.Code)
		}
		return res, dbres.Error
	}

	/*
		if err := setPrice(2, 150); err != nil {
			panic(err)
		}
		log.Println("Price has been updated")
	*/

	codes, err := getCodes()
	if err != nil {
		panic(err)
	}
	log.Println("Codes:", codes)

	updateTagsForId(2)
	initWishlist()
}
