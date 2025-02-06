package main

import (
	"gorm.io/gorm"
)

func migrateSchemas(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(
		&Product{},
		&PlainUser{},
		&User{},
	)
}
