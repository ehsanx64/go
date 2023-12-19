package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbUser    = "test"
	dbPass    = ""
	dbName    = "test"
	dbHost    = "127.0.0.1:3306"
	dbOptions = "charset=utf8mb4&parseTime=True&loc=Local"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func GetDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		dbUser, dbPass, dbHost, dbName, dbOptions)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("wrong database url")
	} else {
		log.Println("Database connection established")
	}

	sqldb, _ := conn.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("database connected")
	}

	fmt.Println("connected to database")
	return conn
}

func InitialMigration() {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.AutoMigrate(User{})
}

func CloseDatabase(connection *gorm.DB) {
	sqldb, _ := connection.DB()
	sqldb.Close()
}
