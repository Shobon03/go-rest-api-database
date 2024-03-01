package main

import (
	"restAPI/database/api"
	"restAPI/database/helpers"

	"os"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
 	dbErr error
)

func main() {
	helpers.LoadEnvironmentVariables()

	db, dbErr = helpers.CreateConnection()

	if helpers.HandleDBError(dbErr) {
		os.Exit(0)
	}

	helpers.MigrateModels(db); 

	router := gin.New()	

	router.Use(gin.Recovery())
	router.Use(helpers.Middleware(db))

	userGroup := router.Group("/user")
	api.SetupUserRoutes(userGroup)

	router.Run("localhost:8081")
}