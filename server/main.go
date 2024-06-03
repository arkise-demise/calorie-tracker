package main

import (
	"calorie-tracker/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id", routes.GetEntryById) 
	router.GET("/ingredient/:ingredient", routes.GetEntriesByIngredient)

	router.PUT("/entry/:id/update", routes.UpdateEntry)            
	router.PUT("/ingredient/:id/update", routes.UpdateIngredient)   
	router.DELETE("/entry/:id/delete", routes.DeleteEntry)         

	router.Run(":" + port)
}
