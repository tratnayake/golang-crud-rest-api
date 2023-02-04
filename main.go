package main

import (
	"golang-crud-rest-api/controllers"
	"golang-crud-rest-api/database"

	"github.com/gin-gonic/gin"
)

func main() {
	//Load configs from config.json using Viper
	LoadAppConfig()

	//Initiatlize DB
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	r := gin.Default()

	// Create a Static Route
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ayylmao",
		})
	})
	// Register the actual routes
	r.GET("/demo/associations/menuproducts", controllers.DemoAssociation)
	r.GET("/api/products", controllers.GetProducts)
	r.GET("/api/products/:id", controllers.GetProductById)
	r.POST("/api/products", controllers.CreateProduct)
	r.PUT("/api/products/:id", controllers.UpdateProduct)
	r.DELETE("/api/products/:id", controllers.DeleteProduct)
	// Menu's
	r.GET("/api/menus", controllers.GetMenus)

	r.Run()

	//Try our database thing.

}
