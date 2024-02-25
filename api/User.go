package api

import (
	"restAPI/database/helpers"
	"restAPI/database/schema/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User models.User

type RequiredUserFields struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	user  User
	users []User
	userFields RequiredUserFields
)

func SetupUserRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", GetUsers)
	routerGroup.POST("/", CreateUser)
	
	routerGroup.GET("/:id", GetUser)
	routerGroup.PUT("/:id", UpdateUser)
	routerGroup.DELETE("/:id", DeleteUser)
}

func GetUsers(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)

	if (!ok) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "could not connect to database"})
	}

	if result := db.Find(&users); result.Error == nil {
		c.IndentedJSON(http.StatusFound, users)
	} else {
		c.IndentedJSON(http.StatusFound, gin.H{"message": result.Error})
	}
}

func GetUser(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)

	if (!ok) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "could not connect to database"})
	}

	userId := c.Param("id")

	if result := db.Where("id = ?", userId).First(&user); result.Error == nil {
		c.IndentedJSON(http.StatusFound, user)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
}

func CreateUser(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)

	if (!ok) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "could not connect to database"})
	}

	if err := c.BindJSON(&userFields); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "malformatted json"})
		return
	}

	// Field validation
	if userFields.Name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "NAME REQUIRED"})
		return
	}
	
	if userFields.Email == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "EMAIL REQUIRED"})
		return
	}

	if userFields.Password == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "PASSWORD REQUIRED"})
		return
	}

	user = User{
		Name: userFields.Name,
		Email: userFields.Email, 
		Password: helpers.HashPassword(userFields.Password),
	}

	if result := db.Create(&user); result.Error == nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"user": user})
	} else {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": result.Error})
	}
}

func UpdateUser(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)

	if (!ok) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "could not connect to database"})
	}

	userId := c.Param("id")

	if err := c.BindJSON(&userFields); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "malformatted json"})
		return
	}

	if result := db.Where("id = ?", userId).First(&user); result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if userFields.Name != "" {
		user.Name = userFields.Name;
	}
	
	if userFields.Email != "" {
		user.Email = userFields.Email;
	}

	if userFields.Password != "" {
		user.Password = helpers.HashPassword(userFields.Password);
	}

	db.Save(&user)

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "user updated successfully"})
}

func DeleteUser(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)

	if (!ok) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "could not connect to database"})
	}

	userId := c.Param("id")

	if result := db.Where("id = ?", userId).First(&user); result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	db.Where("id = ?", userId).Delete(&user)

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "user deleted successfully"})
}