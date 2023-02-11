package controllers

import (
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
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
	database.DB.Model(&models.Menu{}).Preload("Products").Find(&menus)
	c.JSON(http.StatusOK, menus)
}
