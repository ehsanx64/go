package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ProductId uint
	Product   Product
	Content   string
}

func initReview() {
	log.Println("initReview()")
	db.AutoMigrate(&Review{})
	var r []Review

	res := db.Preload("Product").Find(&r)
	if res.Error != nil {
		panic(res.Error)
	}

	if res.RowsAffected < 1 {
		log.Println("No reviews found. Inserting reviews ...")

		r = []Review{
			{
				ProductId: 2,
				Content:   "Good product",
			},
			{
				ProductId: 3,
				Content:   "Not so good",
			},
		}

		db.Create(&r)
	}

	log.Println("Review records count:", res.RowsAffected)

	var meta map[string]interface{}
	var tags []string
	type ResultMeta struct {
		Name string `json:"name"`
		Age  int
	}
	type ResultTags []string
	type Result struct {
		ProductId uint
		Meta      ResultMeta
		Tags      ResultTags
	}
	var result Result
	for _, r := range r {
		if err := json.Unmarshal(r.Product.Tags.Bytes, &tags); err != nil {
			panic(err)
		}

		if err := json.Unmarshal(r.Product.Meta.Bytes, &meta); err != nil {
			panic(err)
		}

		fmt.Printf("%v:%T\n", tags, tags)
		fmt.Println("First tag:", tags[0])

		fmt.Printf("%v:%T\n", meta, meta)
		fmt.Println("Name:", meta["name"])

		result.ProductId = r.Product.ID
		result.Meta = ResultMeta{
			Name: meta["name"].(string),
			Age:  int(meta["age"].(float64)),
		}
		result.Tags = tags

		fmt.Println("result:")
		fmt.Println("Name:", result.Meta.Name)
	}
}
