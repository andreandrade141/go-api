package main

import (
	"go-api/controllers"
	"go-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDB()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)

	r.Run(":8080")
}
