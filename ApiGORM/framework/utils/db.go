package utils

import (
	"log"
	"os"

	"ApiGORM/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Driver postgres
)

// ConnectDB function return connection with database
func ConnectDB() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbtype := os.Getenv("dbtype")
	dsn := os.Getenv("dsn")

	db, err := gorm.Open(dbtype, dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	//defer db.Close()

	db.AutoMigrate(&domain.User{})

	return db

}
