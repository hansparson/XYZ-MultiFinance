package db

import (
	"fmt"
	"log"
	"os"
	dbschema "xyz-multifinance/db/db-schema"

	"github.com/joho/godotenv" // Import the godotenv package
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectGorm() *gorm.DB {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURI := os.Getenv("SQLALCHEMY_DATABASE_URI")

	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Migrate Product Database
	errors := db.AutoMigrate(dbschema.User{}, dbschema.Transaction{}, dbschema.Bill{}, dbschema.MonthlyBilling{}, dbschema.UserLimitBalance{})
	if errors != nil {
		log.Println(errors.Error())
	}

	return db
}

func generateUserID(db *gorm.DB) (string, error) {
	var maxID int
	err := db.Raw("SELECT MAX(CAST(user_id AS UNSIGNED)) FROM users").Scan(&maxID).Error
	if err != nil {
		return "", err
	}
	newID := fmt.Sprintf("%08d", maxID+1) // Membuat string dengan format 8 digit angka
	return newID, nil
}
