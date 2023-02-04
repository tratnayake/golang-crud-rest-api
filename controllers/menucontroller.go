package controllers

import (
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkIfMenuExisxts(menuId string) bool {
	var menu entities.Menu
	database.DB.First(&menu, menuId)
	if menu.ID == 0 {
		return false
	}
	return true
}

func GetMenus(c *gin.Context) {
	var menus []entities.Menu
	database.DB.Find(&menus)
	c.JSON(http.StatusOK, menus)
}
