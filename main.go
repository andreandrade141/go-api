package main

import (
	"go-api/controllers"
	"go-api/middlewares"
	"go-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDB()

	r := gin.Default()

	public := r.Group("/api")
	protected := r.Group("/api/admin")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")
}
