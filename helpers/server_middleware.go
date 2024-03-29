package helpers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Middleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context)  {
		c.Set("db", db)
		c.Next()
	}
}
