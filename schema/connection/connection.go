package connection

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() (db *gorm.DB, err error) {
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	timeZone := os.Getenv("DB_TIMEZONE")

  dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", 
		host, user, password, dbName, port, timeZone,
	)
  
  return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
