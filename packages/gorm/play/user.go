package main

import (
	"database/sql"
	"log"
	"time"

	"gorm.io/gorm"
)

type PlainUser struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
}

type User struct {
	gorm.Model

	Name     string
	Email    *string
	Age      uint8
	Birthday *time.Time

	CompanyID uint
	Company   Company

	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}

type Company struct {
	ID   uint
	Name string
}

func userTests(db *gorm.DB) {
	insertUsers := func() {
		// Insert user Adam
		adam := User{
			Name: "Adam",
			Age:  40,
			Company: Company{
				Name: "Google",
			},
		}

		eve := User{
			Name: "Eve",
			Age:  40,
			Company: Company{
				Name: "Microsoft",
			},
		}

		db.Create(&adam)
		db.Create(&eve)
	}

	if false {
		insertUsers()
	}

	// Get all users with name "Adam"
	var adam []User
	res := db.Where("name = ?", "Adam").Find(&adam)
	log.Println("res:", res)
	log.Println(res.RowsAffected)
	log.Println(adam)
}
