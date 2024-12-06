package routes

import (
	"awesomeProject/database"
	"awesomeProject/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsersHandler(c *gin.Context) {
	var users []models.User
	result := database.DB.Preload("Blogs").Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetOneUserHandler(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	result := database.DB.Preload("Blogs").First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUserHandler(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}
