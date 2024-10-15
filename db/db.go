package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB connection instance
var DB *gorm.DB

// ConnectDB initializes a DB connection.
func ConnectDB() error {
	dsn := os.Getenv("POSTGRES_DSN")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}

// SaveUser saves a user record to the DB (example function).
func SaveUser(name string, age int) error {
	type User struct {
		Name string
		Age  int
	}
	return DB.Create(&User{Name: name, Age: age}).Error
}
