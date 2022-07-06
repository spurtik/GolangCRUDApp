package controllers

import (
	"net/http"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

// GET /users
// Find all users
func FindUsers(ctx *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

// POST /user
// Create user
func CreateUser(ctx *gin.Context) {
	var input models.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Email: input.Email}
	models.DB.Create(&user)
	ctx.JSON(http.StatusOK, gin.H{"users": user})
}

// GET /user/:id
// Get user by id
func FindUserById(ctx *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

// PATCH /user/:id
// Update a user data
func UpdateUserById(c *gin.Context) {
	// check if user exists
	var user models.User
	if err := models.DB.Where("id=?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Vaidate input
	var input models.UpdateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Update user data
	updateUserData := models.User{Name: input.Name, Email: input.Email}
	models.DB.Model(&user).Updates(updateUserData)

	c.JSON(http.StatusOK, gin.H{"user": user})

}

// Delete user by Id
func DeleteUserById(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id=?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found!"})
		return
	}
	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"user": true})
}
