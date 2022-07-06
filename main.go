package main

import (
	"example.com/m/controllers"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func main() {
	req := gin.Default()
	models.InitDb()

	// req.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	// })
	req.GET("/users", controllers.FindUsers)
	req.POST("/user", controllers.CreateUser)
	req.GET("/user/:id", controllers.FindUserById)
	req.PATCH("/user/:id", controllers.UpdateUserById)
	req.DELETE("/user/:id", controllers.DeleteUserById)

	req.Run(":8080")
}
