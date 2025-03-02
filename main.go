package main

import (
	"jwt-authentication-go/controllers"
	"jwt-authentication-go/initializers"
	"jwt-authentication-go/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run(":" + port)
}
