package main

import (
	"log"

	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type FormTags []interface{}
type FormMeta map[string]interface{}

type Product struct {
	gorm.Model
	Code  string
	Price uint
	Tags  pgtype.JSONB `gorm:"type:jsonb;default:'[]';not null" json:"tags"`
	Meta  pgtype.JSONB `gorm:"type:jsonb;default:'{}';not null" json:"meta"`
}

func updateTagsForId(id uint) {
	var product Product
	result := db.Model(&product).Where("id = ?", id).Find(&product)
	if result.Error != nil {
		panic(result.Error)
	}
	product.Tags.Set([]string{"Tag X", "Tag Y"})
	db.Model(&product).Updates(&product)
	log.Println("Product")
	log.Println(product)
}

func seedDatabase() {
	var p = Product{
		Code:  "E77",
		Price: 1000,
	}

	p.Tags.Set(FormTags{"Tag3", "Tag4"})
	p.Meta.Set(FormMeta{
		"name": "Adam Smith",
		"age":  30,
	})
	// Create
	db.Create(&p)

	var prod Product
	db.Model(&Product{}).Find(&prod, 1)
	log.Println(prod)
	return

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}

func setPrice(id uint, price uint) error {
	var product Product

	db.First(&product, id)
	result := db.Model(&product).Updates(Product{Price: price})
	return result.Error
}
