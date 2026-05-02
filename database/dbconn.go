package database

import (
	"fmt"
	"log"
	"os"

	"jwt-auth-api/models" // this will be imported after you've created the User Model in the models.go file

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconn() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbPort, dbUser, dbPass, dbName)

	// gorm.Open() creates a new connection pool whenever it is called
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	// db.AutoMigrate() call helps in creating the table if it is not already present
	db.AutoMigrate(&models.User{}) // we are going to create a models.go file for the User Model.
}
