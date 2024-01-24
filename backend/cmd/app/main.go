package main

import (
	"log"
	"sklad/jwt/controllers"
	"sklad/jwt/database"
	"sklad/jwt/middleware"
	"github.com/gin-gonic/gin"
)

func handlerFunc() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome To This Website")
	   })
	   // Create a new group for the API
	   api := r.Group("/api")
	   {
		// Create a new group for the public routes
		public := api.Group("/public")
		{
		 // Add the login route
		 public.POST("/login", controllers.Login)
		 // Add the signup route
		 public.POST("/signup", controllers.Signup)
		}
		// Add the signup route
		protected := api.Group("/protected").Use(middleware.AuthZ())
		{
		 // Add the profile route
		 protected.GET("/profile", controllers.Profile)   
		}
	   }
	return r
}

func main() {
	err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	router := handlerFunc()
	router.Run("localhost:8080")
}
