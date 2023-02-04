package database

import (
	"fmt"
	"golang-crud-rest-api/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect(connectionString string) {
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	DB.AutoMigrate(&entities.Menu{}, &entities.Product{})
	err := DB.Model(&entities.Menu{}).Association("Products")
	if err != nil {
		fmt.Println(err)
	}

	log.Println("Database Migration Completed...")
	DB.Model(&entities.Menu{}).Association("Products")
	DB.Model(&entities.Menu{}).Association("Products").Find(&entities.Product{})

}
