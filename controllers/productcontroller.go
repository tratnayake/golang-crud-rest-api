package controllers

import (
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Helper
func checkIfProductExists(productId string) bool {
	var product models.Product
	database.DB.First(&product, productId)
	if product.ID == 0 {
		return false
	}
	return true
}

// Demo's
func DemoAssociation(c *gin.Context) {
	menu := models.Menu{
		Products: []*models.Product{
			{
				Name:        "Test1",
				Price:       15,
				Description: "Test",
			},
		},
	}

	database.DB.Create(&menu)
	database.DB.Save(&menu)
	c.JSON(http.StatusOK, menu)
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context) {
	productId := c.Param("id")
	if checkIfProductExists(productId) == false {
		// Product not found, send error response.
		c.JSON(http.StatusNotFound, "Product Not Found!")
		return
	}
	// NTS: Create a new empty struct of type models.Product
	var product models.Product
	// NTS: Store it in that empty struct.
	database.DB.First(&product, productId)
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	// NTS: This is an interesting error definition / handling 1-liner.
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusNotAcceptable, "Unable to bind body")
		return
	}

	database.DB.Create(&product)
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	productId := c.Param("id")
	if checkIfProductExists(productId) == false {
		// Product not found, send error response.
		c.JSON(http.StatusNotFound, "Product Not Found!")
		return
	}
	var product models.Product
	var updatedProduct models.Product
	database.DB.First(&product, productId) // Nothing is done with this.
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusNotAcceptable, "Unable to bind the updated product.")
	}
	database.DB.Save(&updatedProduct)
	c.JSON(http.StatusOK, updatedProduct)
}

func DeleteProduct(c *gin.Context) {
	productId := c.Param("id")

	if checkIfProductExists(productId) == false {
		// Product not found, send error response.
		c.JSON(http.StatusNotFound, "Product Not Found!")
		return
	}

	var product models.Product
	database.DB.Delete(&product, productId)
	c.JSON(http.StatusOK, gin.H{"message": "Product Deleted Successfully"})

}
