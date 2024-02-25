package main

import (
	"restAPI/database/api"
	"restAPI/database/schema/connection"
	"restAPI/database/schema/migrations"

	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
 	dbErr error
)

func handleDBError(dbErr error) bool {
	if dbErr != nil {
		date := strings.Split(time.Now().Format(time.RFC3339), "T")
		f, err := os.OpenFile(
			date[0] + " error.log", 
			os.O_RDWR | os.O_CREATE | os.O_APPEND, 
			0766,
		)

		f.Write([]byte(dbErr.Error()))
		
		if (err != nil)  {
			log.Fatal("Could not write due to permissions")
		}

		defer f.Close()
		return true
	}

	return false
}

func Middleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context)  {
		c.Set("db", db)
		c.Next()
	}
}

func LoadENV() {
  if err := godotenv.Load(); err != nil {
    log.Fatal("Error loading .env file")
  }
}

func main() {
	LoadENV()

	db, dbErr = connection.CreateConnection()

	if handleDBError(dbErr) {
		os.Exit(0)
	}

	migrations.MigrateModels(db); 

	router := gin.New()	

	router.Use(gin.Recovery())
	router.Use(Middleware(db))

	userGroup := router.Group("/user")
	api.SetupUserRoutes(userGroup)

	router.Run("localhost:8081")
}