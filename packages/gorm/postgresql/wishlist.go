package main

import (
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

type Wishlist struct {
	gorm.Model
	ProductId uint
	Product   Product
}

func initWishlist() {
	db.AutoMigrate(&Wishlist{})

	w := Wishlist{
		ProductId: 5,
	}

	db.Create(&w)

	var wl Wishlist
	var tagsJson []interface{}
	var metaJson map[string]interface{}
	const productId = 55

	// Select first record with given product id
	db.Preload("Product").Where("product_id = ?", productId).First(&wl)

	// Print the wishlist id
	log.Println("Wishlist id:", wl.ID)

	// Convert and print product tags field
	log.Println("Product Tags:", string(wl.Product.Tags.Bytes))

	// Unmarshal product tags
	json.Unmarshal(wl.Product.Tags.Bytes, &tagsJson)
	log.Println("Tags json:", tagsJson)
	log.Println("First tag:", tagsJson[0])

	// Unmarshal product meta
	json.Unmarshal(wl.Product.Meta.Bytes, &metaJson)
	log.Println("Meta json:", metaJson)
	log.Println("Meta name:", metaJson["name"])
	log.Println("Meta age:", metaJson["age"])

	log.Printf("Tags dump: %+v\n", wl.Product.Tags)
	tags := wl.Product.Tags.Get()
	log.Printf("%+T\n", tags)
	log.Println("Tags Get():", tags)
}
