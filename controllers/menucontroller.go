package controllers

import (
	"fmt"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func checkIfMenuExists(menuId string) bool {
	var menu models.Menu
	database.DB.First(&menu, menuId)
	if menu.ID == 0 {
		return false
	}
	return true
}

func GetMenus(c *gin.Context) {
	var menus []models.Menu
	// database.DB.Model(&models.Menu{}).Preload("Products").Find(&menus).Omit("Products.name")
	database.DB.Model(&models.Menu{}).Preload("Products", func(tx *gorm.DB) *gorm.DB {
		return tx.Omit("menus")
	}).Find(&menus)
	c.JSON(http.StatusOK, menus)
}

type MenuRequestBody struct {
	Name  string `json:"name" binding:"required"`
	Items []struct {
		Id       string
		Quantity int
	} `json:"items" binding:"required"`
}

func CreateMenu(c *gin.Context) {
	fmt.Println("In Create Menu")
	var requestBody MenuRequestBody

	err := c.BindJSON(&requestBody)
	if err != nil {
		fmt.Println("ERROR")
	}

	fmt.Println("Print out the menur equest body")
	fmt.Println(requestBody)
	c.JSON(http.StatusOK, requestBody)

}
